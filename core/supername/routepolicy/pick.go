package routepolicy

import (
	"time"

	"math/rand"

	"github.com/ironzhang/superlib/superutil/supermodel"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func pick(dests []supermodel.Destination) (cluster string) {
	sum := 0.0
	r := rand.Float64()
	for _, dest := range dests {
		sum += dest.Percent
		if r < sum {
			return dest.Cluster
		}
	}
	if len(dests) > 0 {
		return dests[0].Cluster
	}
	return ""
}
