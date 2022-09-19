package uim

import (
	"net/http"
	"time"
)

type Options struct {
	AppId             string            `default:""`
	Secret            string            `default:""`
	Scheme            string            `default:"HTTPS"`
	Domain            string            `default:"api.uimkit.chat"`
	Port              int32             `default:""`
	BasePath          string            `default:"/providers/v1"`
	Provider          string            `default:""`
	Strategy          string            `default:""`
	IsInsecure        bool              `default:"false"`
	HttpProxy         string            `default:""`
	HttpsProxy        string            `default:""`
	NoProxy           string            `default:""`
	AutoRetry         bool              `default:"false"`
	MaxRetryTime      int32             `default:"3"`
	UserAgent         string            `default:""`
	Debug             bool              `default:"false"`
	HttpTransport     *http.Transport   `default:""`
	Transport         http.RoundTripper `default:""`
	EnableAsync       bool              `default:"false"`
	MaxTaskQueueSize  int32             `default:"1000"`
	GoRoutinePoolSize int32             `default:"5"`
	ReadTimeout       time.Duration     `default:"30000000000"` // 30s
	ConnectTimeout    time.Duration     `default:"10000000000"` // 10s
}

func NewOptions() (options *Options) {
	options = &Options{}
	initStructWithDefaultTag(options)
	return
}

type Option func(*Options)

func WithAppSecret(appId, secret string) Option {
	return func(o *Options) {
		o.AppId = appId
		o.Secret = secret
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

func WithProvider(provider, strategy string) Option {
	return func(o *Options) {
		o.Provider = provider
		o.Strategy = strategy
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

func WithScheme(scheme string) Option {
	return func(o *Options) {
		o.Scheme = scheme
	}
}
