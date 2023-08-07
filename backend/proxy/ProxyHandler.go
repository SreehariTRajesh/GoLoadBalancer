package proxy

import (
	"fmt"
	"io"
	"net/http"
)

var customTransport = http.DefaultTransport

func ProxyHandler(LoadBalancingFunction func() string) func(http.ResponseWriter, *http.Request) {
	proxyhandler := func(res http.ResponseWriter, req *http.Request) {
		Method := req.Method
		Body := req.Body
		ServerURL := LoadBalancingFunction()
		proxyRequest, err := http.NewRequest(Method, ServerURL, Body)
		if err != nil {
			fmt.Println("Error:", err)
		}

		// Copy the headers from the original request to the proxy request
		for name, values := range req.Header {
			for _, value := range values {
				proxyRequest.Header.Add(name, value)
			}
		}

		// Send the proxy request using the custom transport
		response, err := customTransport.RoundTrip(proxyRequest)
		if err != nil {
			http.Error(res, "Error sending proxy request", http.StatusInternalServerError)
			return
		}
		defer response.Body.Close()

		// Copy the headers from the proxy response to the original response
		for name, values := range response.Header {
			for _, value := range values {
				res.Header().Add(name, value)
			}
		}
		io.Copy(res, response.Body)
	}
	return proxyhandler
}
