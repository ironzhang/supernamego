package passthrough

import (
	"context"

	"github.com/ironzhang/superlib/superutil/supermodel"

	"github.com/ironzhang/supernamego/resolver"
)

type passthroughResolver struct {
}

func (p passthroughResolver) Resolve(ctx context.Context, domain string, params map[string]string) (supermodel.Cluster, error) {
	return supermodel.Cluster{
		Name: domain,
		Endpoints: []supermodel.Endpoint{
			{
				Addr:   domain,
				State:  supermodel.Enabled,
				Weight: 100,
			},
		},
	}, nil
}

func init() {
	resolver.Register("passthrough", passthroughResolver{})
}
