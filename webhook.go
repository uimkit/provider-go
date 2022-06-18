package uim

import "encoding/json"

type Webhook struct {
	TimeMs int            `json:"time_ms"` // the timestamp of the request
	Events []WebhookEvent `json:"events"`  // the events associated with the webhook
}

type WebhookEvent struct {
	Event string `json:"type,omitempty"`
	Data  string `json:"data,omitempty"`
}

func unmarshalledWebhook(requestBody []byte) (*Webhook, error) {
	webhook := &Webhook{}
	err := json.Unmarshal(requestBody, &webhook)
	if err != nil {
		return nil, err
	}
	return webhook, nil
}
