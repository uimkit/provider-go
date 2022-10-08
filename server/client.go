package server

import (
	"fmt"

	uim "github.com/uimkit/provider-go"
)

// UIM 服务的 SDK，可用于调用 Provider 服务，响应 Provider 的事件
type Client struct {
	*uim.Client
}

func NewClient(opts ...uim.Option) *Client {
	return &Client{
		Client: uim.NewClient(opts...),
	}
}

func WithServerName(name string) uim.Option {
	return uim.WithEventSource(fmt.Sprintf("uim/%s", name))
}
