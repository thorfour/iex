package types

import (
	"fmt"
	"strings"
)

const (
	APIVersion  = "1.0"
	QuoteStr    = "quote"
	NewsStr     = "news"
	ChartStr    = "chart"
	StockStr    = "stock"
	PriceStr    = "price"
	BatchStr    = "batch"
	LastStr     = "last"
	MrktStr     = "market"
	StatsStr    = "stats"
	EarningsStr = "earnings"
	APIURL      = "https://api.iextrading.com/"
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

// Stat is that stats structure returned from IEX
type Stat struct {
	CompanyName         string      `json:"companyName"`
	Marketcap           int64       `json:"marketcap"`
	Beta                float64     `json:"beta"`
	Week52High          float64     `json:"week52high"`
	Week52Low           float64     `json:"week52low"`
	Week52Change        float64     `json:"week52change"`
	ShortInterest       int         `json:"shortInterest"`
	ShortDate           string      `json:"shortDate"`
	DividendRate        float64     `json:"dividendRate"`
	DividendYield       float64     `json:"dividendYield"`
	ExDividendDate      string      `json:"exDividendDate"`
	LatestEPS           int         `json:"latestEPS"`
	LatestEPSDate       string      `json:"latestEPSDate"`
	SharesOutstanding   int64       `json:"sharesOutstanding"`
	Float               int64       `json:"float"`
	ReturnOnEquity      float64     `json:"returnOnEquity"`
	ConsensusEPS        float64     `json:"consensusEPS"`
	NumberOfEstimates   int         `json:"numberOfEstimates"`
	EPSSurpriseDollar   interface{} `json:"EPSSurpriseDollar"`
	EPSSurprisePercent  float64     `json:"EPSSurprisePercent"`
	Symbol              string      `json:"symbol"`
	EBITDA              int64       `json:"EBITDA"`
	Revenue             int64       `json:"revenue"`
	GrossProfit         int64       `json:"grossProfit"`
	Cash                int64       `json:"cash"`
	Debt                int64       `json:"debt"`
	TtmEPS              float64     `json:"ttmEPS"`
	RevenuePerShare     int         `json:"revenuePerShare"`
	RevenuePerEmployee  int         `json:"revenuePerEmployee"`
	PeRatioHigh         float64     `json:"peRatioHigh"`
	PeRatioLow          int         `json:"peRatioLow"`
	ReturnOnAssets      float64     `json:"returnOnAssets"`
	ReturnOnCapital     interface{} `json:"returnOnCapital"`
	ProfitMargin        float64     `json:"profitMargin"`
	PriceToSales        float64     `json:"priceToSales"`
	PriceToBook         float64     `json:"priceToBook"`
	Day200MovingAvg     float64     `json:"day200MovingAvg"`
	Day50MovingAvg      float64     `json:"day50MovingAvg"`
	InstitutionPercent  float64     `json:"institutionPercent"`
	InsiderPercent      interface{} `json:"insiderPercent"`
	ShortRatio          float64     `json:"shortRatio"`
	Year5ChangePercent  float64     `json:"year5ChangePercent"`
	Year2ChangePercent  float64     `json:"year2ChangePercent"`
	Year1ChangePercent  float64     `json:"year1ChangePercent"`
	YtdChangePercent    float64     `json:"ytdChangePercent"`
	Month6ChangePercent float64     `json:"month6ChangePercent"`
	Month3ChangePercent float64     `json:"month3ChangePercent"`
	Month1ChangePercent float64     `json:"month1ChangePercent"`
	Day5ChangePercent   float64     `json:"day5ChangePercent"`
	Day30ChangePercent  float64     `json:"day30ChangePercent"`
}

// Earnings is returned from iex/earnings
type Earnings struct {
	Symbol   string `json:"symbol"`
	Earnings []struct {
		ActualEPS              float64 `json:"actualEPS"`
		ConsensusEPS           float64 `json:"consensusEPS"`
		EstimatedEPS           float64 `json:"estimatedEPS"`
		AnnounceTime           string  `json:"announceTime"`
		NumberOfEstimates      int     `json:"numberOfEstimates"`
		EPSSurpriseDollar      float64 `json:"EPSSurpriseDollar"`
		EPSReportDate          string  `json:"EPSReportDate"`
		FiscalPeriod           string  `json:"fiscalPeriod"`
		FiscalEndDate          string  `json:"fiscalEndDate"`
		YearAgo                float64 `json:"yearAgo"`
		YearAgoChangePercent   float64 `json:"yearAgoChangePercent"`
		EstimatedChangePercent float64 `json:"estimatedChangePercent"`
		SymbolID               int     `json:"symbolId"`
	} `json:"earnings"`
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
