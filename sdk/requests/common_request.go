package requests

import (
	"bytes"
	"fmt"
	"sort"
)

type CommonRequest struct {
	*baseRequest
}

func NewCommonRequest() (request *CommonRequest) {
	request = &CommonRequest{
		baseRequest: defaultBaseRequest(),
	}
	return
}

func (request *CommonRequest) String() string {
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
	resultBuilder.WriteString(fmt.Sprintf("%s %s\n", request.Method, request.buildPath()))

	// Headers
	resultBuilder.WriteString("Host" + ": " + request.Domain + "\n")
	mapOutput(request.Headers)

	resultBuilder.WriteString("\n")
	// Body
	if len(request.Content) > 0 {
		resultBuilder.WriteString(string(request.Content) + "\n")
	} else {
		mapOutput(request.FormParams)
	}

	return resultBuilder.String()
}
