package main

import (
	"fmt"
	"go_load_balancer/load_balancer"
	"go_load_balancer/roundrobin"
)

func main() {
	roundrobin.InitializeWeightedRobin()
	roundrobin.InitializeRoundRobin()
	for i := 0; i < 100; i++ {
		url := load_balancer.WeightedRoundRobinLoadBalancer()
		fmt.Println(url)
	}
}
