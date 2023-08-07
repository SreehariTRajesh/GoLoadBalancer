package load_balancer

import (
	"encoding/json"
	"fmt"
	"go_load_balancer/roundrobin"
	"go_load_balancer/utils"
	"io"
	"os"
	"path/filepath"
)

type HostHealth struct {
	MaxCPU    float64 `json:"MAX_CPU_USAGE"`
	MaxMemory float64 `json:"MAX_MEMORY_USAGE"`
}

type Metric struct {
	index int
	value float64
}

var cpuIndex int
var memoryIndex int
var cpuMAX float64
var memoryMAX float64

func InitializeHostHealthLoadBalancer() {
	cpuIndex = 0
	memoryIndex = 0
	filePath, _ := filepath.Abs("configurations/HostHealthCheckConfig.json")
	jsonFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error", err)
	}
	byteStream, _ := io.ReadAll(jsonFile)
	var host HostHealth
	json.Unmarshal(byteStream, &host)
	cpuMAX = host.MaxCPU
	memoryMAX = host.MaxMemory
}

func HostHealthCPULoadBalancer() string {
	// Collecting List Of Healthy Servers
	var cpuMetrics []Metric
	for idx, server := range roundrobin.ServerConfiguration.ServerData {
		cpuURL := server.ServerHealthCheckURL + "/cpu"
		cpuUsage, err := utils.GetHostCPUUsage(cpuURL)
		if err != nil {
			panic(err)
		}
		if cpuUsage > cpuMAX {
			var m Metric
			m.index = idx
			m.value = cpuUsage
			cpuMetrics = append(cpuMetrics, m)
		}
	}
	// Running Round Robin
	if len(cpuMetrics) == 0 {
		return RoundRobinLoadBalancer()
	}
	flag := false
	for _, metric := range cpuMetrics {
		if flag == true {
			cpuIndex = metric.index
			flag = false
			break
		}
		if metric.index == cpuIndex {
			flag = true
		}
	}
	if flag == true {
		cpuIndex = 0
	}
	ServerURL := roundrobin.ServerConfiguration.ServerData[cpuIndex]
	return ServerURL.ServerURL
}

func HostHealthMemoryLoadBalancer() string {
	// Collecting List of Healthy Servers
	var memoryMetrics []Metric
	for idx, server := range roundrobin.ServerConfiguration.ServerData {
		memURL := server.ServerHealthCheckURL + "/mem"
		memUsage, err := utils.GetHostCPUUsage(memURL)
		if err != nil {
			panic(err)
		}
		if memUsage > memoryMAX {
			var m Metric
			m.index = idx
			m.value = memUsage
			memoryMetrics = append(memoryMetrics, m)
		}
	}
	// Running Round Robin
	if len(memoryMetrics) == 0 {
		return RoundRobinLoadBalancer()
	}
	flag := false
	for _, metric := range memoryMetrics {
		if flag == true {
			memoryIndex = metric.index
			flag = false
			break
		}
		if metric.index == memoryIndex {
			flag = true
		}
	}
	if flag == true {
		memoryIndex = 0
	}
	ServerURL := roundrobin.ServerConfiguration.ServerData[memoryIndex]
	return ServerURL.ServerURL
}
