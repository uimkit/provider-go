package uim

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	cloudevents "github.com/cloudevents/sdk-go/v2"
)

func TestWebhook(t *testing.T) {
	client, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}

	ts := httptest.NewServer(http.NotFoundHandler())
	defer ts.Close()

	ts.Config.Handler = http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		body, _ := ioutil.ReadAll(req.Body)
		event, err := client.Webhook(req.Header, body)
		if err != nil {
			fmt.Println("Webhook is invalid :(")
		} else {
			fmt.Printf("%+v\n", event)
		}
	})
}

func TestEvent(t *testing.T) {
	sentAt := time.Now()
	message := &Message{
		MessageId:        "1",
		ConversationType: ConversationTypePrivate,
		Seq:              1,
		SentAt:           &sentAt,
		Payload: &MessagePayload{
			Type: MessageTypeText,
			Body: &TextMessageBody{
				Content: "你好",
			},
		},
	}

	ce := cloudevents.NewEvent()
	ce.SetType(UIMEventSendMessage)
	ce.SetData(cloudevents.ApplicationJSON, message)
	ce.SetSource(UIMEventSource)

	client, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}

	client.OnSendMessage(func(message *SendMessage) error {
		t.Logf("OnSendMessage: %v\n", message)
		return nil
	})

	if err = client.processEvent(&ce); err != nil {
		t.Fatal(err)
		return
	}
}
