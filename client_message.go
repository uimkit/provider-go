package uim

import cloudevents "github.com/cloudevents/sdk-go/v2"

func (client *Client) NewMessage(message *Message) error {
	ce := cloudevents.NewEvent()
	ce.SetSource(UIMProviderEventSource)
	ce.SetType(ProviderEventNewMessage)
	ce.SetData(cloudevents.ApplicationJSON, message)

	return client.SendEvent(&ce)
}
