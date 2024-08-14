package routepolicy

import (
	"context"
	"fmt"

	"github.com/ironzhang/superlib/superutil/supermodel"
	"github.com/ironzhang/tlog"

	"github.com/ironzhang/supernamego/core/supername/routepolicy/luascript"
	"github.com/ironzhang/supernamego/core/supername/routepolicy/selection"
	"github.com/ironzhang/supernamego/pkg/public"
)

// Policy 路由策略
type Policy struct {
	routectx map[string]string        // 路由上下文
	service  *supermodel.ServiceModel // 服务模型
	route    *supermodel.RouteModel   // 路由模型
	script   *luascript.Script        // 路由脚本
}

// NewPolicy 构建路由策略
func NewPolicy(rctx map[string]string, service *supermodel.ServiceModel, route *supermodel.RouteModel, script *luascript.Script) *Policy {
	return &Policy{
		routectx: rctx,
		service:  service,
		route:    route,
		script:   script,
	}
}

// Matches 执行路由匹配
func (p *Policy) Matches(ctx context.Context, domain string) (supermodel.Cluster, error) {
	cluster := p.matches(ctx, domain)
	for _, c := range p.service.Clusters {
		if c.Name == cluster {
			return c, nil
		}
	}
	return supermodel.Cluster{}, fmt.Errorf("%s domain can not find %s cluster: %w", domain, cluster, public.ErrClusterNotFound)
}

func (p *Policy) matches(ctx context.Context, domain string) string {
	if p.route.Policy.EnableScript {
		dests := p.matchScript(ctx, domain)
		if len(dests) > 0 {
			return pick(dests)
		}
	}

	cluster, ok := p.matchLabels(ctx)
	if ok {
		return cluster
	}

	return p.service.DefaultDestination
}

func (p *Policy) matchScript(ctx context.Context, domain string) []supermodel.Destination {
	dests, err := p.script.Matches(domain, p.routectx, p.service.Clusters)
	if err != nil {
		tlog.Named(public.LoggerName).WithContext(ctx).Warnw("script matches", "domain", domain, "rctx", p.routectx, "error", err)
		return nil
	}
	return dests
}

func (p *Policy) matchLabels(ctx context.Context) (string, bool) {
	for _, ls := range p.route.Policy.LabelSelectors {
		s := selection.NewSelector(ls...)
		for _, c := range p.service.Clusters {
			store := selection.NewStore(p.routectx, c.Labels)
			if s.Matches(store) {
				return c.Name, true
			}
		}
	}
	return "", false
}
