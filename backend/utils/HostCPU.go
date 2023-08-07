package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

type CPU_Usage struct {
	CoreCount       int       `json:"CoreCount"`
	PerCoreUsage    []float64 `json:"PerCoreUsage"`
	Total_CPU_Usage float64   `json:"Total_CPU_Usage"`
}

func GetHostCPUUsage(HostCPUURL string) (float64, error) {
	response, err := http.Get(HostCPUURL)
	if err != nil {
		return 0, err
	} else {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return 0, nil
		} else {
			var UsageCPU CPU_Usage
			json.Unmarshal(body, &UsageCPU)
			return UsageCPU.Total_CPU_Usage, nil
		}
	}
}
