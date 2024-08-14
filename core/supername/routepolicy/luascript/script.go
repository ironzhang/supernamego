package luascript

import (
	"errors"
	"fmt"
	"sync"

	lua "github.com/yuin/gopher-lua"

	"github.com/ironzhang/superlib/superutil/supermodel"
)

// Script 路由脚本
type Script struct {
	mu     sync.Mutex
	lstate *lua.LState
}

// NewScript 构建路由脚本
func NewScript() *Script {
	return &Script{lstate: lua.NewState()}
}

// Load 加载路由脚本
func (p *Script) Load(path string) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.load(path)
}

// Matches 执行路由匹配
func (p *Script) Matches(domain string, rctx map[string]string, clusters []supermodel.Cluster) ([]supermodel.Destination, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.matches(domain, rctx, clusters)
}

func (p *Script) load(path string) error {
	err := p.lstate.DoFile(path)
	if err != nil {
		return fmt.Errorf("luaroute: do file %q: %w", path, err)
	}
	return nil
}

func (p *Script) matches(domain string, rctx map[string]string, clusters []supermodel.Cluster) ([]supermodel.Destination, error) {
	// 查找脚本函数
	fn, err := p.lookupFunction("MatchFuncs", domain)
	if err != nil {
		return nil, fmt.Errorf("luaroute: lookup function: %w", err)
	}

	// 调用脚本函数
	err = p.lstate.CallByParam(lua.P{
		Fn:      fn,
		NRet:    1,
		Protect: true,
	}, makeMapTable(p.lstate, rctx), makeClusterSliceTable(p.lstate, clusters))
	if err != nil {
		return nil, fmt.Errorf("luaroute: call by param: %w", err)
	}

	// 获取脚本函数调用返回值
	ret := p.lstate.Get(-1)
	p.lstate.Pop(1)
	lt, ok := ret.(*lua.LTable)
	if !ok {
		return nil, errors.New("luaroute: pop value is not a table")
	}

	// 构建目标节点
	dests, err := makeDestinations(lt)
	if err != nil {
		return nil, fmt.Errorf("luaroute: make destinations: %w", err)
	}
	return dests, nil
}

func (p *Script) lookupFunction(table, key string) (*lua.LFunction, error) {
	lt, ok := p.lstate.GetGlobal(table).(*lua.LTable)
	if !ok {
		return nil, fmt.Errorf("%q is not a table", table)
	}
	lf, ok := lt.RawGetString(key).(*lua.LFunction)
	if !ok {
		return nil, fmt.Errorf("can not find function %s[%q]", table, key)
	}
	return lf, nil
}
