package provider

type IQRequest struct {
	*BaseRequest
	ID      string `json:"id,omitempty"`
	Type    string `json:"type,omitempty"`
	From    string `json:"from,omitempty"`
	To      string `json:"to,omitempty"`
	Payload any    `json:"payload,omitempty"`
}

func NewIQRequest(id, typ, from, to string, payload any) *IQRequest {
	return &IQRequest{
		BaseRequest: NewBaseRequestWithPath("/send_iq"),
		ID:          id,
		Type:        typ,
		From:        from,
		To:          to,
		Payload:     payload,
	}
}

type IQResponse struct {
	*BaseResponse
	ID      string `json:"id,omitempty"`
	Type    string `json:"type,omitempty"`
	From    string `json:"from,omitempty"`
	To      string `json:"to,omitempty"`
	Payload any    `json:"payload,omitempty"`
}

func NewIQResponse() *IQResponse {
	return &IQResponse{
		BaseResponse: &BaseResponse{},
	}
}

func (client *Client) SendIQ(request *IQRequest) (response *IQResponse, err error) {
	response = NewIQResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) SendIQAsync(request *IQRequest, callback func(response *IQResponse, err error)) {
	err := client.AddAsyncTask(func() {
		callback(client.SendIQ(request))
	})
	if err != nil {
		callback(nil, err)
	}
}
