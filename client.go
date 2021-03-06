package uim

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var debug Debug

func init() {
	debug = getDebug("sdk")
}

// Version this value will be replaced while build: -ldflags="-X provider.version=x.x.x"
var Version = "0.0.1"
var defaultConnectTimeout = 5 * time.Second
var defaultReadTimeout = 10 * time.Second

var DefaultUserAgent = fmt.Sprintf("UIMKit (%s; %s) Golang/%s Core/%s", runtime.GOOS, runtime.GOARCH, strings.Trim(runtime.Version(), "go"), Version)

const DefaultDomain = "api.uimkit.chat/provider/v1"

type SendMessageHandler func(message *SendMessage) error
type ListAccountsHandler func(query *ListIMAccounts) error
type UpdateUserHandler func(user *UpdateIMUser) error
type UpdateContactHandler func(contact *UpdateContact) error
type ListContactsHandler func(query *ListContacts) error
type ApplyFriendHandler func(apply *NewFriendApply) error
type ApproveFriendApplyHandler func(apply *ApproveFriendApply) error
type NewGroupHandler func(group *NewGroup) error
type UpdateGroupHandler func(group *UpdateGroup) error
type ListGroupsHandler func(query *ListGroups) error
type ApplyJoinGroupHandler func(apply *NewJoinGroupApply) error
type ApproveJoinGroupApplyHandler func(apply *ApproveJoinGroupApply) error
type InviteToGroupHandler func(invite *InviteToGroup) error
type AcceptGroupInvitationHandler func(invite *AcceptGroupInvitation) error
type ListGroupMembersHandler func(query *ListGroupMembers) error

type Client struct {
	Domain              string
	isInsecure          bool
	httpProxy           string
	httpsProxy          string
	noProxy             string
	readTimeout         time.Duration
	connectTimeout      time.Duration
	userAgent           map[string]string
	config              *Config
	httpClient          *http.Client
	logger              *Logger
	asyncTaskQueue      chan func()
	isOpenAsync         bool
	providerEventSource string

	AppId  string
	Secret string

	sendMessageHandlers           []SendMessageHandler
	listAccountsHandlers          []ListAccountsHandler
	updateUserHandlers            []UpdateUserHandler
	updateContactHandlers         []UpdateContactHandler
	listContactsHandlers          []ListContactsHandler
	applyFriendHandlers           []ApplyFriendHandler
	approveFriendApplyHandlers    []ApproveFriendApplyHandler
	newGroupHandlers              []NewGroupHandler
	updateGroupHandlers           []UpdateGroupHandler
	listGroupsHandlers            []ListGroupsHandler
	applyJoinGroupHandlers        []ApplyJoinGroupHandler
	approveJoinGroupApplyHandlers []ApproveJoinGroupApplyHandler
	inviteToGroupHandlers         []InviteToGroupHandler
	acceptGroupInvitationHandlers []AcceptGroupInvitationHandler
	listGroupMembersHandlers      []ListGroupMembersHandler
}

func (client *Client) SetHTTPSInsecure(isInsecure bool) {
	client.isInsecure = isInsecure
}

func (client *Client) GetHTTPSInsecure() bool {
	return client.isInsecure
}

func (client *Client) SetHttpsProxy(httpsProxy string) {
	client.httpsProxy = httpsProxy
}

func (client *Client) GetHttpsProxy() string {
	return client.httpsProxy
}

func (client *Client) SetHttpProxy(httpProxy string) {
	client.httpProxy = httpProxy
}

func (client *Client) GetHttpProxy() string {
	return client.httpProxy
}

func (client *Client) SetNoProxy(noProxy string) {
	client.noProxy = noProxy
}

func (client *Client) GetNoProxy() string {
	return client.noProxy
}

func (client *Client) SetTransport(transport http.RoundTripper) {
	if client.httpClient == nil {
		client.httpClient = &http.Client{}
	}
	client.httpClient.Transport = transport
}

func (client *Client) SetReadTimeout(readTimeout time.Duration) {
	client.readTimeout = readTimeout
}

func (client *Client) SetConnectTimeout(connectTimeout time.Duration) {
	client.connectTimeout = connectTimeout
}

func (client *Client) GetReadTimeout() time.Duration {
	return client.readTimeout
}

func (client *Client) GetConnectTimeout() time.Duration {
	return client.connectTimeout
}

func (client *Client) GetConfig() *Config {
	return client.config
}

func (client *Client) InitWithOptions(config *Config) (err error) {
	client.config = config

	client.httpClient = &http.Client{}
	if config.Transport != nil {
		client.httpClient.Transport = config.Transport
	} else if config.HttpTransport != nil {
		client.httpClient.Transport = config.HttpTransport
	}

	if config.Timeout > 0 {
		client.httpClient.Timeout = config.Timeout
	}

	if config.EnableAsync {
		client.EnableAsync(config.GoRoutinePoolSize, config.MaxTaskQueueSize)
	}

	if config.Provider != "" && config.Strategy != "" {
		client.providerEventSource = fmt.Sprintf(ProviderEventSource, config.Provider, config.Strategy)
	}

	return
}

