package selection

import "strings"

// Store 存储接口
type Store interface {
	GetValues(table, key string) (values []string)
}

// NewStore 构建存储对象
func NewStore(rctx, labels map[string]string) Store {
	return &store{
		rctx:   rctx,
		labels: labels,
	}
}

type store struct {
	rctx   map[string]string
	labels map[string]string
}

func (p *store) GetValues(table, key string) (values []string) {
	var value string
	switch table {
	case "rctx":
		value = p.rctx[key]
	case "labels":
		value = p.labels[key]
	}
	return strings.Split(value, ",")
}
