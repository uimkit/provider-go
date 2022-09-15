package uim

import (
	"net/http"
	"time"
)

type Options struct {
	Domain            string            `default:""`
	Provider          string            `default:""`
	Strategy          string            `default:""`
	IsInsecure        bool              `default:"false"`
	HttpProxy         string            `default:""`
	HttpsProxy        string            `default:""`
	NoProxy           string            `default:""`
	AutoRetry         bool              `default:"true"`
	MaxRetryTime      int               `default:"3"`
	UserAgent         string            `default:""`
	Debug             bool              `default:"false"`
	HttpTransport     *http.Transport   `default:""`
	Transport         http.RoundTripper `default:""`
	EnableAsync       bool              `default:"false"`
	MaxTaskQueueSize  int               `default:"1000"`
	GoRoutinePoolSize int               `default:"5"`
	Scheme            string            `default:"HTTPS"`
	ReadTimeout       time.Duration     `default:"30"`
	ConnectTimeout    time.Duration     `default:"10"`
}

func NewOptions() (options *Options) {
	options = &Options{}
	initStructWithDefaultTag(options)
	return
}

type Option func(*Options)

func WithDomain(domain string) Option {
	return func(o *Options) {
		o.Domain = domain
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

func WithAutoRetry(isAutoRetry bool, maxRetryTime int) Option {
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

func WithTimeout(readTimeout, connectTimeout, timeout time.Duration) Option {
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

func WithAsync(enableAsync bool, maxTaskQueueSize, goRoutinePoolSize int) Option {
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
