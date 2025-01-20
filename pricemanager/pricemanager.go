package pricemanager

import (
	"context"
	"time"

	"github.com/block/ftl/go-runtime/ftl"

	"ftl/nasdaq"
	"ftl/pricedb"
)

//ftl:subscribe nasdaq.prices from=latest
func Prices(ctx context.Context, in nasdaq.Price, savePrice pricedb.SavePriceClient) error {
	ftl.LoggerFromContext(ctx).Infof("Received price: %v", in)
	return savePrice(ctx, pricedb.Price{
		Code:     in.Code,
		Price:    in.Price,
		Time:     time.Now(),
		Currency: "USD",
	})
}
