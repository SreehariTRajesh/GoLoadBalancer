package types

import "io"

type Request struct {
	Method string
	URL    string
	Body   io.ReadCloser
}

func CreateRequest(Method string, URL string, Body io.ReadCloser) Request {
	NewRequest := Request{
		Method: Method,
		URL:    URL,
		Body:   Body,
	}
	return NewRequest
}
