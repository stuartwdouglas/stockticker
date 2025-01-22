package pricemanager

import (
	"context"
	"time"

	"github.com/block/ftl/go-runtime/ftl"

	"ftl/pricedb"
	"ftl/ticker"
)

//ftl:subscribe ticker.prices from=latest
func Prices(ctx context.Context, in ticker.Price, savePrice pricedb.SavePriceClient) error {
	ftl.LoggerFromContext(ctx).Infof("Received price: %v", in)
	savePrice(ctx, pricedb.Price{
		Code:     in.Code,
		Price:    in.Price,
		Time:     time.Now(),
		Currency: "USD",
	})
	return nil
}

//ftl:subscribe asxticker.asxPrices from=latest
func AsxPrices(ctx context.Context, in ticker.Price, savePrice pricedb.SavePriceClient) error {
	ftl.LoggerFromContext(ctx).Infof("Received ASX price: %v", in)
	savePrice(ctx, pricedb.Price{
		Code:     in.Code,
		Price:    in.Price,
		Time:     time.Now(),
		Currency: "USD",
	})
	return nil
}
