package loadbalance

import (
	"context"
	"math/rand"

	"github.com/ironzhang/superlib/superutil/supermodel"
)

// WRLoadBalancer 权重随机均衡算法
type WRLoadBalancer struct {
}

// Pickup 按权重随机挑选地址节点
func (p *WRLoadBalancer) Pickup(ctx context.Context, domain, cluster string,
	endpoints []supermodel.Endpoint) (supermodel.Endpoint, error) {
	if len(endpoints) <= 0 {
		return supermodel.Endpoint{}, ErrInvalidEndpoints
	}

	// 按权重随机挑选节点
	total, availables := calcAvailableEndpoints(endpoints)
	if total > 0 && len(availables) > 0 {
		n := 0
		random := rand.Intn(total)
		for _, ep := range availables {
			n += ep.Weight
			if random < n {
				return ep, nil
			}
		}
	}

	// 随机挑选一个节点
	i := rand.Intn(len(endpoints))
	return endpoints[i], nil
}

func calcAvailableEndpoints(endpoints []supermodel.Endpoint) (int, []supermodel.Endpoint) {
	total := 0
	results := make([]supermodel.Endpoint, 0, len(endpoints))
	for _, ep := range endpoints {
		if ep.State == supermodel.Enabled {
			total += ep.Weight
			results = append(results, ep)
		}
	}
	return total, results
}
