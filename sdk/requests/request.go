/*
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/uimkit/provider-go/sdk/errors"
	"github.com/uimkit/provider-go/sdk/utils"
)

const (
	HTTP  = "HTTP"
	HTTPS = "HTTPS"

	DefaultHttpPort = "80"

	GET     = "GET"
	PUT     = "PUT"
	POST    = "POST"
	DELETE  = "DELETE"
	PATCH   = "PATCH"
	HEAD    = "HEAD"
	OPTIONS = "OPTIONS"

	Json = "application/json"
	Xml  = "application/xml"
	Raw  = "application/octet-stream"
	Form = "application/x-www-form-urlencoded"

	Header = "Header"
	Query  = "Query"
	Body   = "Body"
	Path   = "Path"

	HeaderSeparator = "\n"
)

// interface
type Request interface {
	GetScheme() string
	SetScheme(scheme string)
	GetMethod() string
	GetDomain() string
	SetDomain(domain string)
	GetPort() string
	GetPathPattern() string
	GetHeaders() map[string]string
	GetQueryParams() map[string]string
	GetFormParams() map[string]string
	GetPathParams() map[string]string
	GetContent() []byte
	SetContent(content []byte)
	GetBodyReader() io.Reader
	GetVersion() string
	SetVersion(version string)
	GetAcceptFormat() string
	GetReadTimeout() time.Duration
	SetReadTimeout(readTimeout time.Duration)
	GetConnectTimeout() time.Duration
	SetConnectTimeout(connectTimeout time.Duration)
	GetHTTPSInsecure() *bool
	SetHTTPSInsecure(isInsecure bool)
	GetUserAgent() map[string]string
	SetStringToSign(stringToSign string)
	GetStringToSign() string
	BuildUrl() string
	addHeaderParam(key, value string)
	addQueryParam(key, value string)
	addFormParam(key, value string)
	addPathParam(key, value string)
}

// base class
type baseRequest struct {
	Scheme         string
	Method         string
	Domain         string
	Port           string
	PathPattern    string
	ReadTimeout    time.Duration
	ConnectTimeout time.Duration
	isInsecure     *bool
	userAgent      map[string]string
	version        string
	AcceptFormat   string
	QueryParams    map[string]string
	Headers        map[string]string
	FormParams     map[string]string
	PathParams     map[string]string
	Content        []byte
	stringToSign   string
}

func (request *baseRequest) GetQueryParams() map[string]string {
	return request.QueryParams
}

func (request *baseRequest) GetFormParams() map[string]string {
	return request.FormParams
}

func (request *baseRequest) GetPathParams() map[string]string {
	return request.PathParams
}

func (request *baseRequest) GetReadTimeout() time.Duration {
	return request.ReadTimeout
}

func (request *baseRequest) GetConnectTimeout() time.Duration {
	return request.ConnectTimeout
}

func (request *baseRequest) SetReadTimeout(readTimeout time.Duration) {
	request.ReadTimeout = readTimeout
}

func (request *baseRequest) SetConnectTimeout(connectTimeout time.Duration) {
	request.ConnectTimeout = connectTimeout
}

func (request *baseRequest) GetHTTPSInsecure() *bool {
	return request.isInsecure
}

func (request *baseRequest) SetHTTPSInsecure(isInsecure bool) {
	request.isInsecure = &isInsecure
}

func (request *baseRequest) GetContent() []byte {
	return request.Content
}

func (request *baseRequest) SetVersion(version string) {
	request.version = version
}

func (request *baseRequest) GetVersion() string {
	return request.version
}

func (request *baseRequest) SetContent(content []byte) {
	request.Content = content
}

func (request *baseRequest) GetUserAgent() map[string]string {
	return request.userAgent
}

func (request *baseRequest) AppendUserAgent(key, value string) {
	newkey := true
	if request.userAgent == nil {
		request.userAgent = make(map[string]string)
	}
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

func (request *baseRequest) addHeaderParam(key, value string) {
	request.Headers[key] = value
}

func (request *baseRequest) addQueryParam(key, value string) {
	request.QueryParams[key] = value
}

func (request *baseRequest) addFormParam(key, value string) {
	request.FormParams[key] = value
}

func (request *baseRequest) addPathParam(key, value string) {
	request.PathParams[key] = value
}

func (request *baseRequest) GetAcceptFormat() string {
	return request.AcceptFormat
}

func (request *baseRequest) GetScheme() string {
	return request.Scheme
}

func (request *baseRequest) SetScheme(scheme string) {
	request.Scheme = scheme
}

func (request *baseRequest) GetMethod() string {
	return request.Method
}

func (request *baseRequest) GetDomain() string {
	return request.Domain
}

func (request *baseRequest) SetDomain(host string) {
	request.Domain = host
}

func (request *baseRequest) GetPort() string {
	return request.Port
}

func (request *baseRequest) GetPathPattern() string {
	return request.PathPattern
}

func (request *baseRequest) GetHeaders() map[string]string {
	return request.Headers
}

func (request *baseRequest) SetContentType(contentType string) {
	request.addHeaderParam("Content-Type", contentType)
}

func (request *baseRequest) GetContentType() (contentType string, contains bool) {
	contentType, contains = request.Headers["Content-Type"]
	return
}

func (request *baseRequest) GetBodyReader() io.Reader {
	if request.FormParams != nil && len(request.FormParams) > 0 {
		formString := utils.GetUrlFormedMap(request.FormParams)
		return strings.NewReader(formString)
	} else if len(request.Content) > 0 {
		return bytes.NewReader(request.Content)
	} else {
		return nil
	}
}

func (request *baseRequest) SetStringToSign(stringToSign string) {
	request.stringToSign = stringToSign
}

func (request *baseRequest) GetStringToSign() string {
	return request.stringToSign
}

func (request *baseRequest) buildPath() string {
	path := request.PathPattern
	for key, value := range request.PathParams {
		path = strings.Replace(path, ":"+key, value, 1)
	}
	return path
}

func (request *baseRequest) buildQueryString() string {
	queryParams := request.QueryParams
	// sort QueryParams by key
	q := url.Values{}
	for key, value := range queryParams {
		q.Add(key, value)
	}
	return q.Encode()
}

func (request *baseRequest) BuildUrl() string {
	// for network trans, need url encoded
	scheme := strings.ToLower(request.Scheme)
	domain := request.Domain
	port := request.Port
	path := request.buildPath()
	querystring := request.buildQueryString()
	url := fmt.Sprintf("%s://%s", scheme, domain)
	if len(port) > 0 {
		url = fmt.Sprintf("%s:%s", url, port)
	}
	url = fmt.Sprintf("%s%s", url, path)
	if len(querystring) > 0 {
		url = fmt.Sprintf("%s?%s", url, querystring)
	}
	return url
}

func defaultBaseRequest() (request *baseRequest) {
	request = &baseRequest{
		Scheme:       "",
		AcceptFormat: "JSON",
		Method:       GET,
		QueryParams:  make(map[string]string),
		Headers: map[string]string{
			"x-sdk-client":      "golang/1.0.0",
			"x-sdk-invoke-type": "normal",
			"Accept-Encoding":   "identity",
		},
		FormParams: make(map[string]string),
		PathParams: make(map[string]string),
	}
	return
}

func InitParams(request Request) (err error) {
	requestValue := reflect.ValueOf(request).Elem()
	err = flatRepeatedList(requestValue, request, "", "")
	return
}

func flatRepeatedList(dataValue reflect.Value, request Request, position, prefix string) (err error) {
	dataType := dataValue.Type()
	for i := 0; i < dataType.NumField(); i++ {
		field := dataType.Field(i)
		name, containsNameTag := field.Tag.Lookup("name")
		fieldPosition := position
		if fieldPosition == "" {
			fieldPosition, _ = field.Tag.Lookup("position")
		}
		typeTag, containsTypeTag := field.Tag.Lookup("type")
		if containsNameTag {
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
					return err
				}
				key := prefix + name
				err = addParam(request, fieldPosition, key, string(byt))
				if err != nil {
					return err
				}
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

func handleParam(request Request, dataValue reflect.Value, prefix, key, fieldPosition string) (err error) {
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
				err = handleParam(request, elementValue, prefix, key, fieldPosition)
				if err != nil {
					return err
				}
			} else if v.IsValid() && v.IsNil() {
				err = handleParam(request, v, prefix, key, fieldPosition)
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
			request.addHeaderParam(name, value)
		case Query:
			request.addQueryParam(name, value)
		case Path:
			request.addPathParam(name, value)
		case Body:
			request.addFormParam(name, value)
		default:
			errMsg := fmt.Sprintf(errors.UnsupportedParamPositionErrorMessage, position)
			err = errors.NewClientError(errors.UnsupportedParamPositionErrorCode, errMsg, nil)
		}
	}
	return
}
