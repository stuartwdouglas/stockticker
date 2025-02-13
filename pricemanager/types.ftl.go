// Code generated by FTL. DO NOT EDIT.
package pricemanager

import (
	"context"
	ftlpricedb "ftl/pricedb"
	ftlticker "ftl/ticker"
	"github.com/block/ftl/common/reflection"
	"github.com/block/ftl/go-runtime/server"
)

type PricesClient func(context.Context, ftlticker.Price) error

func init() {
	reflection.Register(

		reflection.ProvideResourcesForVerb(
			Prices,
			server.SinkClient[ftlpricedb.SavePriceClient, ftlpricedb.Price](),
		),
	)
}
