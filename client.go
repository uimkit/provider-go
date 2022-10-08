package uim

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/authok/go-jwt-middleware/v2/jwks"
	"github.com/authok/go-jwt-middleware/v2/validator"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

var debug Debug

func init() {
	debug = getDebug("sdk")
}

// Version will be replaced while build: -ldflags="-X uim.Version=x.x.x"
var Version = "0.0.1"
var DefaultUserAgent = fmt.Sprintf("UIMKit (%s; %s) Golang/%s Core/%s", runtime.GOOS, runtime.GOARCH, strings.Trim(runtime.Version(), "go"), Version)

type EventHandler func(*cloudevents.Event) (any, error)

type Client struct {
	options              *Options
	httpClient           *http.Client
	logger               *Logger
	asyncTaskQueue       chan func()
	isOpenAsync          bool
	eventLock            sync.RWMutex
	eventHandlers        map[string]EventHandler
	accessTokenLock      sync.Mutex
	accessToken          string
	accessTokenExpiresAt time.Time
}

func (client *Client) getHttpProxy(scheme string) (proxy *url.URL, err error) {
	if strings.ToUpper(scheme) == HTTPS {
		if client.options.HttpsProxy != "" {
			proxy, err = url.Parse(client.options.HttpsProxy)
		} else if rawurl := os.Getenv("HTTPS_PROXY"); rawurl != "" {
			proxy, err = url.Parse(rawurl)
		} else if rawurl := os.Getenv("https_proxy"); rawurl != "" {
			proxy, err = url.Parse(rawurl)
		}
	} else {
		if client.options.HttpProxy != "" {
			proxy, err = url.Parse(client.options.HttpProxy)
		} else if rawurl := os.Getenv("HTTP_PROXY"); rawurl != "" {
			proxy, err = url.Parse(rawurl)
		} else if rawurl := os.Getenv("http_proxy"); rawurl != "" {
			proxy, err = url.Parse(rawurl)
		}
	}
	return proxy, err
}

func (client *Client) getNoProxy(scheme string) []string {
	var urls []string
	if client.options.NoProxy != "" {
		urls = strings.Split(client.options.NoProxy, ",")
	} else if rawurl := os.Getenv("NO_PROXY"); rawurl != "" {
		urls = strings.Split(rawurl, ",")
	} else if rawurl := os.Getenv("no_proxy"); rawurl != "" {
		urls = strings.Split(rawurl, ",")
	}
	return urls
}

func (client *Client) getSendUserAgent(requestUserAgent map[string]string) string {
	realUserAgent := ""
	for key, value := range requestUserAgent {
		realUserAgent += fmt.Sprintf(" %s/%s", key, value)
	}
	clientUserAgent := client.options.UserAgent
	if clientUserAgent != "" {
		return realUserAgent + fmt.Sprintf(" Extra/%s", clientUserAgent)
	}
	return realUserAgent
}

func (client *Client) getHTTPSInsecure(request Request) (insecure bool) {
	if request.GetHTTPSInsecure() != nil {
		insecure = *request.GetHTTPSInsecure()
	} else {
		insecure = client.options.IsInsecure
	}
	return insecure
}

// EnableAsync enable the async task queue
func (client *Client) enableAsync(routinePoolSize, maxTaskQueueSize int32) {
	if client.isOpenAsync {
		fmt.Println("warning: Please not call EnableAsync repeatedly")
		return
	}
	client.isOpenAsync = true
	client.asyncTaskQueue = make(chan func(), maxTaskQueueSize)
	for i := 0; i < int(routinePoolSize); i++ {
		go func() {
			for {
				task, notClosed := <-client.asyncTaskQueue
				if !notClosed {
					return
				} else {
					task()
				}
			}
		}()
	}
}

func (client *Client) Shutdown() {
	if client.asyncTaskQueue != nil {
		close(client.asyncTaskQueue)
	}
	client.isOpenAsync = false
}

