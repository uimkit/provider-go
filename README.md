# UIM Chat Provider SDK


# 目录
- [安装](#installation)

## 安装
```sh
go get github.com/uimkit/provider-go
```

## 快速开始
```go
package main

import (
    "github.com/uimkit/provider-go"
    "github.com/gwatts/gin-adapter"
)

func main() {
    providerClient := uim.NewClient(config)
    providerClient.OnSendMessage(func (message *uim.Message) {
        // send to upstream
        miniMessage := ... translate from message
        miniClient.SendMessage(miniMessage)
    })

    miniClient := miniprogram.NewClient(...)
    miniClient.OnMessage(func (message *miniprogram.Message) {
        // send to downstream
        uimMessage := ... translate from message
        providerClient.NewMessage(message)
    })

    engine := gin.Default()
    engine.Post("uim/webhook", adapter.Wrap(client.WebhookHandler))
    engine.Post("miniprogram/webhook", miniClient.WebhookHandler)

    engine.Run(":8910")
}
```