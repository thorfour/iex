package types

import (
	"fmt"
	"strings"
)

const (
	APIVersion = "1.0"
	QuoteStr   = "quote"
	NewsStr    = "news"
	ChartStr   = "chart"
	StockStr   = "stock"
	PriceStr   = "price"
	BatchStr   = "batch"
	LastStr    = "last"
	MrktStr    = "market"
	APIURL     = "https://api.iextrading.com/"
)

// Quote repesents the format returned for a quote from IEX(https://iextrading.com)
type Quote struct {
	Symbol                string  `json:"symbol"`
	CompanyName           string  `json:"companyName"`
	PrimaryExchange       string  `json:"primaryExchange"`
	Sector                string  `json:"sector"`
	CalculationPrice      string  `json:"calculationPrice"`
	Open                  float64 `json:"open"`
	OpenTime              int64   `json:"openTime"`
	Close                 float64 `json:"close"`
	CloseTime             int64   `json:"closeTime"`
	High                  float64 `json:"high"`
	Low                   float64 `json:"low"`
	LatestPrice           float64 `json:"latestPrice"`
	LatestSource          string  `json:"latestSource"`
	LatestTime            string  `json:"latestTime"`
	LatestUpdate          int64   `json:"latestUpdate"`
	LatestVolume          int     `json:"latestVolume"`
	IexRealtimePrice      float64 `json:"iexRealtimePrice"`
	IexRealtimeSize       int     `json:"iexRealtimeSize"`
	IexLastUpdated        int64   `json:"iexLastUpdated"`
	DelayedPrice          float64 `json:"delayedPrice"`
	DelayedPriceTime      int64   `json:"delayedPriceTime"`
	ExtendedPrice         float64 `json:"extendedPrice"`
	ExtendedChange        float64 `json:"extendedChange"`
	ExtendedChangePercent float64 `json:"extendedChangePercent"`
	ExtendedPriceTime     int64   `json:"extendedPriceTime"`
	PreviousClose         float64 `json:"previousClose"`
	Change                float64 `json:"change"`
	ChangePercent         float64 `json:"changePercent"`
	IexMarketPercent      float64 `json:"iexMarketPercent"`
	IexVolume             int     `json:"iexVolume"`
	AvgTotalVolume        int     `json:"avgTotalVolume"`
	IexBidPrice           int     `json:"iexBidPrice"`
	IexBidSize            int     `json:"iexBidSize"`
	IexAskPrice           int     `json:"iexAskPrice"`
	IexAskSize            int     `json:"iexAskSize"`
	MarketCap             int64   `json:"marketCap"`
	PeRatio               float64 `json:"peRatio"`
	Week52High            float64 `json:"week52High"`
	Week52Low             float64 `json:"week52Low"`
	YtdChange             float64 `json:"ytdChange"`
}

// News is the news structure returned from IEX
type News struct {
	Datetime string `json:"datetime"`
	Headline string `json:"headline"`
	Source   string `json:"source"`
	URL      string `json:"url"`
	Summary  string `json:"summary"`
	Related  string `json:"related"`
	Image    string `json:"image"`
}

// Batch is a []Quote
type Batch map[string]map[string]Quote

// Quote returns the quote in a iex batch for a specific ticker
// returns error if symbol does not exist
func (i Batch) Quote(ticker string) (Quote, error) {
	ticker = strings.ToUpper(ticker)
	t, ok := i[ticker]
	if !ok {
		return Quote{}, fmt.Errorf("Failed to find %v in batch request", ticker)
	}

	q, ok := t[QuoteStr]
	if !ok {
		return Quote{}, fmt.Errorf("Failed to find quote for %v in batch request", ticker)
	}

	return q, nil
}
