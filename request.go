package uim

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	HTTP  = "HTTP"
	HTTPS = "HTTPS"

	Json = "application/json"
	Xml  = "application/xml"
	Raw  = "application/octet-stream"
	Form = "application/x-www-form-urlencoded"

	Header = "header"
	Query  = "query"
	Body   = "body"
	Path   = "path"
)

type Request interface {
	SetMethod(method string)
	GetMethod() string
	SetScheme(scheme string)
	GetScheme() string
	SetDomain(domain string)
	GetDomain() string
	SetPort(port int32)
	GetPort() int32
	SetPath(path string)
	GetPath() string
	AddHeaderParam(key, value string)
	GetHeaders() map[string]string
	AddQueryParam(key, value string)
	GetQueryParams() map[string]string
	AddFormParam(key, value string)
	GetFormParams() map[string]string
	AddPathParam(key, value string)
	GetPathParams() map[string]string
	SetReadTimeout(readTimeout time.Duration)
	GetReadTimeout() time.Duration
	SetConnectTimeout(connectTimeout time.Duration)
	GetConnectTimeout() time.Duration
	SetHTTPSInsecure(isInsecure bool)
	GetHTTPSInsecure() *bool
	SetVersion(version string)
	GetVersion() string
	SetStringToSign(stringToSign string)
	GetStringToSign() string
	GetUserAgent() map[string]string
	SetContentType(string)
	GetContentType() (string, bool)
	SetContent(content []byte)
	GetContent() []byte
	SetAcceptFormat(format string)
	GetAcceptFormat() string
	GetBodyReader() io.Reader
	BuildUrl() string
}

type BaseRequest struct {
	method         string
	scheme         string
	domain         string
	port           int32
	path           string
	headers        map[string]string
	queryParams    map[string]string
	formParams     map[string]string
	pathParams     map[string]string
	readTimeout    time.Duration
	connectTimeout time.Duration
	isInsecure     *bool
	version        string
	stringToSign   string
	userAgent      map[string]string
	content        []byte
	acceptFormat   string
}

func (request *BaseRequest) SetMethod(method string) {
	request.method = method
}

func (request *BaseRequest) GetMethod() string {
	return request.method
}

func (request *BaseRequest) SetScheme(scheme string) {
	request.scheme = scheme
}

func (request *BaseRequest) GetScheme() string {
	return request.scheme
}

func (request *BaseRequest) SetDomain(domain string) {
	request.domain = domain
}

func (request *BaseRequest) GetDomain() string {
	return request.domain
}

func (request *BaseRequest) SetPort(port int32) {
	request.port = port
}

func (request *BaseRequest) GetPort() int32 {
	return request.port
}

func (request *BaseRequest) SetPath(path string) {
	request.path = path
}

func (request *BaseRequest) GetPath() string {
	return request.path
}

func (request *BaseRequest) AddHeaderParam(key, value string) {
	request.headers[key] = value
}

func (request *BaseRequest) GetHeaders() map[string]string {
	return request.headers
}

func (request *BaseRequest) AddQueryParam(key, value string) {
	request.queryParams[key] = value
}

func (request *BaseRequest) GetQueryParams() map[string]string {
	return request.queryParams
}

func (request *BaseRequest) AddFormParam(key, value string) {
	request.formParams[key] = value
}

func (request *BaseRequest) GetFormParams() map[string]string {
	return request.formParams
}

func (request *BaseRequest) AddPathParam(key, value string) {
	request.pathParams[key] = value
}

func (request *BaseRequest) GetPathParams() map[string]string {
	return request.pathParams
}

func (request *BaseRequest) SetReadTimeout(readTimeout time.Duration) {
	request.readTimeout = readTimeout
}

func (request *BaseRequest) GetReadTimeout() time.Duration {
	return request.readTimeout
}

func (request *BaseRequest) SetConnectTimeout(connectTimeout time.Duration) {
	request.connectTimeout = connectTimeout
}

func (request *BaseRequest) GetConnectTimeout() time.Duration {
	return request.connectTimeout
}

