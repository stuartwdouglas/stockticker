package pricedb

import (
	"context"
	"time"

	"github.com/block/ftl/go-runtime/ftl" // Import the FTL SDK.
)

type Price struct {
	Code     string
	Price    float64
	Time     time.Time
	Currency string
}

type PriceDatasource struct {
	ftl.DefaultMySQLDatabaseConfig
}

func (PriceDatasource) Name() string { return "prices" }

//ftl:verb export
func SavePrice(ctx context.Context, price Price, client StorePriceClient) error {
	return client(ctx, StorePriceQuery{
		Code:      price.Code,
		Price:     price.Price,
		Timestamp: price.Time,
		Currency:  "USD",
	})
}

//ftl:verb export
func QueryPrices(ctx context.Context, client LoadPricesClient) ([]Price, error) {
	prices, err := client(ctx)
	if err != nil {
		return nil, err
	}
	var items []Price
	for _, i := range prices {
		items = append(items, Price{
			Code:     i.Code,
			Price:    i.Price,
			Time:     i.Timestamp,
			Currency: i.Currency,
		})
	}
	return items, nil
}
