package supernamego

import (
	"context"
	"fmt"

	"github.com/ironzhang/tlog"

	"github.com/ironzhang/supernamego/supername"
)

var supernameResolver = &supername.Resolver{}

// Setup 初始化设置
func Setup(opts Options) (err error) {
	// 初始化选项设置默认值
	if err = opts.setupDefaults(); err != nil {
		tlog.Errorw("options setup defaults", "error", err)
		return fmt.Errorf("options setup defaults: %w", err)
	}

	// 构造服务发现解析程序
	supernameResolver = &supername.Resolver{
		Tags:             opts.Tags,
		LoadBalancer:     opts.LoadBalancer,
		SkipPreloadError: opts.Misc.SkipPreloadError,
	}

	// 预加载域名
	if len(opts.PreloadDomains) > 0 {
		err = supernameResolver.Preload(context.Background(), opts.PreloadDomains)
		if err != nil {
			tlog.Errorw("supername resolver preload", "domains", opts.PreloadDomains, "error", err)
			return fmt.Errorf("supername resolver preload: %w", err)
		}
	}
	return nil
}

// AutoSetup 无参初始化
func AutoSetup() error {
	return Setup(Options{})
}

// WithLoadBalancer 构建一个新的服务发现解析程序，并重置负载均衡器
func WithLoadBalancer(lb supername.LoadBalancer) *supername.Resolver {
	return supernameResolver.WithLoadBalancer(lb)
}

// Lookup 查找地址节点
func Lookup(ctx context.Context, domain string, tags map[string]string) (addr, cluster string, err error) {
	return supernameResolver.Lookup(ctx, domain, tags)
}
