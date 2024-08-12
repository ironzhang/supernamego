package supername

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/ironzhang/tlog"

	"github.com/ironzhang/superlib/superutil/supermodel"

	"github.com/ironzhang/supernamego/core/supername/routepolicy"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func mergeRouteContext(src map[string]string, dst map[string]string) map[string]string {
	if len(src) <= 0 {
		return dst
	}
	if len(dst) <= 0 {
		return src
	}

	m := make(map[string]string, len(src)+len(dst))
	for k, v := range src {
		m[k] = v
	}
	for k, v := range dst {
		m[k] = v
	}
	return m
}

type lookuper struct {
	service  *supermodel.ServiceModel
	route    *supermodel.RouteModel
	policy   *routepolicy.Policy
	routectx map[string]string
}

func (p *lookuper) MakeClusterMap() map[string]supermodel.Cluster {
	m := make(map[string]supermodel.Cluster, len(p.service.Clusters))
	for _, c := range p.service.Clusters {
		m[c.Name] = c
	}
	return m
}

func (p *lookuper) MatchRouteScript(ctx context.Context, domain string, rctx map[string]string) []supermodel.Destination {
	rctx = mergeRouteContext(p.routectx, rctx)
	dests, err := p.policy.MatchRoute(domain, rctx, p.service.Clusters)
	if err != nil {
		tlog.Named("supername").WithContext(ctx).Warnw("policy match route", "domain", domain, "rctx", rctx, "error", err)
		return nil
	}
	return dests
}

func (p *lookuper) MatchRoute(ctx context.Context, domain string, rctx map[string]string) []supermodel.Destination {
	if p.route.Policy.EnableScript {
		dests := p.MatchRouteScript(ctx, domain, rctx)
		if len(dests) > 0 {
			return dests
		}
	}
	return p.route.Policy.DefaultDestinations
}

func (p *lookuper) Pick(dests []supermodel.Destination) (cluster string, err error) {
	sum := 0.0
	r := rand.Float64()
	for _, dest := range dests {
		sum += dest.Percent
		if r < sum {
			return dest.Cluster, nil
		}
	}
	if len(dests) > 0 {
		return dests[0].Cluster, nil
	}
	if len(p.service.Clusters) > 0 {
		return p.service.Clusters[0].Name, nil
	}
	return "", ErrNoAvalibaleCluster
}

func (p *lookuper) Lookup(ctx context.Context, domain string, rctx map[string]string) (supermodel.Cluster, error) {
	dests := p.MatchRoute(ctx, domain, rctx)
	cname, err := p.Pick(dests)
	if err != nil {
		return supermodel.Cluster{}, fmt.Errorf("%s domain can not pick cluster: %w", domain, err)
	}

	clusters := p.MakeClusterMap()
	c, ok := clusters[cname]
	if !ok {
		return supermodel.Cluster{}, fmt.Errorf("%s domain can not find %s cluster: %w", domain, cname, ErrClusterNotFound)
	}
	return c, nil
}
