package sns

import (
	"fmt"

	"github.com/ironzhang/superlib/fileutil"
)

// Misc 杂项信息
type Misc struct {
	// 忽略预加载错误
	SkipPreloadError bool
}

// Options 初始化选项
type Options struct {
	// 路由标签，为 nil 则读取文件 superoptions/route-params.json
	RouteParams map[string]string

	// 预加载域名列表，为 nil 则读取文件 superoptions/preload.json
	PreloadDomains []string

	// 杂项信息，为 nil 则读取文件 superoptions/misc.json
	Misc *Misc
}

func (p *Options) setupDefaults() (err error) {
	if p.RouteParams == nil {
		p.RouteParams, err = readRouteParams()
		if err != nil {
			return fmt.Errorf("read route params: %w", err)
		}
	}

	if p.PreloadDomains == nil {
		p.PreloadDomains, err = readPreloadDomains()
		if err != nil {
			return fmt.Errorf("read preload domains: %w", err)
		}
	}

	if p.Misc == nil {
		p.Misc, err = readMisc()
		if err != nil {
			return fmt.Errorf("read misc: %w", err)
		}
	}

	return nil
}

func readRouteParams() (map[string]string, error) {
	var params map[string]string
	const path = "superoptions/route-params.json"
	if fileutil.FileExist(path) {
		err := fileutil.ReadJSON(path, &params)
		if err != nil {
			return nil, err
		}
	}
	return params, nil
}

func readPreloadDomains() (domains []string, err error) {
	const path = "superoptions/preload.json"
	if fileutil.FileExist(path) {
		err = fileutil.ReadJSON(path, &domains)
		if err != nil {
			return nil, err
		}
	}
	return domains, nil
}

func readMisc() (*Misc, error) {
	var misc Misc
	const path = "superoptions/misc.json"
	if fileutil.FileExist(path) {
		err := fileutil.ReadJSON(path, &misc)
		if err != nil {
			return nil, err
		}
	}
	return &misc, nil
}