/**
only block when any one of the following occurs:
1. the asyncTaskQueue is full, increase the queue size to avoid this
2. Shutdown() in progressing, the client is being closed
**/
func (client *Client) AddAsyncTask(task func()) (err error) {
	if client.asyncTaskQueue != nil {
		if client.isOpenAsync {
			client.asyncTaskQueue <- task
		}
	} else {
		err = NewClientError(AsyncFunctionNotEnabledCode, AsyncFunctionNotEnabledMessage, nil)
	}
	return
}

func (client *Client) getAccessToken() (string, error) {
	if client.accessToken != "" && client.accessTokenExpiresAt.After(time.Now()) {
		return client.accessToken, nil
	}

	client.accessTokenLock.Lock()
	defer client.accessTokenLock.Unlock()

	if client.accessToken != "" && client.accessTokenExpiresAt.After(time.Now()) {
		return client.accessToken, nil
	}

	accessToken, expiresIn, err := client.Authorize()
	if err != nil {
		return "", err
	}
	expiresAt := time.Now().Add(time.Duration(expiresIn-300) * time.Second)
	client.accessToken = accessToken
	client.accessTokenExpiresAt = expiresAt
	return accessToken, nil
}

func (client *Client) buildRequest(request Request) (httpRequest *http.Request, err error) {
	// add clientVersion
	request.GetHeaders()["x-sdk-core-version"] = Version

	// add authorization
	if client.options.EnableAuthorization {
		accessToken, err := client.getAccessToken()
		if err != nil {
			return nil, err
		}
		request.GetHeaders()["authorization"] = fmt.Sprintf("Bearer %s", accessToken)
	}

	// accept format
	if accept := request.GetAcceptFormat(); accept != "" {
		request.GetHeaders()["accept"] = request.GetAcceptFormat()
	}

	if request.GetDomain() == "" {
		request.SetDomain(client.options.Domain)
	}

	if request.GetScheme() == "" {
		request.SetScheme(client.options.Scheme)
	}

	if request.GetPort() == 0 {
		request.SetPort(client.options.Port)
	}

	if request.GetBasePath() == "" {
		request.SetBasePath(client.options.BasePath)
	}

	request.SetPath(request.GetBasePath() + request.GetPath())

	// init request params
	err = initParams(request)
	if err != nil {
		return
	}

	err = marshalBody(request)
	if err != nil {
		return
	}

	httpRequest, err = buildHttpRequest(request)
	if err == nil {
		userAgent := DefaultUserAgent + client.getSendUserAgent(request.GetUserAgent())
		httpRequest.Header.Set("User-Agent", userAgent)
	}

	return
}

func (client *Client) getTimeout(request Request) (time.Duration, time.Duration) {
	readTimeout := time.Duration(0)
	connectTimeout := time.Duration(0)

	reqReadTimeout := request.GetReadTimeout()
	reqConnectTimeout := request.GetConnectTimeout()
	if reqReadTimeout > 0 {
		readTimeout = reqReadTimeout
	} else if client.options.ReadTimeout > 0 {
		readTimeout = client.options.ReadTimeout
	} else if client.httpClient.Timeout > 0 {
		readTimeout = client.httpClient.Timeout
	}

	if reqConnectTimeout > 0 {
		connectTimeout = reqConnectTimeout
	} else if client.options.ConnectTimeout > 0 {
		connectTimeout = client.options.ConnectTimeout
	}
	return readTimeout, connectTimeout
}

func timeoutDialer(connectTimeout time.Duration) func(cxt context.Context, net, addr string) (c net.Conn, err error) {
	return func(ctx context.Context, network, address string) (net.Conn, error) {
		return (&net.Dialer{
			Timeout:   connectTimeout,
			DualStack: true,
		}).DialContext(ctx, network, address)
	}
}

func (client *Client) setTimeout(request Request) {
	readTimeout, connectTimeout := client.getTimeout(request)
	client.httpClient.Timeout = readTimeout
	if trans, ok := client.httpClient.Transport.(*http.Transport); ok && trans != nil {
		trans.DialContext = timeoutDialer(connectTimeout)
		client.httpClient.Transport = trans
	} else if client.httpClient.Transport == nil {
		client.httpClient.Transport = &http.Transport{
			DialContext: timeoutDialer(connectTimeout),
		}
	}
}

