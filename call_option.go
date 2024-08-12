package supernamego

import "github.com/ironzhang/supernamego/core/loadbalance"

type callInfo struct {
	LoadBalancer loadbalance.LoadBalancer
	RouteContext map[string]string
}

type CallOption func(info *callInfo)

func SetLoadBalancer(lb loadbalance.LoadBalancer) CallOption {
	return func(info *callInfo) {
		info.LoadBalancer = lb
	}
}

func SetRouteContext(rctx map[string]string) CallOption {
	return func(info *callInfo) {
		info.RouteContext = rctx
	}
}
