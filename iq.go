package provider

type SendIQRequest struct {
	*BaseRequest
	ID      string `json:"id,omitempty"`
	Type    string `json:"type,omitempty"`
	From    string `json:"from,omitempty"`
	To      string `json:"to,omitempty"`
	Payload any    `json:"payload,omitempty"`
}

func NewSendIQRequest(id, t, from, to string, payload any) *SendIQRequest {
	return &SendIQRequest{
		BaseRequest: NewBaseRequestWithPath("send_iq"),
		ID:          id,
		Type:        t,
		From:        from,
		To:          to,
		Payload:     payload,
	}
}

type SendIQResponse struct {
	*BaseResponse
	ID      string `json:"id,omitempty"`
	Type    string `json:"type,omitempty"`
	From    string `json:"from,omitempty"`
	To      string `json:"to,omitempty"`
	Payload any    `json:"payload,omitempty"`
}

func NewSendIQResponse() *SendIQResponse {
	return &SendIQResponse{
		BaseResponse: &BaseResponse{},
	}
}

func (client *Client) SendIQ(request *SendIQRequest) (response *SendIQResponse, err error) {
	response = NewSendIQResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) SendIQAsync(request *SendIQRequest, callback func(response *SendIQResponse, err error)) {
	err := client.AddAsyncTask(func() {
		callback(client.SendIQ(request))
	})
	if err != nil {
		callback(nil, err)
	}
}
