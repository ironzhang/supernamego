package main

import (
	"context"
	"net/http"
	"time"

	"github.com/ironzhang/superlib/fileutil"
	"github.com/ironzhang/superlib/httputils/httpclient"
	"github.com/ironzhang/superlib/httputils/httpclient/interceptors"
	"github.com/ironzhang/supernamego"
	"github.com/ironzhang/tlog"
)

type Config struct {
	Servers []string
}

func LoadConfig() (Config, error) {
	var cfg Config
	err := fileutil.ReadTOML("config.toml", &cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil
}

func SupernameResolve(ctx context.Context, addr string) (string, error) {
	host, _, err := supernamego.Lookup(ctx, addr)
	return host, err
}

func Echo(c *httpclient.Client, in string) (string, error) {
	var out string
	err := c.Post(context.TODO(), "/echo", nil, in, &out)
	if err != nil {
		return "", err
	}
	return out, nil
}

func RunEcho(addr string) {
	c := &httpclient.Client{
		Addr: addr,
		Client: http.Client{
			Timeout: time.Second,
		},
		Resolve: SupernameResolve,
		Interceptors: []httpclient.Interceptor{
			interceptors.AccessLogInterceptor(),
		},
	}

	out, err := Echo(c, "hello, world")
	if err != nil {
		tlog.Errorw("echo", "addr", addr, "error", err)
		return
	}
	tlog.Infof("%s", out)
}

func main() {
	cfg, err := LoadConfig()
	if err != nil {
		tlog.Errorw("load config", "error", err)
		return
	}

	for _, server := range cfg.Servers {
		RunEcho(server)
	}
}
