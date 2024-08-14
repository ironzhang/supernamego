package selection_test

import (
	"testing"

	"github.com/ironzhang/superlib/superutil/supermodel"
	"github.com/ironzhang/supernamego/core/supername/routepolicy/selection"
)

func TestSelectorMatches(t *testing.T) {
	requirements1 := []supermodel.Requirement{
		{
			Not:      false,
			Operator: supermodel.Equals,
			Left: supermodel.Token{
				Type:  supermodel.Table,
				Table: "labels",
				Key:   "X-Zone",
			},
			Right: supermodel.Token{
				Type:  supermodel.Table,
				Table: "rctx",
				Key:   "X-Zone",
			},
		},
		{
			Not:      false,
			Operator: supermodel.Equals,
			Left: supermodel.Token{
				Type:  supermodel.Table,
				Table: "labels",
				Key:   "X-Lane",
			},
			Right: supermodel.Token{
				Type:   supermodel.Const,
				Consts: []string{"default"},
			},
		},
		{
			Not:      false,
			Operator: supermodel.Equals,
			Left: supermodel.Token{
				Type:  supermodel.Table,
				Table: "labels",
				Key:   "X-Kind",
			},
			Right: supermodel.Token{
				Type:   supermodel.Const,
				Consts: []string{"k8s"},
			},
		},
	}
	requirements2 := []supermodel.Requirement{
		{
			Not:      false,
			Operator: supermodel.Equals,
			Left: supermodel.Token{
				Type:  supermodel.Table,
				Table: "labels",
				Key:   "X-Zone",
			},
			Right: supermodel.Token{
				Type:   supermodel.Const,
				Consts: []string{"dev"},
			},
		},
		{
			Not:      false,
			Operator: supermodel.Equals,
			Left: supermodel.Token{
				Type:  supermodel.Table,
				Table: "labels",
				Key:   "X-Lane",
			},
			Right: supermodel.Token{
				Type:   supermodel.Const,
				Consts: []string{"default"},
			},
		},
		{
			Not:      false,
			Operator: supermodel.Equals,
			Left: supermodel.Token{
				Type:  supermodel.Table,
				Table: "labels",
				Key:   "X-Kind",
			},
			Right: supermodel.Token{
				Type:   supermodel.Const,
				Consts: []string{"k8s"},
			},
		},
	}

	tests := []struct {
		rctx         map[string]string
		labels       map[string]string
		requirements []supermodel.Requirement
		result       bool
	}{
		{
			rctx: map[string]string{
				"X-Zone": "dev",
			},
			labels: map[string]string{
				"X-Zone": "dev",
				"X-Lane": "default",
				"X-Kind": "k8s",
			},
			requirements: requirements1,
			result:       true,
		},
		{
			rctx: map[string]string{},
			labels: map[string]string{
				"X-Zone": "dev",
				"X-Lane": "default",
				"X-Kind": "k8s",
			},
			requirements: requirements1,
			result:       false,
		},
		{
			rctx: map[string]string{},
			labels: map[string]string{
				"X-Zone": "dev",
				"X-Lane": "default",
				"X-Kind": "k8s",
			},
			requirements: requirements2,
			result:       true,
		},
		{
			rctx: map[string]string{},
			labels: map[string]string{
				"X-Zone": "az00",
				"X-Lane": "default",
				"X-Kind": "k8s",
			},
			requirements: requirements2,
			result:       false,
		},
	}
	for i, tt := range tests {
		store := selection.NewStore(tt.rctx, tt.labels)
		selector := selection.NewSelector(tt.requirements...)
		if got, want := selector.Matches(store), tt.result; got != want {
			t.Fatalf("%d: matches: got %v, want %v", i, got, want)
		}
	}
}
