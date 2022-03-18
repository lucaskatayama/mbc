package mbc

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/shopspring/decimal"

	"github.com/lucaskatayama/mbc/v4/utils"
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
	Asks [][]decimal.Decimal `json:"asks"`
	Bids [][]decimal.Decimal `json:"bids"`
	Ts   utils.UnixTime      `json:"timestamp"`
}

// OrderbookParams represents an orderbook request param
// InstrumentSymbol is an instrument symbol with format BASE-QUOTE
type OrderbookParams struct {
	Symbol InstrumentSymbol `url:"-"`
	Limit  int64            `url:"limit"`
}

// GetOrderbook fetches an orderbook given an instrument symbol
func (s *PublicDataService) GetOrderbook(ctx context.Context, params OrderbookParams, opts ...RequestOpt) (Orderbook, *http.Response, error) {
	var err error

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
	Pair   InstrumentSymbol `json:"pair"`
	High   string           `json:"high"`
	Low    string           `json:"low"`
	Volume string           `json:"vol"`
	Last   string           `json:"last"`
	Buy    string           `json:"buy"`
	Sell   string           `json:"sell"`
	Open   string           `json:"open"`
	Ts     utils.UnixTime   `json:"date"`
}

// TickerParams a ticker request param
// Symbols represents a list of instrument symbols with format BASE-QUOTE
type TickerParams struct {
	Symbols []InstrumentSymbol `url:"symbols,comma"`
}

// ListTickers fetches a list of tickers given a list of instrument symbols
func (s *PublicDataService) ListTickers(ctx context.Context, params TickerParams, opts ...RequestOpt) ([]Ticker, *http.Response, error) {
	var err error
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.newRequest(ctx, http.MethodGet, "/tickers", params, opts)
	if err != nil {
		return nil, nil, err
	}
	var tickers []Ticker
	resp, err := s.client.do(req, &tickers)
	if err != nil {
		return nil, nil, err
	}

	return tickers, resp, nil
}

// Trade respresents a single trade operation
type Trade struct {
	Tid    int64          `json:"tid"`
	Ts     utils.UnixTime `json:"date"`
	Type   string         `json:"type"`
	Price  string         `json:"price"`
	Amount string         `json:"amount"`
}

// TradeParams represents a ListTrades request param
type TradeParams struct {
	Symbol InstrumentSymbol `url:"-"`
	Tid    int64            `url:"tid,omitempty"`
	Since  int64            `url:"since,omitempty"`
	From   time.Time        `url:"from,omitempty,unix"`
	To     time.Time        `url:"to,omitempty,unix"`
}

// ListTrades fetches trades given an instrument symbol
func (s *PublicDataService) ListTrades(ctx context.Context, params TradeParams, opts ...RequestOpt) ([]Trade, *http.Response, error) {
	var err error

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
