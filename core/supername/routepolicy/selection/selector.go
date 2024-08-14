package selection

import "github.com/ironzhang/superlib/superutil/supermodel"

// Selector 选择器
type Selector struct {
	requirements []supermodel.Requirement
}

// NewSelector 构造选择器
func NewSelector(reqs ...supermodel.Requirement) *Selector {
	return &Selector{
		requirements: reqs,
	}
}

// AddRequirements 添加匹配条件
func (p *Selector) AddRequirements(reqs ...supermodel.Requirement) {
	p.requirements = append(p.requirements, reqs...)
}

// Matches 依据条件执行匹配
func (p *Selector) Matches(s Store) bool {
	for _, r := range p.requirements {
		if !matches(s, r) {
			return false
		}
	}
	return true
}
