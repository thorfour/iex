package iex

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// Quote returns a stock quote for a given ticker
func Quote(ticker string) (*IEXQuote, error) {
	url := api().stock().ticker(ticker).quote()
	jsonQuote, err := getJSON(url)
	if err != nil {
		return nil, err
	}

	// Parse into IEX quote
	var quote IEXQuote
	err = json.Unmarshal(jsonQuote, &quote)
	if err != nil {
		return nil, err
	}

	return &quote, nil
}

// Price returns the current price of a ticker
func Price(ticker string) (float64, error) {
	url := api().stock().ticker(ticker).price()
	jsonQuote, err := getJSON(url)
	if err != nil {
		return -1, err
	}

	price, err := strconv.ParseFloat(string(jsonQuote), 64)
	if err != nil {
		return -1, err
	}

	return price, nil
}

// BatchQuotes returns quotes for multiple tickers using a batch request
func BatchQuotes(tickers []string) (IEXBatch, error) {

	url := api().stock().market().batch().symbols().tickers(tickers).and().types(typeQuoteStr)
	jsonQuote, err := getJSON(url)
	if err != nil {
		return nil, err
	}

	// Parse into IEX quote
	var batch IEXBatch
	err = json.Unmarshal(jsonQuote, &batch)
	if err != nil {
		return nil, err
	}

	return batch, nil
}

// Quote returns the quote in a iex batch for a specific ticker
// returns error if symbol does not exist
func (i IEXBatch) Quote(ticker string) (IEXQuote, error) {
	ticker = strings.ToUpper(ticker)
	t, ok := i[ticker]
	if !ok {
		return IEXQuote{}, fmt.Errorf("Failed to find %v in batch request", ticker)
	}

	q, ok := t[typeQuoteStr]
	if !ok {
		return IEXQuote{}, fmt.Errorf("Failed to find quote for %v in batch request", ticker)
	}

	return q, nil
}

// News returns the news for a given symbol
func News(ticker string) ([]IEXNews, error) {

	url := api().stock().ticker(ticker).news().last().integer(5)
	jsonQuote, err := getJSON(url)
	if err != nil {
		return nil, err
	}

	// Parse into IEXNews
	var news []IEXNews
	err = json.Unmarshal(jsonQuote, &news)
	if err != nil {
		return nil, err
	}

	return news, nil
}

// getJSON returns the JSON response from a url
func getJSON(url apiString) ([]byte, error) {

	resp, err := http.Get(url.String())
	if err != nil {
		return nil, err
	}

	// Read the quote into the slice
	defer resp.Body.Close()
	jsonQuote, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return jsonQuote, nil
}
