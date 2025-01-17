package pricedb

import (
	"context"
	"database/sql"
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
func SavePrice(ctx context.Context, price Price, db ftl.DatabaseHandle[PriceDatasource]) error {
	var database *sql.DB = db.Get(ctx) // Get the database connection.
	_, err := database.Exec("INSERT INTO prices (code, price, timestamp, currency) VALUES (?, ?, ?, ?)", price.Code, price.Price, price.Time, price.Currency)
	return err
}

//ftl:verb export
func QueryPrices(ctx context.Context, db ftl.DatabaseHandle[PriceDatasource]) ([]Price, error) {
	var database *sql.DB = db.Get(ctx) // Get the database connection.
	// The following code is standard golang SQL code, it has nothing FTL specific.
	rows, err := database.QueryContext(ctx, "SELECT code, price, timestamp, currency FROM prices")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Price
	for rows.Next() {
		var i Price
		if err := rows.Scan(
			&i.Code,
			&i.Price,
			&i.Time,
			&i.Currency,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
