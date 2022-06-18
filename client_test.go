package uim

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
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
		webhook, err := client.Webhook(req.Header, body)
		if err != nil {
			fmt.Println("Webhook is invalid :(")
		} else {
			fmt.Printf("%+v\n", webhook.Events)
		}
	})
}
