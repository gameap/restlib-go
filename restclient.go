package main

// #cgo CFLAGS: -g -Wall
// #include <stdlib.h>
// #include "request.h"
// #include "response.h"
// extern Response DoRequest(Request request);
import "C"

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

//export DoRequest
func DoRequest(request C.Request) C.Response {
	var response *resty.Response
	var err error

	restyClient, err := convertCRequestToRestyRequest(request)
	if err != nil {
		return C.Response {
			StatusCode: -1,
			Error: C.CString(err.Error()),
		}
	}

	restyClient.EnableTrace()

	switch request.Method {
	case C.GET:
		response, err = restyClient.Get(C.GoString(request.URL))
	case C.POST:
		restyClient.SetBody(C.GoString(request.Body))
		response, err = restyClient.Post(C.GoString(request.URL))
	case C.PUT:
		restyClient.SetBody(C.GoString(request.Body))
		response, err = restyClient.Put(C.GoString(request.URL))
	case C.PATCH:
		restyClient.SetBody(C.GoString(request.Body))
		response, err = restyClient.Patch(C.GoString(request.URL))
	case C.DELETE:
		response, err = restyClient.Delete(C.GoString(request.URL))
	case C.HEAD:
		response, err = restyClient.Head(C.GoString(request.URL))
	case C.OPTIONS:
		response, err = restyClient.Options(C.GoString(request.URL))
	default:
		return C.Response {
			StatusCode: -1,
			Error: C.CString("invalid method"),
		}
	}

	if err != nil {
		return C.Response {
			StatusCode: -1,
			Error: C.CString(err.Error()),
		}
	}
	if response == nil {
		return C.Response {
			StatusCode: -1,
			Error: C.CString("Unexpected empty response"),
		}
	}

	r, err := convertRestyResponseToCResponse(response)
	if err != nil {
		return C.Response {
			StatusCode: -1,
			Error: C.CString(err.Error()),
		}
	}

	return r
}

func convertRestyResponseToCResponse(response *resty.Response) (C.Response, error) {
	var header string

	for headerName, headerValues := range response.Header() {
		for _, headerValue := range headerValues {
			header += headerName + ": " + headerValue + "\n"
		}
	}

	return C.Response{
		StatusCode: C.int(response.StatusCode()),
		Status: C.CString(response.Status()),
		Proto: C.CString(response.Proto()),
		Body: C.CString(string(response.Body())),
		Header: C.CString(header),
	}, nil
}

func convertCRequestToRestyRequest(request C.Request) (*resty.Request, error) {
	client := resty.New()

	r := client.R()

	headers := C.GoString(request.Header)
	if headers != "" {
		for _, v := range strings.Split(headers, "\n") {
			h := strings.Split(v, ":")

			if len(h) != 2 {
				return nil, fmt.Errorf("invalid header")
			}

			r.SetHeader(h[0], h[1])
		}
	}

	userAgent := C.GoString(request.UserAgent)
	if userAgent != "" {
		r.SetHeader("User-Agent", userAgent)
	} else {
		r.SetHeader("User-Agent", "restclient-go")
	}

	return r, nil
}

func main() {}