func (client *Client) DoAction(request Request, response Response, opts ...RequestOption) (err error) {
	for _, opt := range opts {
		opt(request)
	}
	fieldMap := make(map[string]string)
	initLogMsg(fieldMap)
	defer func() {
		client.printLog(fieldMap, err)
	}()

	httpRequest, err := client.buildRequest(request)
	if err != nil {
		return
	}

	client.setTimeout(request)
	proxy, err := client.getHttpProxy(httpRequest.URL.Scheme)
	if err != nil {
		return err
	}
	noProxy := client.getNoProxy(httpRequest.URL.Scheme)

	var withoutProxy bool
	for _, value := range noProxy {
		if strings.HasPrefix(value, "*") {
			value = fmt.Sprintf(".%s", value)
		}
		noProxyReg, err := regexp.Compile(value)
		if err != nil {
			return err
		}
		if noProxyReg.MatchString(httpRequest.Host) {
			withoutProxy = true
			break
		}
	}

	// Set whether to ignore certificate validation.
	// Default InsecureSkipVerify is false.
	if trans, ok := client.httpClient.Transport.(*http.Transport); ok && trans != nil {
		if trans.TLSClientConfig != nil {
			trans.TLSClientConfig.InsecureSkipVerify = client.getHTTPSInsecure(request)
		} else {
			trans.TLSClientConfig = &tls.Config{
				InsecureSkipVerify: client.getHTTPSInsecure(request),
			}
		}
		if proxy != nil && !withoutProxy {
			trans.Proxy = http.ProxyURL(proxy)
		}
		client.httpClient.Transport = trans
	}

	var httpResponse *http.Response
	for retryTimes := 0; retryTimes <= int(client.options.MaxRetryTime); retryTimes++ {
		if retryTimes > 0 {
			client.printLog(fieldMap, err)
			initLogMsg(fieldMap)
		}
		putMsgToMap(fieldMap, httpRequest)

		debug("> %s %s %s", httpRequest.Method, httpRequest.URL.RequestURI(), httpRequest.Proto)
		debug("> Host: %s", httpRequest.Host)
		for key, value := range httpRequest.Header {
			debug("> %s: %v", key, strings.Join(value, ""))
		}
		debug(">")
		debug(" Retry Times: %d.", retryTimes)

		startTime := time.Now()
		fieldMap["{start_time}"] = startTime.Format("2006-01-02 15:04:05")
		httpResponse, err = client.httpClient.Do(httpRequest)
		fieldMap["{cost}"] = time.Since(startTime).String()

		if err == nil {
			fieldMap["{code}"] = strconv.Itoa(httpResponse.StatusCode)
			fieldMap["{res_headers}"] = toString(httpResponse.Header)

			debug("< %s %s", httpResponse.Proto, httpResponse.Status)
			for key, value := range httpResponse.Header {
				debug("< %s: %v", key, strings.Join(value, ""))
			}
		}
		debug("<")

		// receive error
		if err != nil {
			debug(" Error: %s.", err.Error())
			if !client.options.AutoRetry {
				return
			} else if retryTimes >= int(client.options.MaxRetryTime) {
				// timeout but reached the max retry times, return
				if strings.Contains(err.Error(), "Client.Timeout") {
					times := strconv.Itoa(retryTimes + 1)
					timeoutErrorMsg := fmt.Sprintf(TimeoutErrorMessage, times, times)
					err = NewClientError(TimeoutErrorCode, timeoutErrorMsg, err)
				} else if _, ok := err.(*url.Error); ok {
					err = NewClientError(NetworkErrorCode, NetworkErrorMessage, err)
				}
				return
			}
		}

		if isCertificateError(err) {
			return
		}

		//  if status code >= 500 or timeout, will trigger retry
		if client.options.AutoRetry && (err != nil || isServerError(httpResponse)) {
			client.setTimeout(request)
			// rewrite signatureNonce and signature
			httpRequest, err = client.buildRequest(request)
			// buildHttpRequest(request, finalSigner, regionId)
			if err != nil {
				return
			}
			continue
		}
		break
	}

	err = unmarshalResponse(response, httpResponse, request.GetAcceptFormat())
	fieldMap["{res_body}"] = response.GetHttpContentString()
	debug("%s", response.GetHttpContentString())
	return
}

