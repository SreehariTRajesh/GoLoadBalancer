package load_balancer

import (
	"go_load_balancer/roundrobin"
)

func RoundRobinLoadBalancer() string {
	ServerURL := roundrobin.ServerConfiguration.ServerData[roundrobin.RobinIndex]
	roundrobin.IncrementRobinIndex()
	return ServerURL.ServerURL
}
