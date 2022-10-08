package provider

import (
	"fmt"

	uim "github.com/uimkit/provider-go"
)

// 服务商的 SDK，可用于调用 UIM 服务，响应 UIM 的事件
type Client struct {
	*uim.Client
}

func NewClient(opts ...uim.Option) *Client {
	return &Client{
		Client: uim.NewClient(opts...),
	}
}

func WithProvider(provider, strategy string) uim.Option {
	return uim.WithEventSource(fmt.Sprintf("provider.source/%s/%s", provider, strategy))
}