func (client *Client) Authorize() (accessToken string, expiresIn int64, err error) {
	payload, _ := json.Marshal(map[string]string{
		"client_id":     client.options.ClientId,
		"client_secret": client.options.ClientSecret,
		"audience":      client.options.ClientAudience,
		"grant_type":    "client_credentials",
	})
	req, _ := http.NewRequest("POST", client.options.TokenEndpoint, bytes.NewReader(payload))
	req.Header.Add("content-type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		debug("%v", err)
		return "", 0, err
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusMultipleChoices {
		result := make(map[string]string)
		if err = json.Unmarshal(body, &result); err != nil {
			debug("%s", string(body))
			return "", 0, err
		}
		return "", 0, NewClientError(AuthenticationFailedErrorCode, AuthenticationFailedErrorMessage, nil)
	}

	result := make(map[string]any)
	if err = json.Unmarshal(body, &result); err != nil {
		debug("%s", string(body))
		return "", 0, err
	}
	accessToken = result["access_token"].(string)
	expiresIn = int64(result["expires_in"].(float64))
	return
}

func (client *Client) ValidateToken(token string) (any, error) {
	issuerURL, err := url.Parse(client.options.ServerIssuer)
	if err != nil {
		return nil, err
	}
	provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

	// Set up the validator.
	jwtValidator, err := validator.New(
		provider.KeyFunc,
		validator.RS256,
		issuerURL.String(),
		[]string{client.options.ServerAudience},
	)
	if err != nil {
		return nil, err
	}
	return jwtValidator.ValidateToken(context.TODO(), token)
}

func (client *Client) newEvent(eventType string, data any) *cloudevents.Event {
	id, _ := gonanoid.New()
	ce := cloudevents.NewEvent()
	ce.SetID(id)
	ce.SetSource(client.options.EventSource)
	ce.SetType(eventType)
	ce.SetData(cloudevents.ApplicationJSON, data)
	return &ce
}

func (client *Client) SendEvent(eventType string, data any, opts ...RequestOption) (err error) {
	event := client.newEvent(eventType, data)
	content, _ := json.Marshal(event)
	req := NewBaseRequest()
	req.SetContent(content)
	return client.DoAction(req, &BaseResponse{}, opts...)
}

func (client *Client) Invoke(commandType string, data any, resp Response, opts ...RequestOption) (Response, error) {
	command := client.newEvent(commandType, data)
	content, _ := json.Marshal(command)
	req := NewBaseRequest()
	req.SetContent(content)
	err := client.DoAction(req, resp, opts...)
	return resp, err
}

func (c *Client) OnEvent(event string, handler EventHandler) {
	c.eventLock.Lock()
	defer c.eventLock.Unlock()
	c.eventHandlers[event] = handler
}

func (c *Client) EventHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" || !strings.HasPrefix(token, "Bearer ") {
			writeError(w, NewServerError(
				UnauthorizedErrorStatus,
				UnauthorizedErrorCode,
				UnauthorizedErrorMessage,
				nil,
			))
			return
		}
		token = strings.Split(token, " ")[1]
		_, err := c.ValidateToken(token)
		if err != nil {
			writeError(w, NewServerError(
				UnauthorizedErrorStatus,
				UnauthorizedErrorCode,
				UnauthorizedErrorMessage,
				err,
			))
			return
		}

		c.eventLock.RLock()
		defer c.eventLock.RUnlock()

		body, _ := ioutil.ReadAll(r.Body)
		event := cloudevents.NewEvent()
		err = json.Unmarshal(body, &event)
		if err != nil {
			writeError(w, NewServerError(
				InvalidEventFormatErrorStatus,
				InvalidEventFormatErrorCode,
				InvalidEventFormatErrorMessage,
				nil,
			))
			return
		}

		if handler, ok := c.eventHandlers[event.Type()]; ok {
			if resp, err := handler(&event); err == nil {
				if resp == nil {
					w.WriteHeader(http.StatusOK)
					return

				} else {
					switch r.Header.Get("accept") {
					default: // Json
						if body, err := json.Marshal(resp); err != nil {
							writeError(w, NewServerError(
								UnsupportedResponseFormatErrorStatus,
								UnsupportedResponseFormatErrorCode,
								fmt.Sprintf(UnsupportedResponseFormatErrorMessage, Json),
								err,
							))
						} else {
							w.WriteHeader(http.StatusOK)
							_, _ = w.Write(body)
						}
					}
					return
				}

			} else {
				writeError(w, err)
				return
			}

		} else {
			writeError(w, NewServerError(
				UnsupportedEventTypeErrorStatus,
				UnsupportedEventTypeErrorCode,
				fmt.Sprintf(UnsupportedEventTypeErrorMessage, event.Type()),
				nil,
			))
			return
		}
	}
}

