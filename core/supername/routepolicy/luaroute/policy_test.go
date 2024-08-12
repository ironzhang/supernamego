package luaroute

import (
	"reflect"
	"regexp"
	"testing"

	"github.com/ironzhang/superlib/superutil/supermodel"
)

func matchError(t testing.TB, err error, errstr string) bool {
	switch {
	case err == nil && errstr == "":
		return true
	case err != nil && errstr == "":
		return false
	case err == nil && errstr != "":
		return false
	case err != nil && errstr != "":
		matched, e := regexp.MatchString(errstr, err.Error())
		if e != nil {
			t.Fatalf("match: regexp match string: %v", e)
		}
		return matched
	}
	panic("never reach")
}

func TestPolicy(t *testing.T) {
	p := NewPolicy()

	// 加载脚本
	scripts := []string{"./testdata/test1.lua", "./testdata/test2.lua"}
	for _, script := range scripts {
		if err := p.Load(script); err != nil {
			t.Fatalf("load %q: %v", script, err)
		}
	}

	// 测试集群列表
	clusters := []supermodel.Cluster{
		supermodel.Cluster{
			Name: "az00.default.k8s",
			Labels: map[string]string{
				supermodel.ZoneKey: "az00",
				supermodel.LaneKey: "default",
			},
			Endpoints: make([]supermodel.Endpoint, 10),
		},
		supermodel.Cluster{
			Name: "az00.small.k8s",
			Labels: map[string]string{
				supermodel.ZoneKey: "az00",
				supermodel.LaneKey: "small",
			},
			Endpoints: make([]supermodel.Endpoint, 10),
		},
		supermodel.Cluster{
			Name: "az01.default.k8s",
			Labels: map[string]string{
				supermodel.ZoneKey: "az01",
				supermodel.LaneKey: "default",
			},
			Endpoints: make([]supermodel.Endpoint, 10),
		},
	}

	// 测试用例
	tests := []struct {
		domain       string
		routectx     map[string]string
		clusters     []supermodel.Cluster
		destinations []supermodel.Destination
		err          string
	}{
		{
			domain:       "www.test1.com",
			routectx:     map[string]string{},
			clusters:     clusters,
			destinations: nil,
			err:          "",
		},
		{
			domain: "www.test1.com",
			routectx: map[string]string{
				supermodel.ZoneKey: "az03",
				supermodel.LaneKey: "default",
			},
			clusters:     clusters,
			destinations: nil,
			err:          "",
		},
		{
			domain: "www.test1.com",
			routectx: map[string]string{
				supermodel.ZoneKey: "az00",
				supermodel.LaneKey: "default",
			},
			clusters: clusters,
			destinations: []supermodel.Destination{
				{Cluster: "az00.default.k8s", Percent: 1},
			},
			err: "",
		},
		{
			domain: "www.test1.com",
			routectx: map[string]string{
				supermodel.ZoneKey: "az00",
				supermodel.LaneKey: "small",
			},
			clusters: clusters,
			destinations: []supermodel.Destination{
				{Cluster: "az00.default.k8s", Percent: 1},
			},
			err: "",
		},
		{
			domain: "www.test1.com",
			routectx: map[string]string{
				supermodel.ZoneKey: "az00",
				supermodel.LaneKey: "read",
			},
			clusters: clusters,
			destinations: []supermodel.Destination{
				{Cluster: "az00.default.k8s", Percent: 1},
			},
			err: "",
		},

		{
			domain:       "www.test2.com",
			routectx:     map[string]string{},
			clusters:     clusters,
			destinations: nil,
			err:          "",
		},
		{
			domain: "www.test2.com",
			routectx: map[string]string{
				supermodel.ZoneKey: "az00",
				supermodel.LaneKey: "default",
			},
			clusters: clusters,
			destinations: []supermodel.Destination{
				{Cluster: "az00.default.k8s", Percent: 1},
			},
			err: "",
		},
		{
			domain: "www.test2.com",
			routectx: map[string]string{
				supermodel.ZoneKey: "az00",
				supermodel.LaneKey: "small",
			},
			clusters: clusters,
			destinations: []supermodel.Destination{
				{Cluster: "az00.small.k8s", Percent: 1},
			},
			err: "",
		},
		{
			domain: "www.test2.com",
			routectx: map[string]string{
				supermodel.ZoneKey: "az01",
				supermodel.LaneKey: "small",
			},
			clusters: clusters,
			destinations: []supermodel.Destination{
				{Cluster: "az01.default.k8s", Percent: 1},
			},
			err: "",
		},
	}
	for i, tt := range tests {
		dests, err := p.MatchRoute(tt.domain, tt.routectx, tt.clusters)
		if got, want := err, tt.err; !matchError(t, got, want) {
			t.Fatalf("%d: match route, domain=%q: error is not match, got %v, want %v", i, tt.domain, got, want)
		}
		if err != nil {
			continue
		}
		if got, want := dests, tt.destinations; !reflect.DeepEqual(got, want) {
			t.Fatalf("%d: match route, domain=%q: destinations is unexpected, got %v, want %v", i, tt.domain, got, want)
		}
	}
}
