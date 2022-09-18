package uim

import (
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

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/google/uuid"
)

var debug Debug

func init() {
	debug = getDebug("sdk")
}

// Version will be replaced while build: -ldflags="-X uim.Version=x.x.x"
var Version = "0.0.1"
var DefaultUserAgent = fmt.Sprintf("UIMKit (%s; %s) Golang/%s Core/%s", runtime.GOOS, runtime.GOARCH, strings.Trim(runtime.Version(), "go"), Version)
var DefaultDomain = "api.uimkit.chat/provider/v1"

var defaultConnectTimeout = 30 * time.Second
var defaultReadTimeout = 10 * time.Second

type EventHandler func(*cloudevents.Event) (any, error)

type Client struct {
	appId          string
	secret         string
	eventSource    string
	options        *Options
	httpClient     *http.Client
	logger         *Logger
	asyncTaskQueue chan func()
	isOpenAsync    bool
	eventLock      sync.RWMutex
	eventHandlers  map[string]EventHandler
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
func (client *Client) enableAsync(routinePoolSize, maxTaskQueueSize int) {
	if client.isOpenAsync {
		fmt.Println("warning: Please not call EnableAsync repeatedly")
		return
	}
	client.isOpenAsync = true
	client.asyncTaskQueue = make(chan func(), maxTaskQueueSize)
	for i := 0; i < routinePoolSize; i++ {
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

func (client *Client) buildRequest(request Request) (httpRequest *http.Request, err error) {
	// add clientVersion
	request.GetHeaders()["x-sdk-core-version"] = Version

	// accept format
	if accept := request.GetAcceptFormat(); accept != "" {
		request.GetHeaders()["accept"] = request.GetAcceptFormat()
	}

	// resolve endpoint
	endpoint := request.GetDomain()
	if endpoint == "" && client.options.Domain != "" {
		endpoint = client.options.Domain
	}
	if endpoint == "" {
		endpoint = DefaultDomain
	}
	request.SetDomain(endpoint)

	if request.GetScheme() == "" {
		request.SetScheme(client.options.Scheme)
	}
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
	readTimeout := defaultReadTimeout
	connectTimeout := defaultConnectTimeout

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

func (client *Client) DoAction(request Request, response Response) (err error) {
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
	for retryTimes := 0; retryTimes <= client.options.MaxRetryTime; retryTimes++ {
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
			} else if retryTimes >= client.options.MaxRetryTime {
				// timeout but reached the max retry times, return
				times := strconv.Itoa(retryTimes + 1)
				timeoutErrorMsg := fmt.Sprintf(TimeoutErrorMessage, times, times)
				if strings.Contains(err.Error(), "Client.Timeout") {
					timeoutErrorMsg += " Read timeout. Please set a valid ReadTimeout."
				} else {
					timeoutErrorMsg += " Connect timeout. Please set a valid ConnectTimeout."
				}
				err = NewClientError(TimeoutErrorCode, timeoutErrorMsg, err)
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

	// wrap server errors
	if serverErr, ok := err.(*ServerError); ok {
		var wrapInfo = map[string]string{}
		serverErr.RespHeaders = response.GetHttpHeaders()
		wrapInfo["StringToSign"] = request.GetStringToSign()
		err = wrapServerError(serverErr, wrapInfo)
	}
	return
}

func (client *Client) newEvent(eventType string, data any) *cloudevents.Event {
	ce := cloudevents.NewEvent()
	ce.SetID(uuid.NewString())
	ce.SetSource(client.eventSource)
	ce.SetType(eventType)
	ce.SetData(cloudevents.ApplicationJSON, data)
	return &ce
}

func (client *Client) SendEvent(event *cloudevents.Event) (err error) {
	content, _ := json.Marshal(event)
	req := NewBaseRequest()
	req.SetContent(content)
	return client.DoAction(req, &BaseResponse{})
}

func (client *Client) InvokeCommand(command *cloudevents.Event, resp Response) (Response, error) {
	content, _ := json.Marshal(command)
	req := NewBaseRequest()
	req.SetContent(content)
	err := client.DoAction(req, resp)
	return resp, err
}

func castCommandResponse[T Response](resp Response, err error) (T, error) {
	return resp.(T), err
}

func (c *Client) OnEvent(event string, handler EventHandler) {
	c.eventLock.Lock()
	defer c.eventLock.Unlock()
	c.eventHandlers[event] = handler
}

func castEventHandler[D any](handler func(*D) error) EventHandler {
	return func(event *cloudevents.Event) (any, error) {
		data := new(D)
		if err := event.DataAs(data); err != nil {
			return nil, err
		}
		err := handler(data)
		return nil, err
	}
}

func castCommandHandler[D any, R any](handler func(*D) (*R, error)) EventHandler {
	return func(event *cloudevents.Event) (any, error) {
		data := new(D)
		if err := event.DataAs(data); err != nil {
			return nil, err
		}
		return handler(data)
	}
}

func (c *Client) EventHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c.eventLock.RLock()
		defer c.eventLock.RUnlock()

		body, _ := ioutil.ReadAll(r.Body)
		event := cloudevents.NewEvent()
		err := json.Unmarshal(body, &event)
		if err != nil {
			debug("Handle event error: %s.", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if handler, ok := c.eventHandlers[event.Type()]; ok {
			if resp, err := handler(&event); err == nil {
				if resp == nil {
					debug("Handle event success: %+v.", &event)
					w.WriteHeader(http.StatusOK)
					return

				} else {
					debug("Handle event success: %+v, resp: %+v.", &event, resp)
					switch r.Header.Get("accept") {
					default: // Json
						body, _ := json.Marshal(resp)
						_, _ = w.Write(body)
					}
					w.WriteHeader(http.StatusOK)
					return
				}

			} else {
				debug("Handle event error: %+v.", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

		} else {
			debug("Handle unknown event: %s.", event.Type())
			w.WriteHeader(http.StatusBadRequest)
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

func NewClient(appId, secret string, opts ...Option) (client *Client) {
	client = &Client{
		appId:   appId,
		secret:  secret,
		options: NewOptions(),
	}
	for _, opt := range opts {
		opt(client.options)
	}
	options := client.options

	client.eventSource = fmt.Sprintf("provider.source/%s/%s", options.Provider, options.Strategy)

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