func (client *Client) getHttpProxy(scheme string) (proxy *url.URL, err error) {
	if strings.ToUpper(scheme) == HTTPS {
		if client.GetHttpsProxy() != "" {
			proxy, err = url.Parse(client.httpsProxy)
		} else if rawurl := os.Getenv("HTTPS_PROXY"); rawurl != "" {
			proxy, err = url.Parse(rawurl)
		} else if rawurl := os.Getenv("https_proxy"); rawurl != "" {
			proxy, err = url.Parse(rawurl)
		}
	} else {
		if client.GetHttpProxy() != "" {
			proxy, err = url.Parse(client.httpProxy)
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
	if client.GetNoProxy() != "" {
		urls = strings.Split(client.noProxy, ",")
	} else if rawurl := os.Getenv("NO_PROXY"); rawurl != "" {
		urls = strings.Split(rawurl, ",")
	} else if rawurl := os.Getenv("no_proxy"); rawurl != "" {
		urls = strings.Split(rawurl, ",")
	}
	return urls
}

func getSendUserAgent(configUserAgent string, clientUserAgent, requestUserAgent map[string]string) string {
	realUserAgent := ""
	for key1, value1 := range clientUserAgent {
		for key2 := range requestUserAgent {
			if key1 == key2 {
				key1 = ""
			}
		}
		if key1 != "" {
			realUserAgent += fmt.Sprintf(" %s/%s", key1, value1)
		}
	}
	for key, value := range requestUserAgent {
		realUserAgent += fmt.Sprintf(" %s/%s", key, value)
	}
	if configUserAgent != "" {
		return realUserAgent + fmt.Sprintf(" Extra/%s", configUserAgent)
	}
	return realUserAgent
}

func (client *Client) getHTTPSInsecure(request Request) (insecure bool) {
	if request.GetHTTPSInsecure() != nil {
		insecure = *request.GetHTTPSInsecure()
	} else {
		insecure = client.GetHTTPSInsecure()
	}
	return insecure
}

// EnableAsync enable the async task queue
func (client *Client) EnableAsync(routinePoolSize, maxTaskQueueSize int) {
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

	// resolve endpoint
	endpoint := request.GetDomain()
	if endpoint == "" && client.Domain != "" {
		endpoint = client.Domain
	}
	if endpoint == "" {
		endpoint = DefaultDomain
	}
	request.SetDomain(endpoint)

	if request.GetScheme() == "" {
		request.SetScheme(client.config.Scheme)
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
		userAgent := DefaultUserAgent + getSendUserAgent(client.config.UserAgent, client.userAgent, request.GetUserAgent())
		httpRequest.Header.Set("User-Agent", userAgent)
	}

	return
}

func (client *Client) AppendUserAgent(key, value string) {
	if client.userAgent == nil {
		client.userAgent = make(map[string]string)
	}
	newkey := true
	if strings.ToLower(key) != "core" && strings.ToLower(key) != "go" {
		for tag := range client.userAgent {
			if tag == key {
				client.userAgent[tag] = value
				newkey = false
			}
		}
		if newkey {
			client.userAgent[key] = value
		}
	}
}

func (client *Client) getTimeout(request Request) (time.Duration, time.Duration) {
	readTimeout := defaultReadTimeout
	connectTimeout := defaultConnectTimeout

	reqReadTimeout := request.GetReadTimeout()
	reqConnectTimeout := request.GetConnectTimeout()
	if reqReadTimeout != 0*time.Millisecond {
		readTimeout = reqReadTimeout
	} else if client.readTimeout != 0*time.Millisecond {
		readTimeout = client.readTimeout
	} else if client.httpClient.Timeout != 0 {
		readTimeout = client.httpClient.Timeout
	}

	if reqConnectTimeout != 0*time.Millisecond {
		connectTimeout = reqConnectTimeout
	} else if client.connectTimeout != 0*time.Millisecond {
		connectTimeout = client.connectTimeout
	}
	return readTimeout, connectTimeout
}

func Timeout(connectTimeout time.Duration) func(cxt context.Context, net, addr string) (c net.Conn, err error) {
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
		trans.DialContext = Timeout(connectTimeout)
		client.httpClient.Transport = trans
	} else if client.httpClient.Transport == nil {
		client.httpClient.Transport = &http.Transport{
			DialContext: Timeout(connectTimeout),
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

	var flag bool
	for _, value := range noProxy {
		if strings.HasPrefix(value, "*") {
			value = fmt.Sprintf(".%s", value)
		}
		noProxyReg, err := regexp.Compile(value)
		if err != nil {
			return err
		}
		if noProxyReg.MatchString(httpRequest.Host) {
			flag = true
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
		if proxy != nil && !flag {
			trans.Proxy = http.ProxyURL(proxy)
		}
		client.httpClient.Transport = trans
	}

	var httpResponse *http.Response
	for retryTimes := 0; retryTimes <= client.config.MaxRetryTime; retryTimes++ {
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
			if !client.config.AutoRetry {
				return
			} else if retryTimes >= client.config.MaxRetryTime {
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
		if client.config.AutoRetry && (err != nil || isServerError(httpResponse)) {
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

func NewClient() (client *Client, err error) {
	client, err = NewClientWithOptions(NewConfig())
	return
}

func NewClientWithOptions(config *Config) (client *Client, err error) {
	client = &Client{}
	err = client.InitWithOptions(config)
	return
}
