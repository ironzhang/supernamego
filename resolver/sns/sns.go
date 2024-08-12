package sns

import (
	"context"

	"github.com/ironzhang/superlib/superutil/supermodel"

	"github.com/ironzhang/supernamego/resolver"
)

type snsResolver struct {
}

func (p snsResolver) Resolve(ctx context.Context, domain string, rctx map[string]string) (supermodel.Cluster, error) {
	return supernameResolver.Resolve(ctx, domain, rctx)
}

func init() {
	resolver.Register("sns", snsResolver{})
}
