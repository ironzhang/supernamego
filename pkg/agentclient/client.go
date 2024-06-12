package agentclient

import (
	"context"
	"net/http"
	"time"

	"github.com/ironzhang/superlib/httputils/httpclient"
	"github.com/ironzhang/superlib/timeutil"
)

// Options agent client options.
type Options struct {
	Addr    string
	Timeout time.Duration
}

// Client is a client to call agent api.
type Client struct {
	hc httpclient.Client
}

// New returns an instance of agent client.
func New(opts Options) *Client {
	return &Client{
		hc: httpclient.Client{
			Addr: opts.Addr,
			Client: http.Client{
				Timeout: opts.Timeout,
			},
		},
	}
}

// WatchDomains subscribes the given domains.
func (p *Client) WatchDomains(ctx context.Context, domains []string, ttl time.Duration, async bool) error {
	req := _WatchDomainsReq{
		Domains:      domains,
		TTL:          timeutil.Duration(ttl),
		Asynchronous: async,
	}
	return p.hc.Post(ctx, "/sns/agent/api/v1/watch/domains", nil, req, nil)
}
