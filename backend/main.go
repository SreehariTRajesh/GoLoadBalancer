package main

import (
	"encoding/json"
	"go_load_balancer/load_balancer"
	"go_load_balancer/proxy"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type LoadBalancerConfig struct {
	BalancingAlgorithm string `json:"BalancingAlgorithm"`
}

func main() {

	filePath, err := filepath.Abs("configurations/LoadBalancerConfig.json")
	if err != nil {
		panic(err)
	}
	lbConfigFile, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	byteStream, _ := io.ReadAll(lbConfigFile)
	var lbConfig LoadBalancerConfig
	json.Unmarshal(byteStream, &lbConfig)
	var lbFunction func() string
	if lbConfig.BalancingAlgorithm == "RoundRobin" {
		lbFunction = load_balancer.RoundRobinLoadBalancer
	} else if lbConfig.BalancingAlgorithm == "WeightedRoundRobin" {
		lbFunction = load_balancer.WeightedRoundRobinLoadBalancer
	} else if lbConfig.BalancingAlgorithm == "HostHealthCheckCPU" {
		lbFunction = load_balancer.HostHealthCPULoadBalancer
	} else if lbConfig.BalancingAlgorithm == "HostHealthCheckMemory" {
		lbFunction = load_balancer.HostHealthMemoryLoadBalancer
	}
	http.HandleFunc("/", proxy.ProxyHandler(lbFunction))
	log.Println("Server Running on http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
