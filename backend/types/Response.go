package types

import "io"

type Response struct {
	Header string
	URL    string
	Body   io.ReadCloser
}

func CreateResponse(Method string, URL string, Body io.ReadCloser) Request {
	NewResponse := Request{
		Method: Method,
		URL:    URL,
		Body:   Body,
	}
	return NewResponse
}
