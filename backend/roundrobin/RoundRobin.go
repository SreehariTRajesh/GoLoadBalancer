package roundrobin

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type ServerDetails struct {
	ServerURL            string `json:"ServerURL"`
	ServerHealthCheckURL string `json:"HealthCheckURL"`
}

type ServerConfig struct {
	ServerCount int             `json:"ServerCount"`
	ServerData  []ServerDetails `json:"ServerData"`
}

var RobinIndex int
var RobinCount int
var ServerConfiguration ServerConfig

func InitializeRoundRobin() {
	RobinIndex = 0
	filePath, _ := filepath.Abs("configurations/ServerData.json")
	jsonFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error", err)
	}
	byteStream, _ := io.ReadAll(jsonFile)
	json.Unmarshal(byteStream, &ServerConfiguration)
	RobinCount = ServerConfiguration.ServerCount
}

func IncrementRobinIndex() {
	RobinIndex = (RobinIndex + 1) % RobinCount
}
