package uim

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Response interface {
	IsSuccess() bool
	GetHttpStatus() int
	GetHttpHeaders() map[string][]string
	GetHttpContentString() string
	GetHttpContentBytes() []byte
	GetOriginHttpResponse() *http.Response
	parseFromHttpResponse(httpResponse *http.Response) error
}

// Unmarshal object from http response body to target Response
func unmarshalResponse(response Response, httpResponse *http.Response, format string) (err error) {
	err = response.parseFromHttpResponse(httpResponse)
	if err != nil {
		return
	}
	if !response.IsSuccess() {
		err = NewServerError(response.GetHttpStatus(), response.GetHttpContentString(), "")
		return
	}

	if len(response.GetHttpContentBytes()) == 0 {
		return
	}

	if format == Json {
		err = json.Unmarshal(response.GetHttpContentBytes(), response)
		if err != nil {
			err = NewClientError(JsonUnmarshalErrorCode, JsonUnmarshalErrorMessage, err)
		}
	}
	return
}

type BaseResponse struct {
	httpStatus         int
	httpHeaders        map[string][]string
	httpContentString  string
	httpContentBytes   []byte
	originHttpResponse *http.Response
}

func (baseResponse *BaseResponse) GetHttpStatus() int {
	return baseResponse.httpStatus
}

func (baseResponse *BaseResponse) GetHttpHeaders() map[string][]string {
	return baseResponse.httpHeaders
}

func (baseResponse *BaseResponse) GetHttpContentString() string {
	return baseResponse.httpContentString
}

func (baseResponse *BaseResponse) GetHttpContentBytes() []byte {
	return baseResponse.httpContentBytes
}

func (baseResponse *BaseResponse) GetOriginHttpResponse() *http.Response {
	return baseResponse.originHttpResponse
}

func (baseResponse *BaseResponse) IsSuccess() bool {
	if baseResponse.GetHttpStatus() >= 200 && baseResponse.GetHttpStatus() < 300 {
		return true
	}
	return false
}

func (baseResponse *BaseResponse) parseFromHttpResponse(httpResponse *http.Response) (err error) {
	defer httpResponse.Body.Close()
	body, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return
	}
	baseResponse.httpStatus = httpResponse.StatusCode
	baseResponse.httpHeaders = httpResponse.Header
	baseResponse.httpContentBytes = body
	baseResponse.httpContentString = string(body)
	baseResponse.originHttpResponse = httpResponse
	return
}

func (baseResponse *BaseResponse) String() string {
	resultBuilder := bytes.Buffer{}
	// statusCode
	// resultBuilder.WriteString("\n")
	resultBuilder.WriteString(fmt.Sprintf("%s %s\n", baseResponse.originHttpResponse.Proto, baseResponse.originHttpResponse.Status))
	// httpHeaders
	//resultBuilder.WriteString("Headers:\n")
	for key, value := range baseResponse.httpHeaders {
		resultBuilder.WriteString(key + ": " + strings.Join(value, ";") + "\n")
	}
	resultBuilder.WriteString("\n")
	// content
	//resultBuilder.WriteString("Content:\n")
	resultBuilder.WriteString(baseResponse.httpContentString + "\n")
	return resultBuilder.String()
}
