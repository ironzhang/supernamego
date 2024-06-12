package supername

import (
	"sync/atomic"

	"github.com/ironzhang/superlib/superutil/supermodel"
)

type provider struct {
	domain  string
	service atomic.Value // *model.ServiceModel
	route   atomic.Value // *model.RouteModel
}

func (p *provider) StoreServiceModel(s *supermodel.ServiceModel) {
	p.service.Store(s)
}

func (p *provider) LoadServiceModel() (*supermodel.ServiceModel, bool) {
	s, ok := p.service.Load().(*supermodel.ServiceModel)
	return s, ok
}

func (p *provider) StoreRouteModel(r *supermodel.RouteModel) {
	p.route.Store(r)
}

func (p *provider) LoadRouteModel() (*supermodel.RouteModel, bool) {
	r, ok := p.route.Load().(*supermodel.RouteModel)
	if !ok {
		return &supermodel.RouteModel{Domain: p.domain}, true
	}
	return r, ok
}
