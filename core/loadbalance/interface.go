package loadbalance

import (
	"context"

	"github.com/ironzhang/superlib/superutil/supermodel"
)

type LoadBalancer interface {
	Pickup(ctx context.Context, domain, cluster string, endpoints []supermodel.Endpoint) (supermodel.Endpoint, error)
}
