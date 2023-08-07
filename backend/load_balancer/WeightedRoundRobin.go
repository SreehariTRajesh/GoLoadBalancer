package load_balancer

import "go_load_balancer/roundrobin"

func WeightedRoundRobinLoadBalancer() string {
	ServerURL := roundrobin.ServerConfiguration.ServerData[roundrobin.WeightedRobinIndex]
	roundrobin.IncrementWeightedRobinIndex()
	return ServerURL.ServerURL
}
