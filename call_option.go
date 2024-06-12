package supernamego

import "github.com/ironzhang/supernamego/core/loadbalance"

type callInfo struct {
	LoadBalancer loadbalance.LoadBalancer
	RouteParams  map[string]string
}

type CallOption func(info *callInfo)

func SetLoadBalancer(lb loadbalance.LoadBalancer) CallOption {
	return func(info *callInfo) {
		info.LoadBalancer = lb
	}
}

func SetRouteParams(params map[string]string) CallOption {
	return func(info *callInfo) {
		info.RouteParams = params
	}
}
