package uim

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Error interface {
	error
	HttpStatus() int
	ErrorCode() string
	Message() string
	OriginError() error
}

// 客户端错误
const (
	DefaultClientErrorStatus = 400
	DefaultClientErrorCode   = "SDK.ClientError"

	UnsupportedParamPositionErrorCode    = "SDK.UnsupportedParamPosition"
	UnsupportedParamPositionErrorMessage = "Specified param position (%s) is not supported, please upgrade sdk and retry"

	AsyncFunctionNotEnabledCode    = "SDK.AsyncFunctionNotEnabled"
	AsyncFunctionNotEnabledMessage = "Async function is not enabled in client, please invoke 'client.EnableAsync' function"

	JsonMarshalErrorCode    = "SDK.JsonMarshalError"
	JsonMarshalErrorMessage = "Failed to marshal request"

	JsonUnmarshalErrorCode    = "SDK.JsonUnmarshalError"
	JsonUnmarshalErrorMessage = "Failed to unmarshal response, but you can get the data via response.GetHttpStatusCode() and response.GetHttpContentString()"

	TimeoutErrorCode    = "SDK.TimeoutError"
	TimeoutErrorMessage = "The request timed out %s times(%s for retry), perhaps we should have the threshold raised a little?"

	NetworkErrorCode    = "SDK.NetworkError"
	NetworkErrorMessage = "Failed to make the http request."

	AuthenticationFailedErrorCode    = "SDK.AuthenticationFailed"
	AuthenticationFailedErrorMessage = "Authentication failed, please check 'client_id' & 'client_secret'"
)

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

// 服务端错误
const (
	DefaultServerErrorStatus = http.StatusInternalServerError
	DefaultServerErrorCode   = "SDK.ServerError"

	UnsupportedEventTypeErrorStatus  = http.StatusBadRequest
	UnsupportedEventTypeErrorCode    = "SDK.UnsupportedEventType"
	UnsupportedEventTypeErrorMessage = "Unsupported event type \"%s\""

	InvalidEventFormatErrorStatus  = http.StatusBadRequest
	InvalidEventFormatErrorCode    = "SDK.InvalidEventFormat"
	InvalidEventFormatErrorMessage = "Event must comply with the cloudevents specification"

	UnsupportedResponseFormatErrorStatus  = http.StatusBadRequest
	UnsupportedResponseFormatErrorCode    = "SDK.UnsupportedResponseFormat"
	UnsupportedResponseFormatErrorMessage = "Could not marshal response in \"%s\", pelease set proper \"Accept\" header in request"

	InvalidEventDataErrorStatus = http.StatusBadRequest
	InvalidEventDataErrorCode   = "SDK.InvalidEventData"

	ResourceNotFoundErrorStatus = http.StatusNotFound
	ResourceNotFoundErrorCode   = "SDK.ResourceNotFound"

	UnregisteredProviderErrorStatus  = http.StatusUnauthorized
	UnregisteredProviderErrorCode    = "SDK.UnregisteredProvider"
	UnregisteredProviderErrorMessage = "Please set proper \"Provider\" and \"Strategy\" client options"
)

type ServerError struct {
	httpStatus  int
	errorCode   string
	message     string
	originError error
}

func NewServerError(httpStatus int, errorCode, message string, originError error) Error {
	return &ServerError{
		httpStatus:  httpStatus,
		errorCode:   errorCode,
		message:     message,
		originError: originError,
	}
}

func (err *ServerError) Error() string {
	serverErrMsg := fmt.Sprintf("[%s] %s", err.ErrorCode(), err.message)
	if err.originError != nil {
		return serverErrMsg + "\ncaused by:\n" + err.originError.Error()
	}
	return serverErrMsg
}

func (err *ServerError) OriginError() error {
	return err.originError
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

func parseServerError(httpStatus int, responseContent string) Error {
	result := &ServerError{
		httpStatus: httpStatus,
	}
	data := make(map[string]string)
	err := json.Unmarshal([]byte(responseContent), &data)
	if err == nil {
		result.errorCode = data["code"]
		result.message = data["message"]
	} else {
		result.errorCode = DefaultServerErrorCode
		result.message = fmt.Sprintf("Server runtime error caused by: %s", responseContent)
	}
	return result
}

func writeError(w http.ResponseWriter, err error) {
	if serverError, ok := err.(*ServerError); ok {
		b, _ := json.Marshal(map[string]string{
			"code":    serverError.errorCode,
			"message": serverError.message,
		})
		w.WriteHeader(serverError.httpStatus)
		w.Write(b)
	} else {
		b, _ := json.Marshal(map[string]string{
			"code":    DefaultServerErrorCode,
			"message": fmt.Sprintf("Server runtime error caused by: %s", err.Error()),
		})
		w.WriteHeader(DefaultServerErrorStatus)
		w.Write(b)
	}
}
