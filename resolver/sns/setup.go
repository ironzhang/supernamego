package sns

import (
	"context"
	"fmt"

	"github.com/ironzhang/tlog"

	"github.com/ironzhang/supernamego/core/supername"
)

var supernameResolver = &supername.Resolver{}

// Setup 初始化设置
func Setup(opts Options) error {
	var err error

	// 初始化选项设置默认值
	if err = opts.setupDefaults(); err != nil {
		tlog.Errorw("options setup defaults", "error", err)
		return fmt.Errorf("options setup defaults: %w", err)
	}

	// 构造服务发现解析程序
	supernameResolver = &supername.Resolver{
		RouteContext:     opts.RouteContext,
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
