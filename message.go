package provider

type SendMessageRequest struct {
	*BaseRequest
	Type string `json:"type,omitempty"`
	From string `json:"from,omitempty"`
	To   string `json:"to,omitempty"`
	Body any    `json:"body,omitempty"`
}

func NewSendMessageRequest(t, from, to string, body any) *SendMessageRequest {
	return &SendMessageRequest{
		BaseRequest: NewBaseRequestWithPath("/send_message"),
		Type:        t,
		From:        from,
		To:          to,
		Body:        body,
	}
}

type SendMessageResponse struct {
	*BaseResponse
}

func NewSendMessageResponse() *SendMessageResponse {
	return &SendMessageResponse{
		BaseResponse: &BaseResponse{},
	}
}

func (client *Client) SendMessage(request *SendMessageRequest) (response *SendMessageResponse, err error) {
	response = NewSendMessageResponse()
	err = client.DoAction(request, response)
	return
}
