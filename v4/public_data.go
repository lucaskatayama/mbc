package mbc

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// PublicDataService contains all Public Data API operations
type PublicDataService struct {
	client *Client
}

func (s *PublicDataService) normalizeInstrumentSymbol(symbol string) (string, error) {
	p := strings.Split(symbol, "-")
	if len(p) != 2 {
		return "", errors.New("instrument symbol should be in the form BASE-QUOTE")
	}
	return strings.ToUpper(symbol), nil
}

func (s *PublicDataService) normalizeInstrumentSymbols(symbols []string) ([]string, error) {
	var ret []string
	for _, y := range symbols {
		nSymbol, err := s.normalizeInstrumentSymbol(y)
		if err != nil {
			return nil, fmt.Errorf("%v: %s", err, y)
		}
		ret = append(ret, nSymbol)
	}
	return ret, nil
}

// Orderbook represents an orderbook
type Orderbook struct {
	Asks [][]float64 `json:"asks"`
	Bids [][]float64 `json:"bids"`
	Ts   UnixTime    `json:"timestamp"`
}

// OrderbookParams represents an orderbook request param
// Symbol is an instrument symbol with format BASE-QUOTE
type OrderbookParams struct {
	Symbol string `url:"-"`
}

// GetOrderbook fetches an orderbook given an instrument symbol
func (s *PublicDataService) GetOrderbook(ctx context.Context, params OrderbookParams, opts ...RequestOpt) (Orderbook, *http.Response, error) {
	var err error
	params.Symbol, err = s.normalizeInstrumentSymbol(params.Symbol)
	if err != nil {
		return Orderbook{}, nil, err
	}

	p := fmt.Sprintf("/%s/orderbook", params.Symbol)
	req, err := s.client.newRequest(ctx, http.MethodGet, p, params, opts)
	if err != nil {
		return Orderbook{}, nil, err
	}
	var book Orderbook
	resp, err := s.client.do(req, &book)
	if err != nil {
		return Orderbook{}, nil, err
	}

	return book, resp, nil
}

// Ticker represents a ticker
type Ticker struct {
	Pair   string   `json:"pair"`
	High   string   `json:"high"`
	Low    string   `json:"low"`
	Volume string   `json:"vol"`
	Last   string   `json:"last"`
	Buy    string   `json:"buy"`
	Sell   string   `json:"sell"`
	Open   string   `json:"open"`
	Ts     UnixTime `json:"date"`
}

// TickerParams a ticker request param
// Symbols represents a list of instrument symbols with format BASE-QUOTE
type TickerParams struct {
	Symbols []string `url:"symbols,comma"`
}

// ListTickers fetches a list of tickers given a list of instrument symbols
func (s *PublicDataService) ListTickers(ctx context.Context, params TickerParams, opts ...RequestOpt) ([]Ticker, error) {
	var err error
	params.Symbols, err = s.normalizeInstrumentSymbols(params.Symbols)
	if err != nil {
		return nil, err
	}

	req, err := s.client.newRequest(ctx, http.MethodGet, "/tickers", params, opts)
	if err != nil {
		return nil, err
	}
	var tickers []Ticker
	if resp, err := s.client.do(req, &tickers); err != nil {
		fmt.Print(resp)
		return nil, err
	}

	return tickers, nil
}

// Trade respresents a single trade operation
type Trade struct {
	Tid    int64    `json:"tid"`
	Ts     UnixTime `json:"date"`
	Type   string   `json:"type"`
	Price  float64  `json:"price"`
	Amount float64  `json:"amount"`
}

// TradeParams represents a ListTrades request param
type TradeParams struct {
	Symbol string    `url:"-"`
	Tid    int64     `url:"tid,omitempty"`
	Since  int64     `url:"since,omitempty"`
	From   time.Time `url:"from,omitempty,unix"`
	To     time.Time `url:"to,omitempty,unix"`
}

// ListTrades fetches trades given an instrument symbol
func (s *PublicDataService) ListTrades(ctx context.Context, params TradeParams, opts ...RequestOpt) ([]Trade, *http.Response, error) {
	var err error
	params.Symbol, err = s.normalizeInstrumentSymbol(params.Symbol)

	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.newRequest(ctx, http.MethodGet, fmt.Sprintf("/%s/trades", params.Symbol), params, opts)
	if err != nil {
		return nil, nil, err
	}

	var data []Trade
	resp, err := s.client.do(req, &data)
	if err != nil {
		fmt.Print(resp)
		return nil, nil, err
	}

	return data, resp, nil
}