func (request *BaseRequest) SetHTTPSInsecure(isInsecure bool) {
	request.isInsecure = &isInsecure
}

func (request *BaseRequest) GetHTTPSInsecure() *bool {
	return request.isInsecure
}

func (request *BaseRequest) SetVersion(version string) {
	request.version = version
}

func (request *BaseRequest) GetVersion() string {
	return request.version
}

func (request *BaseRequest) SetStringToSign(stringToSign string) {
	request.stringToSign = stringToSign
}

func (request *BaseRequest) GetStringToSign() string {
	return request.stringToSign
}

func (request *BaseRequest) GetUserAgent() map[string]string {
	return request.userAgent
}

func (request *BaseRequest) AppendUserAgent(key, value string) {
	if request.userAgent == nil {
		request.userAgent = make(map[string]string)
	}
	newkey := true
	if strings.ToLower(key) != "core" && strings.ToLower(key) != "go" {
		for tag := range request.userAgent {
			if tag == key {
				request.userAgent[tag] = value
				newkey = false
			}
		}
		if newkey {
			request.userAgent[key] = value
		}
	}
}

func (request *BaseRequest) SetContentType(contentType string) {
	request.AddHeaderParam("Content-Type", contentType)
}

func (request *BaseRequest) GetContentType() (contentType string, contains bool) {
	contentType, contains = request.headers["Content-Type"]
	return
}

func (request *BaseRequest) SetContent(content []byte) {
	request.content = content
}

func (request *BaseRequest) GetContent() []byte {
	return request.content
}

func (request *BaseRequest) SetAcceptFormat(format string) {
	request.acceptFormat = format
}

func (request *BaseRequest) GetAcceptFormat() string {
	return request.acceptFormat
}

func (request *BaseRequest) GetBodyReader() io.Reader {
	if request.formParams != nil && len(request.formParams) > 0 {
		formString := getUrlFormedMap(request.formParams)
		return strings.NewReader(formString)
	} else if len(request.content) > 0 {
		return bytes.NewReader(request.content)
	} else {
		return nil
	}
}

func (request *BaseRequest) buildPath() string {
	path := request.path
	for key, value := range request.pathParams {
		path = strings.Replace(path, ":"+key, value, 1)
	}
	return path
}

func (request *BaseRequest) BuildUrl() string {
	scheme := strings.ToLower(request.scheme)
	domain := request.domain
	port := request.port
	path := request.buildPath()
	querystring := getUrlFormedMap(request.queryParams)
	url := fmt.Sprintf("%s://%s", scheme, domain)
	if port > 0 {
		url = fmt.Sprintf("%s:%d", url, port)
	}
	url = fmt.Sprintf("%s%s", url, path)
	if len(querystring) > 0 {
		url = fmt.Sprintf("%s?%s", url, querystring)
	}
	return url
}

func (request *BaseRequest) String() string {
	resultBuilder := bytes.Buffer{}

	mapOutput := func(m map[string]string) {
		if len(m) > 0 {
			sortedKeys := make([]string, 0)
			for k := range m {
				sortedKeys = append(sortedKeys, k)
			}

			// sort 'string' key in increasing order
			sort.Strings(sortedKeys)

			for _, key := range sortedKeys {
				resultBuilder.WriteString(key + ": " + m[key] + "\n")
			}
		}
	}

	// Request Line
	resultBuilder.WriteString(fmt.Sprintf("%s %s\n", request.method, request.buildPath()))

	// Headers
	resultBuilder.WriteString("Host" + ": " + request.domain + "\n")
	mapOutput(request.headers)

	resultBuilder.WriteString("\n")
	// Body
	if len(request.GetContent()) > 0 {
		resultBuilder.WriteString(string(request.content) + "\n")
	} else {
		mapOutput(request.formParams)
	}

	return resultBuilder.String()
}

