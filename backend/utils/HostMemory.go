package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

type Memory_Usage struct {
	TotalMemory   uint64  `json: "TotalMemory"`
	UsedMemory    uint64  `json: "UsedMemory"`
	FreeMemory    uint64  `json: "FreeMemory`
	MemoryUsedPCT float64 `json: "MemoryUsedPCT"`
}

func GetHostMemoryUsage(HostMemoryURL string) (float64, error) {
	response, err := http.Get(HostMemoryURL)
	if err != nil {
		return 0, err
	} else {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return 0, nil
		} else {
			var UsageMemory Memory_Usage
			json.Unmarshal(body, &UsageMemory)
			return UsageMemory.MemoryUsedPCT, nil
		}
	}
}
