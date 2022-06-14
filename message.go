package provider

type MessageRequest struct {
	*BaseRequest
	Type string `json:"type,omitempty"`
	From string `json:"from,omitempty"`
	To   string `json:"to,omitempty"`
	Body any    `json:"body,omitempty"`
}

func NewMessageRequest(t, from, to string, body any) *MessageRequest {
	return &MessageRequest{
		BaseRequest: NewBaseRequestWithPath("/send_message"),
		Type:        t,
		From:        from,
		To:          to,
		Body:        body,
	}
}

func (client *Client) SendMessage(request *MessageRequest) (err error) {
	response := &BaseResponse{}
	err = client.DoAction(request, response)
	return
}