func NewBaseRequest() (request *BaseRequest) {
	request = &BaseRequest{
		method:       http.MethodPost,
		acceptFormat: Json,
		headers: map[string]string{
			"x-sdk-client":      "golang/1.0.0",
			"x-sdk-invoke-type": "normal",
			"Accept-Encoding":   "identity",
			"Content-Type":      Json,
		},
		queryParams: make(map[string]string),
		formParams:  make(map[string]string),
		pathParams:  make(map[string]string),
	}
	return
}

func NewBaseRequestWithPath(path string) (request *BaseRequest) {
	request = NewBaseRequest()
	request.path = path
	return
}

func initParams(request Request) (err error) {
	requestValue := reflect.ValueOf(request).Elem()
	if err = flatRepeatedList(requestValue, request, "", ""); err != nil {
		return
	}
	return
}

// 通过字段标签填充请求数据，支持标签 name、position、type。其中：
// name 代表字段名，必填。
// position 是参数位置，必填，有 header 作为 HTTP 头，query 作为 query 参数，body 作为 form 表单参数，path 作为路径参数。
// type 用来表明字段的数据类型，可选。数据类型包括：
// 为空默认代表字符串，如果值实际是 map，则为 json.Marshal 序列化后的 json 字符串；
// Json 代表字段值可以通过 json.Marshal 序列化为 json 字符串；
// Struct 代表字段值是嵌套对象；
// Map 代表字段值是 map；
// Repeated 代表字段值是 slice 或指向 slice 的指针。
func flatRepeatedList(dataValue reflect.Value, request Request, position, prefix string) (err error) {
	dataType := dataValue.Type()
	for i := 0; i < dataType.NumField(); i++ {
		field := dataType.Field(i)
		name, containsNameTag := field.Tag.Lookup("name")
		if !containsNameTag {
			continue
		}
		fieldPosition := position
		if fieldPosition == "" {
			fieldPosition, _ = field.Tag.Lookup("position")
		}
		typeTag, containsTypeTag := field.Tag.Lookup("type")
		if !containsTypeTag {
			// simple param
			key := prefix + name
			value := dataValue.Field(i).String()
			if dataValue.Field(i).Kind().String() == "map" {
				byt, _ := json.Marshal(dataValue.Field(i).Interface())
				value = string(byt)
				if value == "null" {
					value = ""
				}
			}
			err = addParam(request, fieldPosition, key, value)
			if err != nil {
				return
			}
		} else if typeTag == "Repeated" {
			// repeated param
			err = handleRepeatedParams(request, dataValue, prefix, name, fieldPosition, i)
			if err != nil {
				return
			}
		} else if typeTag == "Struct" {
			err = handleStruct(request, dataValue, prefix, name, fieldPosition, i)
			if err != nil {
				return
			}
		} else if typeTag == "Map" {
			err = handleMap(request, dataValue, prefix, name, fieldPosition, i)
			if err != nil {
				return err
			}
		} else if typeTag == "Json" {
			byt, err := json.Marshal(dataValue.Field(i).Interface())
			if err != nil {
				err = NewClientError(JsonMarshalErrorCode, JsonMarshalErrorMessage, err)
				return err
			}
			key := prefix + name
			err = addParam(request, fieldPosition, key, string(byt))
			if err != nil {
				return err
			}
		}
	}
	return
}

func handleRepeatedParams(request Request, dataValue reflect.Value, prefix, name, fieldPosition string, index int) (err error) {
	repeatedFieldValue := dataValue.Field(index)
	if repeatedFieldValue.Kind() != reflect.Slice {
		repeatedFieldValue = repeatedFieldValue.Elem()
	}
	if repeatedFieldValue.IsValid() && !repeatedFieldValue.IsNil() {
		for m := 0; m < repeatedFieldValue.Len(); m++ {
			elementValue := repeatedFieldValue.Index(m)
			key := prefix + name + "." + strconv.Itoa(m+1)
			if elementValue.Type().Kind().String() == "string" {
				value := elementValue.String()
				err = addParam(request, fieldPosition, key, value)
				if err != nil {
					return
				}
			} else {
				err = flatRepeatedList(elementValue, request, fieldPosition, key+".")
				if err != nil {
					return
				}
			}
		}
	}
	return nil
}

