package endpoint

import (
	"fmt"
	"strings"

	"github.com/thorfour/iex/pkg/types"
)

type APIString string

// API returns the url api endpint
func API() APIString {
	return types.APIURL + types.APIVersion + "/"
}

// Market adds the market type
func (s APIString) Market() APIString {
	return APIString(string(s) + types.MrktStr + "/")
}

// Stock adds the stock type
func (s APIString) Stock() APIString {
	return APIString(string(s) + types.StockStr + "/")
}

// Quote adds the quote type
func (s APIString) Quote() APIString {
	return APIString(string(s) + types.QuoteStr + "/")
}

// Ticker adds the ticker
func (s APIString) Ticker(t string) APIString {
	return APIString(string(s) + t + "/")
}

// Price adds the price type
func (s APIString) Price() APIString {
	return APIString(string(s) + types.PriceStr + "/")
}

// Batch adds the batch type
func (s APIString) Batch() APIString {
	return APIString(string(s) + types.BatchStr + "?")
}

// Tickers adds the tickers as a comma separated list
func (s APIString) Tickers(t []string) APIString {
	return APIString(string(s) + strings.Join(t, ","))
}

// Symbols adds the symbols type
func (s APIString) Symbols() APIString {
	return APIString(string(s) + "symbols=")
}

// And adds the ampersand
func (s APIString) And() APIString {
	return APIString(string(s) + "&")
}

// Types adds the types= argument
func (s APIString) Types(t ...string) APIString {
	return APIString(string(s) + "types=" + strings.Join(t, ","))
}

// News adds the news type
func (s APIString) News() APIString {
	return APIString(string(s) + types.NewsStr + "/")
}

// Last adds the last type
func (s APIString) Last() APIString {
	return APIString(string(s) + types.LastStr + "/")
}

// Integer adds a integer argument
func (s APIString) Integer(a int) APIString {
	return APIString(string(s) + fmt.Sprintf("%v", a) + "/")
}

// String prints the url
func (s APIString) String() string {
	return string(s)
}

// Query appends a query string onto APIString
func (s APIString) Query(args map[string]interface{}) APIString {
	api := string(s) + "?"
	for k, v := range args {
		api = api + k + "=" + fmt.Sprintf("%v", v) + ","
	}

	api = strings.TrimSuffix(api, ",")
	return APIString(api)
}

// Stats adds the stats type
func (s APIString) Stats() APIString {
	return APIString(string(s) + types.StatsStr + "/")
}

// Earnings adds the earnings type
func (s APIString) Earnings() APIString {
	return APIString(string(s) + types.EarningsStr + "/")
}

// Dividends adds the dividends type
func (s APIString) Dividends() APIString {
	return APIString(string(s) + types.DividendsStr + "/")
}

// AddString adds an arbitrary string to the end
func (s APIString) AddString(t string) APIString {
	return APIString(string(s) + t)
}
