package iex

import (
	"fmt"
	"strings"
)

type apiString string

func api() apiString {
	return apiURL + IEXAPIVersion + "/"
}

func (s apiString) market() apiString {
	return apiString(string(s) + typeMrktStr + "/")
}

func (s apiString) stock() apiString {
	return apiString(string(s) + typeStockStr + "/")
}

func (s apiString) quote() apiString {
	return apiString(string(s) + typeQuoteStr + "/")
}

func (s apiString) ticker(t string) apiString {
	return apiString(string(s) + t + "/")
}

func (s apiString) price() apiString {
	return apiString(string(s) + typePriceStr + "/")
}

func (s apiString) batch() apiString {
	return apiString(string(s) + typeBatchStr + "?")
}

func (s apiString) tickers(t []string) apiString {
	return apiString(string(s) + strings.Join(t, ","))
}

func (s apiString) symbols() apiString {
	return apiString(string(s) + "symbols=")
}

func (s apiString) and() apiString {
	return apiString(string(s) + "&")
}

func (s apiString) types(t ...string) apiString {
	return apiString(string(s) + "types=" + strings.Join(t, ","))
}

func (s apiString) news() apiString {
	return apiString(string(s) + typeNewsStr + "/")
}

func (s apiString) last() apiString {
	return apiString(string(s) + typeLastStr + "/")
}

func (s apiString) integer(a int) apiString {
	return apiString(string(s) + fmt.Sprintf("%v", a) + "/")
}

func (s apiString) String() string {
	return string(s)
}
