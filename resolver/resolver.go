package resolver

import (
	"context"
	"fmt"
	"strings"

	"github.com/ironzhang/superlib/superutil/supermodel"
)

type Resolver interface {
	Resolve(ctx context.Context, domain string, params map[string]string) (supermodel.Cluster, error)
}

var resolvers = make(map[string]Resolver)

func Register(scheme string, r Resolver) {
	_, ok := resolvers[scheme]
	if ok {
		panic(fmt.Sprintf("%s scheme resolver is registered", scheme))
	}
	resolvers[scheme] = r
}

func Resolve(ctx context.Context, endpoint string, params map[string]string) (supermodel.Cluster, error) {
	scheme, domain := parseEndpoint(endpoint)
	r, ok := resolvers[scheme]
	if ok {
		return r.Resolve(ctx, domain, params)
	}
	return supermodel.Cluster{}, fmt.Errorf("can not find %q scheme resolver", scheme)
}

func parseEndpoint(endpoint string) (scheme, domain string) {
	results := strings.SplitN(endpoint, "/", 2)
	if len(results) == 2 {
		return results[0], results[1]
	}
	return "passthrough", endpoint
}
