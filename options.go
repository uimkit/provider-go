package uim

import (
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Options struct {
	ClientId            string            `default:""`
	ClientSecret        string            `default:""`
	ClientAudience      string            `default:""`
	ServerIssuer        string            `default:"https://uim.cn.authok.cn/"`
	ServerAudience      string            `default:""`
	TokenEndpoint       string            `default:"https://uim.cn.authok.cn/oauth/token"`
	EnableAuthorization bool              `default:"true"`
	EventSource         string            `default:""`
	Scheme              string            `default:"HTTPS"`
	Domain              string            `default:"api.uimkit.chat"`
	Port                int32             `default:""`
	BasePath            string            `default:""`
	IsInsecure          bool              `default:"false"` // 是否可以跳过证书验证
	HttpProxy           string            `default:""`
	HttpsProxy          string            `default:""`
	NoProxy             string            `default:""`
	AutoRetry           bool              `default:"false"`
	MaxRetryTime        int32             `default:"3"`
	UserAgent           string            `default:""`
	Debug               bool              `default:"false"`
	HttpTransport       *http.Transport   `default:""`
	Transport           http.RoundTripper `default:""`
	EnableAsync         bool              `default:"false"`
	MaxTaskQueueSize    int32             `default:"1000"`
	GoRoutinePoolSize   int32             `default:"5"`
	ReadTimeout         time.Duration     `default:"30000000000"` // 30s
	ConnectTimeout      time.Duration     `default:"10000000000"` // 10s
}

func NewOptions() (options *Options) {
	options = &Options{}
	initStructWithDefaultTag(options)
	return
}

type Option func(*Options)

func WithClient(clientId, clientSecret, audience string) Option {
	return func(o *Options) {
		o.ClientId = clientId
		o.ClientSecret = clientSecret
		o.ClientAudience = audience
	}
}

func WithServer(issuer, audience string) Option {
	return func(o *Options) {
		o.ServerIssuer = issuer
		o.ServerAudience = audience
	}
}

func WithTokenEndpoint(endpoint string) Option {
	return func(o *Options) {
		o.TokenEndpoint = endpoint
	}
}

func WithAuthorization(enable bool) Option {
	return func(o *Options) {
		o.EnableAuthorization = enable
	}
}

func WithEventSource(es string) Option {
	return func(o *Options) {
		o.EventSource = es
	}
}

func WithBaseUrl(baseUrl string) Option {
	return func(o *Options) {
		parsed, _ := url.Parse(baseUrl)
		port, _ := strconv.ParseInt(parsed.Port(), 10, 64)
		WithScheme(parsed.Scheme)(o)
		WithDomain(parsed.Hostname())(o)
		WithPort(int32(port))(o)
		WithBasePath(parsed.Path)(o)
	}
}

func WithScheme(scheme string) Option {
	return func(o *Options) {
		o.Scheme = scheme
	}
}

func WithDomain(domain string) Option {
	return func(o *Options) {
		o.Domain = domain
	}
}

func WithPort(port int32) Option {
	return func(o *Options) {
		o.Port = port
	}
}

func WithBasePath(basePath string) Option {
	return func(o *Options) {
		o.BasePath = basePath
	}
}

func WithSSL(insecure bool) Option {
	return func(o *Options) {
		o.IsInsecure = insecure
	}
}

// 设置代理
// httpProxy 是 http 的代理 host
// httpsProxy 是 https 的代理 host
// noProxy 是跳过代理的 host 列表，逗号分隔
func WithProxy(httpProxy, httpsProxy, noProxy string) Option {
	return func(o *Options) {
		o.HttpProxy = httpProxy
		o.HttpsProxy = httpsProxy
		o.NoProxy = noProxy
	}
}

func WithAutoRetry(isAutoRetry bool, maxRetryTime int32) Option {
	return func(o *Options) {
		o.AutoRetry = isAutoRetry
		o.MaxRetryTime = maxRetryTime
	}
}

func WithUserAgent(userAgent string) Option {
	return func(o *Options) {
		o.UserAgent = userAgent
	}
}

func WithDebug(isDebug bool) Option {
	return func(o *Options) {
		o.Debug = isDebug
	}
}

func WithTimeout(readTimeout, connectTimeout time.Duration) Option {
	return func(o *Options) {
		o.ReadTimeout = readTimeout
		o.ConnectTimeout = connectTimeout
	}
}

func WithHttpTransport(httpTransport *http.Transport) Option {
	return func(o *Options) {
		o.HttpTransport = httpTransport
	}
}

func WithTransport(transport http.RoundTripper) Option {
	return func(o *Options) {
		o.Transport = transport
	}
}

func WithAsync(enableAsync bool, maxTaskQueueSize, goRoutinePoolSize int32) Option {
	return func(o *Options) {
		o.EnableAsync = enableAsync
		o.MaxTaskQueueSize = maxTaskQueueSize
		o.GoRoutinePoolSize = goRoutinePoolSize
	}
}
