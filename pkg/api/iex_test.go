package iex

import (
	"fmt"
	"testing"
)

func TestBatchQuotes(t *testing.T) {
	b, err := BatchQuotes([]string{"amd", "tsla"})
	if err != nil {
		t.Fatal("Failed to get batch quotes: ", err)
	}

	q, err := b.Quote("amd")
	if err != nil {
		t.Fatal("Failed to get quote from batch")
	}

	if q.Symbol != "AMD" {
		t.Error("Symbol was differnt than requested")
	}
}

func TestQuote(t *testing.T) {
	symbol := "WDC"
	info, err := Quote(symbol, false)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(info)
}

func TestPrice(t *testing.T) {
	symbol := "WDC"
	price, err := Price(symbol)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(price)
}

func TestPriceIndex(t *testing.T) {
	symbol := "VOO"
	price, err := Price(symbol)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(price)
}

func TestNews(t *testing.T) {

	news, err := News("amd")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(news)
}

func TestStat(t *testing.T) {
	_, err := Stats("aapl", true)
	if err != nil {
		t.Fatal(err)
	}
}

func TestEarnings(t *testing.T) {
	_, err := Earnings("aapl")
	if err != nil {
		t.Fatal(err)
	}
}

func TestDividends(t *testing.T) {
	_, err := Dividends("aapl", OneYear)
	if err != nil {
		t.Fatal(err)
	}
}
