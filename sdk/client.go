/*
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package sdk

import (
	"context"
	"crypto/tls"
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

	"github.com/uimkit/provider-go/sdk/errors"
	"github.com/uimkit/provider-go/sdk/requests"
	"github.com/uimkit/provider-go/sdk/responses"
	"github.com/uimkit/provider-go/sdk/utils"
)

var debug utils.Debug

func init() {
	debug = utils.Init("sdk")
}

// Version this value will be replaced while build: -ldflags="-X sdk.version=x.x.x"
var Version = "0.0.1"
var defaultConnectTimeout = 5 * time.Second
var defaultReadTimeout = 10 * time.Second

var DefaultUserAgent = fmt.Sprintf("UIMKit (%s; %s) Golang/%s Core/%s", runtime.GOOS, runtime.GOARCH, strings.Trim(runtime.Version(), "go"), Version)

var hookDo = func(fn func(req *http.Request) (*http.Response, error)) func(req *http.Request) (*http.Response, error) {
	return fn
}

// Client the type Client
type Client struct {
	SourceIp        string
	SecureTransport string
	isInsecure      bool
	config          *Config
	httpProxy       string
	httpsProxy      string
	noProxy         string
	logger          *Logger
	userAgent       map[string]string
	httpClient      *http.Client
	asyncTaskQueue  chan func()
	readTimeout     time.Duration
	connectTimeout  time.Duration
	Domain          string
	isOpenAsync     bool
}

func (client *Client) Init() (err error) {
	panic("not support yet")
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
	return
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

func (client *Client) getHttpProxy(scheme string) (proxy *url.URL, err error) {
	if scheme == "https" {
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

func (client *Client) InitClientConfig() (config *Config) {
	if client.config != nil {
		return client.config
	} else {
		return NewConfig()
	}
}

func (client *Client) buildRequest(request requests.Request) (httpRequest *http.Request, err error) {
	// add clientVersion
	request.GetHeaders()["x-sdk-core-version"] = Version

	// resolve endpoint
	endpoint := request.GetDomain()

	if endpoint == "" && client.Domain != "" {
		endpoint = client.Domain
	}

	request.SetDomain(endpoint)
	if request.GetScheme() == "" {
		request.SetScheme(client.config.Scheme)
	}
	// init request params
	err = requests.InitParams(request)
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

func (client *Client) AppendUserAgent(key, value string) {
	newkey := true

	if client.userAgent == nil {
		client.userAgent = make(map[string]string)
	}
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

func (client *Client) BuildRequest(request requests.Request) (err error) {
	_, err = client.buildRequest(request)
	return
}

func (client *Client) getTimeout(request requests.Request) (time.Duration, time.Duration) {
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

func (client *Client) setTimeout(request requests.Request) {
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

func (client *Client) getHTTPSInsecure(request requests.Request) (insecure bool) {
	if request.GetHTTPSInsecure() != nil {
		insecure = *request.GetHTTPSInsecure()
	} else {
		insecure = client.GetHTTPSInsecure()
	}
	return insecure
}

func (client *Client) DoAction(request requests.Request, response responses.Response) (err error) {
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
		httpResponse, err = hookDo(client.httpClient.Do)(httpRequest)
		fieldMap["{cost}"] = time.Since(startTime).String()
		if err == nil {
			fieldMap["{code}"] = strconv.Itoa(httpResponse.StatusCode)
			fieldMap["{res_headers}"] = TransToString(httpResponse.Header)
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
				timeoutErrorMsg := fmt.Sprintf(errors.TimeoutErrorMessage, times, times)
				if strings.Contains(err.Error(), "Client.Timeout") {
					timeoutErrorMsg += " Read timeout. Please set a valid ReadTimeout."
				} else {
					timeoutErrorMsg += " Connect timeout. Please set a valid ConnectTimeout."
				}
				err = errors.NewClientError(errors.TimeoutErrorCode, timeoutErrorMsg, err)
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

	err = responses.Unmarshal(response, httpResponse, request.GetAcceptFormat())
	fieldMap["{res_body}"] = response.GetHttpContentString()
	debug("%s", response.GetHttpContentString())
	// wrap server errors
	if serverErr, ok := err.(*errors.ServerError); ok {
		var wrapInfo = map[string]string{}
		serverErr.RespHeaders = response.GetHttpHeaders()
		wrapInfo["StringToSign"] = request.GetStringToSign()
		err = errors.WrapServerError(serverErr, wrapInfo)
	}
	return
}

func isCertificateError(err error) bool {
	if err != nil && strings.Contains(err.Error(), "x509: certificate signed by unknown authority") {
		return true
	}
	return false
}

func putMsgToMap(fieldMap map[string]string, request *http.Request) {
	fieldMap["{host}"] = request.Host
	fieldMap["{method}"] = request.Method
	fieldMap["{uri}"] = request.URL.RequestURI()
	fieldMap["{pid}"] = strconv.Itoa(os.Getpid())
	fieldMap["{version}"] = strings.Split(request.Proto, "/")[1]
	hostname, _ := os.Hostname()
	fieldMap["{hostname}"] = hostname
	fieldMap["{req_headers}"] = TransToString(request.Header)
	fieldMap["{target}"] = request.URL.Path + request.URL.RawQuery
}

func buildHttpRequest(request requests.Request) (httpRequest *http.Request, err error) {
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

func isServerError(httpResponse *http.Response) bool {
	return httpResponse.StatusCode >= http.StatusInternalServerError
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
		err = errors.NewClientError(errors.AsyncFunctionNotEnabledCode, errors.AsyncFunctionNotEnabledMessage, nil)
	}
	return
}

func (client *Client) GetConfig() *Config {
	return client.config
}

func NewClient() (client *Client, err error) {
	client = &Client{}
	err = client.Init()
	return
}

func NewClientWithOptions(config *Config) (client *Client, err error) {
	client = &Client{}
	err = client.InitWithOptions(config)
	return
}

func (client *Client) Shutdown() {
	if client.asyncTaskQueue != nil {
		close(client.asyncTaskQueue)
	}

	client.isOpenAsync = false
}