func buildHttpRequest(request Request) (httpRequest *http.Request, err error) {
	requestMethod := request.GetMethod()
	requestUrl := request.BuildUrl()
	body := request.GetBodyReader()
	httpRequest, err = http.NewRequest(requestMethod, requestUrl, body)
	if err != nil {
		return
	}
	for key, value := range request.GetHeaders() {
		httpRequest.Header[key] = []string{value}
	}
	// host is a special case
	if host, containsHost := request.GetHeaders()["Host"]; containsHost {
		httpRequest.Host = host
	}
	return
}

func marshalBody(request Request) (err error) {
	// don't overwrite if body has been set
	if request.GetContent() != nil {
		return nil
	}
	if contentType, contains := request.GetContentType(); contains {
		var content []byte
		if contentType == Json {
			if content, err = json.Marshal(request); err != nil {
				err = NewClientError(JsonMarshalErrorCode, JsonMarshalErrorMessage, err)
				return
			}
			request.SetContent(content)
		}
	}
	return
}

func isCertificateError(err error) bool {
	return err != nil && strings.Contains(err.Error(), "x509: certificate signed by unknown authority")
}

func putMsgToMap(fieldMap map[string]string, request *http.Request) {
	fieldMap["{host}"] = request.Host
	fieldMap["{method}"] = request.Method
	fieldMap["{uri}"] = request.URL.RequestURI()
	fieldMap["{pid}"] = strconv.Itoa(os.Getpid())
	fieldMap["{version}"] = strings.Split(request.Proto, "/")[1]
	hostname, _ := os.Hostname()
	fieldMap["{hostname}"] = hostname
	fieldMap["{req_headers}"] = toString(request.Header)
	fieldMap["{target}"] = request.URL.Path + request.URL.RawQuery
}

func isServerError(httpResponse *http.Response) bool {
	return httpResponse.StatusCode >= http.StatusInternalServerError
}

func NewClient(opts ...Option) (client *Client) {
	client = &Client{
		options:       NewOptions(),
		eventHandlers: make(map[string]EventHandler),
	}
	for _, opt := range opts {
		opt(client.options)
	}
	options := client.options

	client.httpClient = &http.Client{}
	if options.Transport != nil {
		client.httpClient.Transport = options.Transport
	} else if options.HttpTransport != nil {
		client.httpClient.Transport = options.HttpTransport
	}
	if options.ReadTimeout > 0 {
		client.httpClient.Timeout = options.ReadTimeout
	}

	if options.EnableAsync {
		client.enableAsync(options.GoRoutinePoolSize, options.MaxTaskQueueSize)
	}

	return client
}

func CastCommandResponse[T Response](resp Response, err error) (T, error) {
	return resp.(T), err
}

func CastEventHandler[D any](handler func(*cloudevents.Event, *D) error) EventHandler {
	return func(event *cloudevents.Event) (any, error) {
		data := new(D)
		if err := event.DataAs(data); err != nil {
			return nil, err
		}
		err := handler(event, data)
		return nil, err
	}
}

func CastCommandHandler[D any, R any](handler func(*cloudevents.Event, *D) (*R, error)) EventHandler {
	return func(event *cloudevents.Event) (any, error) {
		data := new(D)
		if err := event.DataAs(data); err != nil {
			return nil, err
		}
		return handler(event, data)
	}
}
