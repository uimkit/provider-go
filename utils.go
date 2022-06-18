package uim

import (
	"encoding/json"
	"net/url"
	"reflect"
	"strconv"
	"time"
)

func timeISO8601() (timeStr string) {
	gmt := time.FixedZone("GMT", 0)
	return time.Now().In(gmt).Format("2006-01-02T15:04:05Z")
}

func getUrlFormedMap(source map[string]string) (urlEncoded string) {
	urlEncoder := url.Values{}
	for key, value := range source {
		urlEncoder.Add(key, value)
	}
	urlEncoded = urlEncoder.Encode()
	return
}

func initStructWithDefaultTag(bean interface{}) {
	configType := reflect.TypeOf(bean)
	for i := 0; i < configType.Elem().NumField(); i++ {
		field := configType.Elem().Field(i)
		defaultValue := field.Tag.Get("default")
		if defaultValue == "" {
			continue
		}
		setter := reflect.ValueOf(bean).Elem().Field(i)
		switch field.Type.String() {
		case "int":
			intValue, _ := strconv.ParseInt(defaultValue, 10, 64)
			setter.SetInt(intValue)
		case "time.Duration":
			intValue, _ := strconv.ParseInt(defaultValue, 10, 64)
			setter.SetInt(intValue)
		case "string":
			setter.SetString(defaultValue)
		case "bool":
			boolValue, _ := strconv.ParseBool(defaultValue)
			setter.SetBool(boolValue)
		}
	}
}

func toString(object interface{}) string {
	byt, err := json.Marshal(object)
	if err != nil {
		return ""
	}
	return string(byt)
}
