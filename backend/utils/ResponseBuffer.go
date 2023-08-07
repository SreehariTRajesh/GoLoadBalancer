package utils

import (
	"net/http"
)

var ResponseBuffer []http.Response

func InsertResponseToBuffer(res http.Response) {
	ResponseBuffer = append(ResponseBuffer, res)
}

func CleanResponseBuffer(Size int) []http.Response {
	CurrentResponse := ResponseBuffer[:Size]
	ResponseBuffer = ResponseBuffer[Size:]
	return CurrentResponse
}