func handleParam(request Request, dataValue reflect.Value, key, fieldPosition string) (err error) {
	if dataValue.Type().String() == "[]string" {
		if dataValue.IsNil() {
			return
		}
		for j := 0; j < dataValue.Len(); j++ {
			err = addParam(request, fieldPosition, key+"."+strconv.Itoa(j+1), dataValue.Index(j).String())
			if err != nil {
				return
			}
		}
	} else {
		if dataValue.Type().Kind().String() == "string" {
			value := dataValue.String()
			err = addParam(request, fieldPosition, key, value)
			if err != nil {
				return
			}
		} else if dataValue.Type().Kind().String() == "struct" {
			err = flatRepeatedList(dataValue, request, fieldPosition, key+".")
			if err != nil {
				return
			}
		} else if dataValue.Type().Kind().String() == "int" {
			value := dataValue.Int()
			err = addParam(request, fieldPosition, key, strconv.Itoa(int(value)))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func handleMap(request Request, dataValue reflect.Value, prefix, name, fieldPosition string, index int) (err error) {
	valueField := dataValue.Field(index)
	if valueField.IsValid() && !valueField.IsNil() {
		iter := valueField.MapRange()
		for iter.Next() {
			k := iter.Key()
			v := iter.Value()
			key := prefix + name + ".#" + strconv.Itoa(k.Len()) + "#" + k.String()
			if v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
				elementValue := v.Elem()
				err = handleParam(request, elementValue, key, fieldPosition)
				if err != nil {
					return err
				}
			} else if v.IsValid() && v.IsNil() {
				err = handleParam(request, v, key, fieldPosition)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func handleStruct(request Request, dataValue reflect.Value, prefix, name, fieldPosition string, index int) (err error) {
	valueField := dataValue.Field(index)
	if valueField.IsValid() && valueField.String() != "" {
		valueFieldType := valueField.Type()
		for m := 0; m < valueFieldType.NumField(); m++ {
			fieldName := valueFieldType.Field(m).Name
			elementValue := valueField.FieldByName(fieldName)
			key := prefix + name + "." + fieldName
			if elementValue.Type().String() == "[]string" {
				if elementValue.IsNil() {
					continue
				}
				for j := 0; j < elementValue.Len(); j++ {
					err = addParam(request, fieldPosition, key+"."+strconv.Itoa(j+1), elementValue.Index(j).String())
					if err != nil {
						return
					}
				}
			} else {
				if elementValue.Type().Kind().String() == "string" {
					value := elementValue.String()
					err = addParam(request, fieldPosition, key, value)
					if err != nil {
						return
					}
				} else if elementValue.Type().Kind().String() == "struct" {
					err = flatRepeatedList(elementValue, request, fieldPosition, key+".")
					if err != nil {
						return
					}
				} else if !elementValue.IsNil() {
					repeatedFieldValue := elementValue.Elem()
					if repeatedFieldValue.IsValid() && !repeatedFieldValue.IsNil() {
						for m := 0; m < repeatedFieldValue.Len(); m++ {
							elementValue := repeatedFieldValue.Index(m)
							if elementValue.Type().Kind().String() == "string" {
								value := elementValue.String()
								err := addParam(request, fieldPosition, key+"."+strconv.Itoa(m+1), value)
								if err != nil {
									return err
								}
							} else {
								err = flatRepeatedList(elementValue, request, fieldPosition, key+"."+strconv.Itoa(m+1)+".")
								if err != nil {
									return
								}
							}
						}
					}
				}
			}
		}
	}
	return nil
}

func addParam(request Request, position, name, value string) (err error) {
	if len(value) > 0 {
		switch position {
		case Header:
			request.AddHeaderParam(name, value)
		case Query:
			request.AddQueryParam(name, value)
		case Path:
			request.AddPathParam(name, value)
		case Body:
			request.AddFormParam(name, value)
		default:
			errMsg := fmt.Sprintf(UnsupportedParamPositionErrorMessage, position)
			err = NewClientError(UnsupportedParamPositionErrorCode, errMsg, nil)
		}
	}
	return
}
