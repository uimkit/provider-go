package provider

import (
	"encoding/json"
	"fmt"
)

const (
	DefaultClientErrorStatus = 400
	DefaultClientErrorCode   = "SDK.ClientError"

	UnsupportedParamPositionErrorCode    = "SDK.UnsupportedParamPosition"
	UnsupportedParamPositionErrorMessage = "Specified param position (%s) is not supported, please upgrade sdk and retry"

	AsyncFunctionNotEnabledCode    = "SDK.AsyncFunctionNotEnabled"
	AsyncFunctionNotEnabledMessage = "Async function is not enabled in client, please invoke 'client.EnableAsync' function"

	MissingParamErrorCode = "SDK.MissingParam"
	InvalidParamErrorCode = "SDK.InvalidParam"

	JsonMarshalErrorCode    = "SDK.JsonMarshalError"
	JsonMarshalErrorMessage = "Failed to marshal request"

	JsonUnmarshalErrorCode    = "SDK.JsonUnmarshalError"
	JsonUnmarshalErrorMessage = "Failed to unmarshal response, but you can get the data via response.GetHttpStatusCode() and response.GetHttpContentString()"

	TimeoutErrorCode    = "SDK.TimeoutError"
	TimeoutErrorMessage = "The request timed out %s times(%s for retry), perhaps we should have the threshold raised a little?"
)

type Error interface {
	error
	HttpStatus() int
	ErrorCode() string
	Message() string
	OriginError() error
}

type ClientError struct {
	errorCode   string
	message     string
	originError error
}

func NewClientError(errorCode, message string, originErr error) Error {
	return &ClientError{
		errorCode:   errorCode,
		message:     message,
		originError: originErr,
	}
}

func (err *ClientError) Error() string {
	clientErrMsg := fmt.Sprintf("[%s] %s", err.ErrorCode(), err.message)
	if err.originError != nil {
		return clientErrMsg + "\ncaused by:\n" + err.originError.Error()
	}
	return clientErrMsg
}

func (err *ClientError) OriginError() error {
	return err.originError
}

func (*ClientError) HttpStatus() int {
	return DefaultClientErrorStatus
}

func (err *ClientError) ErrorCode() string {
	if err.errorCode == "" {
		return DefaultClientErrorCode
	} else {
		return err.errorCode
	}
}

func (err *ClientError) Message() string {
	return err.message
}

func (err *ClientError) String() string {
	return err.Error()
}

var wrapperList = []ServerErrorWrapper{}

type ServerError struct {
	RespHeaders map[string][]string
	httpStatus  int
	requestId   string
	hostId      string
	errorCode   string
	comment     string
	recommend   string
	message     string
}

type ServerErrorWrapper interface {
	tryWrap(error *ServerError, wrapInfo map[string]string) bool
}

func (err *ServerError) Error() string {
	return fmt.Sprintf("SDK.ServerError\nErrorCode: %s\nRecommend: %s\nRequestId: %s\nHostId: %s\nMessage: %s\nRespHeaders: %s",
		err.errorCode, err.comment+err.recommend, err.requestId, err.hostId, err.message, err.RespHeaders)
}

func NewServerError(httpStatus int, responseContent, comment string) Error {
	result := &ServerError{
		httpStatus: httpStatus,
		comment:    comment,
	}

	data := make(map[string]string)
	err := json.Unmarshal([]byte(responseContent), &data)
	if err == nil {
		result.requestId = data["request_id"]
		result.hostId = data["host_id"]
		result.errorCode = data["code"]
		result.recommend = data["recommend"]
		result.message = data["message"]
	}
	return result
}

func wrapServerError(originError *ServerError, wrapInfo map[string]string) *ServerError {
	for _, wrapper := range wrapperList {
		ok := wrapper.tryWrap(originError, wrapInfo)
		if ok {
			return originError
		}
	}
	return originError
}

func (err *ServerError) HttpStatus() int {
	return err.httpStatus
}

func (err *ServerError) ErrorCode() string {
	return err.errorCode
}

func (err *ServerError) Message() string {
	return err.message
}

func (err *ServerError) OriginError() error {
	return nil
}

func (err *ServerError) HostId() string {
	return err.hostId
}

func (err *ServerError) RequestId() string {
	return err.requestId
}

func (err *ServerError) Recommend() string {
	return err.recommend
}

func (err *ServerError) Comment() string {
	return err.comment
}
