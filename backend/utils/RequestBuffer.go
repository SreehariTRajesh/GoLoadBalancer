package utils

import (
	"go_load_balancer/types"
)

var RequestBuffer []types.Request

func InsertRequestToBuffer(req types.Request) {
	RequestBuffer = append(RequestBuffer, req)
}

func CleanRequestBuffer(Size int) []types.Request {
	CurrentRequests := RequestBuffer[:Size]
	RequestBuffer = RequestBuffer[Size:]
	return CurrentRequests
}
