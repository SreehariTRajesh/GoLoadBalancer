package roundrobin

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
)

type WeightedServerConfig struct {
	ServerCount        int   `json:"ServerCount"`
	WeightedServerData []int `json:"ServerWeightData"`
}

var WeightedRobinServerWeights WeightedServerConfig
var WeightedRobinIndex int
var WeightedRobinCount int

func InitializeWeightedRobin() {
	WeightedRobinCount = 0
	WeightedRobinIndex = 0
	filePath, _ := filepath.Abs("configurations/WeightedLoadBalancerConfig.json")
	jsonFile, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	byteStream, _ := io.ReadAll(jsonFile)
	json.Unmarshal(byteStream, &WeightedRobinServerWeights)

}

func IncrementWeightedRobinIndex() {
	if (WeightedRobinCount+1)%WeightedRobinServerWeights.WeightedServerData[WeightedRobinIndex] == 0 {
		WeightedRobinIndex = (WeightedRobinIndex + 1) % WeightedRobinServerWeights.ServerCount
		WeightedRobinCount = 0
	} else {
		WeightedRobinCount++
	}
}
