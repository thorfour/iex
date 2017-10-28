package iex

const (
	// Current IEX API version
	IEXAPIVersion = "1.0"
	typeQuoteStr  = "quote"
	typeNewsStr   = "news"
	typeChartStr  = "chart"
	typeStockStr  = "stock"
	typePriceStr  = "price"
	typeBatchStr  = "batch"
	typeLastStr   = "last"
	typeMrktStr   = "market"
	apiURL        = "https://api.iextrading.com/"
)

// IEXQuote repesents the format returned for a quote from IEX(https://iextrading.com)
type IEXQuote struct {
	Symbol           string  `json:symbol`
	CompanyName      string  `json:companyName`
	PrimaryExchange  string  `json:primaryExchange`
	CalculationPrice string  `json:calculationPrice`
	IexRealtimePrice float64 `json:iexRealtimePrice`
	IexRealtimeSize  float64 `json:iexRealtimeSize`
	IexLastUpdated   float64 `json:iexLastUpdated`
	DelayedPrice     float64 `json:delayedPrice`
	DelayedPriceTime float64 `json:delayedPriceTime`
	PreviousClose    float64 `json:previousClose`
	Change           float64 `json:change`
	ChangePercent    float64 `json:changePercent`
	IexMarketPercent float64 `json:iexMarketPercent`
	IexVolume        float64 `json:iexVolume`
	AvgTotalVolume   float64 `json:avgTotalVolume`
	IexBidPrice      float64 `json:iexBidPrice`
	IexBidSize       float64 `json:iexBidSize`
	IexAskPrice      float64 `json:iexAskPrice`
	IexAskSize       float64 `json:iexAskSize`
	MarketCap        float64 `json:marketCap`
	//PeRatio          float64 `json:peRatio`
	Week52High float64 `json:week52High`
	Week52Low  float64 `json:week52Low`
}

type IEXNews struct {
	DateTime string `json:datetime`
	Headline string `json:headline`
	Source   string `json:source`
	URL      string `json:url`
	Summary  string `json:summar`
	Related  string `json:related`
}

// IEXBatch is a []IEXQuote
type IEXBatch map[string]map[string]IEXQuote
