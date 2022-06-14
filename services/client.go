package services

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"net/http"

	"github.com/uimkit/provider-go/sdk"
	"github.com/uimkit/provider-go/sdk/requests"
	"github.com/uimkit/provider-go/sdk/responses"
)

// Client is the sdk client struct, each func corresponds to an OpenAPI
type Client struct {
	sdk.Client
}

// NewClientWithOptions creates a sdk client with regionId/sdkConfig/credential
// this is the common api to create a sdk client
func NewClientWithOptions(config *sdk.Config) (client *Client, err error) {
	client = &Client{}
	err = client.InitWithOptions(config)
	return
}

type SendMessageRequest struct {
	*requests.CommonRequest
	Type string `json:"type,omitempty"`
	From string `json:"from,omitempty"`
	To   string `json:"to,omitempty"`
	Body any    `json:"body,omitempty"`
}

func NewSendMessageRequest(t, from, to string, body any) *SendMessageRequest {
	request := &SendMessageRequest{
		CommonRequest: requests.NewCommonRequest(),
		Type:          t,
		From:          from,
		To:            to,
		Body:          body,
	}
	request.Method = http.MethodPost
	request.PathPattern = "/send_message"
	request.SetContentType(requests.Json)
	return request
}

type SendMessageResponse struct {
	*responses.BaseResponse
}

func NewSendMessageResponse() *SendMessageResponse {
	return &SendMessageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
}

func (client *Client) SendMessage(request *SendMessageRequest) (response *SendMessageResponse, err error) {
	response = NewSendMessageResponse()
	err = client.DoAction(request, response)
	return
}

type SendIQRequest struct {
	*requests.CommonRequest
	ID      string `json:"id,omitempty"`
	Type    string `json:"type,omitempty"`
	From    string `json:"from,omitempty"`
	To      string `json:"to,omitempty"`
	Payload any    `json:"payload,omitempty"`
}

func NewSendIQRequest(id, t, from, to string, payload any) *SendIQRequest {
	request := &SendIQRequest{
		CommonRequest: requests.NewCommonRequest(),
		ID:            id,
		Type:          t,
		From:          from,
		To:            to,
		Payload:       payload,
	}
	request.Method = http.MethodPost
	request.PathPattern = "/send_iq"
	request.SetContentType(requests.Json)
	return request
}

type SendIQResponse struct {
	*responses.BaseResponse
	ID      string `json:"id,omitempty"`
	Type    string `json:"type,omitempty"`
	From    string `json:"from,omitempty"`
	To      string `json:"to,omitempty"`
	Payload any    `json:"payload,omitempty"`
}

func NewSendIQResponse() *SendIQResponse {
	return &SendIQResponse{
		BaseResponse: &responses.BaseResponse{},
	}
}

func (client *Client) SendIQ(request *SendIQRequest, callback func(response *SendIQResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SendIQResponse
		var err error
		defer close(result)
		response = NewSendIQResponse()
		err = client.DoAction(request, response)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}
