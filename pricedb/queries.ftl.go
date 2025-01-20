// Code generated by FTL. DO NOT EDIT.
package pricedb

import (
    "context"
    "github.com/alecthomas/types/tuple"
    "github.com/block/ftl/common/reflection"
    "github.com/block/ftl/go-runtime/server"
    stdtime "time"
)

type LoadPricesResult struct {
  	Code string
  	Price float64
  	Timestamp stdtime.Time
  	Currency string
}
type StorePriceQuery struct {
  	Code string
  	Price float64
  	Timestamp stdtime.Time
  	Currency string
}
	
type LoadPricesClient func(context.Context) ([]LoadPricesResult, error)
	
type StorePriceClient func(context.Context, StorePriceQuery) error

func init() {
	reflection.Register(
		server.QuerySource[LoadPricesResult]("pricedb", "loadPrices", reflection.CommandTypeMany, "prices", "SELECT code, price, timestamp, currency FROM prices ORDER BY timestamp DESC", []string{}, []tuple.Pair[string,string]{tuple.PairOf("code", "Code"),tuple.PairOf("price", "Price"),tuple.PairOf("timestamp", "Timestamp"),tuple.PairOf("currency", "Currency")}),
		server.QuerySink[StorePriceQuery]("pricedb", "storePrice", reflection.CommandTypeExec, "prices", "INSERT INTO prices (code, price, timestamp, currency) VALUES (?, ?, ?, ?)", []string{"Code","Price","Timestamp","Currency"}, []tuple.Pair[string,string]{}),
	)
}