package uim

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	cloudevents "github.com/cloudevents/sdk-go/v2"
)

const (
	UIMEventSendMessage   = "uim.send_message"   // 发送消息
	UIMEventFriendRequest = "uim.friend_request" // 好友请求
)

func (c *Client) WebhookHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		event, err := c.Webhook(r.Header, body)
		if err != nil {
			fmt.Println("Webhook is invalid :(")

			return
		}

		err = c.processEvent(event)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (c *Client) Webhook(header http.Header, body []byte) (*cloudevents.Event, error) {
	for _, token := range header["X-UIM-Key"] {
		if token == c.AppId && checkSignature(header.Get("X-UIM-Signature"), c.Secret, body) {
			var event *cloudevents.Event

			err := json.Unmarshal(body, event)
			if err != nil {
				return nil, err
			}

			return event, nil
		}
	}
	return nil, errors.New("Invalid webhook")
}

func (c *Client) triggerNewMessage(message *Message) error {
	for _, handler := range c.messageHandlers {
		if err := handler(message); err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) OnSendMessage(handler MessageHandler) {
	c.messageHandlers = append(c.messageHandlers, handler)
}

func (c *Client) processEvent(event *cloudevents.Event) error {
	switch event.Type() {
	case UIMEventSendMessage:
		var message *Message
		if err := event.DataAs(&message); err != nil {
			return err
		}

		return c.triggerNewMessage(message)
	default:
	}

	return nil
}
