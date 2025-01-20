package nasdaq

import (
	"context"
	"math/rand/v2"

	"github.com/block/ftl/go-runtime/ftl" // Import the FTL SDK.
)

type Price struct {
	Code  string
	Price float64
}

//ftl:topic export partitions=1
type Prices = ftl.TopicHandle[Price, ftl.SinglePartitionMap[Price]]

//ftl:cron 3s
func NasdaqTick(ctx context.Context, prices Prices) error {
	return prices.Publish(ctx, Price{Code: "FOO", Price: rand.Float64() * 100})
}
