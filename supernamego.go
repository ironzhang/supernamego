package supernamego

import (
	"context"

	"github.com/ironzhang/supernamego/core/loadbalance"
	"github.com/ironzhang/supernamego/resolver"
	_ "github.com/ironzhang/supernamego/resolver/passthrough"
	"github.com/ironzhang/supernamego/resolver/sns"
)

// AutoSetup 无参初始化
func AutoSetup() error {
	return sns.Setup(sns.Options{})
}

// Lookup 查找地址节点
func Lookup(ctx context.Context, domain string, opts ...CallOption) (addr, cluster string, err error) {
	// 构造调用信息
	info := makeCallInfo(opts)

	// 解析域名
	c, err := resolver.Resolve(ctx, domain, info.RouteParams)
	if err != nil {
		return "", "", err
	}

	// 负载均衡
	ep, err := info.LoadBalancer.Pickup(ctx, domain, c.Name, c.Endpoints)
	if err != nil {
		return "", "", err
	}
	return ep.Addr, c.Name, nil
}

func makeCallInfo(opts []CallOption) callInfo {
	info := callInfo{}
	for _, o := range opts {
		o(&info)
	}
	if info.LoadBalancer == nil {
		info.LoadBalancer = &loadbalance.WRLoadBalancer{}
	}
	return info
}
