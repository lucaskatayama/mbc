package mbc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type PublicDataService struct {
	client *Client
}

type Orderbook struct {
	Asks [][]float64 `json:"asks"`
	Bids [][]float64 `json:"bids"`
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

type OrderbookParams struct {
	Symbol string `url:"-"`
}

func (s *PublicDataService) GetOrderbook(ctx context.Context, params OrderbookParams, opts ...RequestOpt) (Orderbook, error) {
	var err error
	params.Symbol, err = s.normalizeInstrumentSymbol(params.Symbol)
	if err != nil {
		return Orderbook{}, err
	}

	p := fmt.Sprintf("/%s/orderbook", params.Symbol)
	req, err := s.client.NewRequest(ctx, http.MethodGet, p, params, opts)
	if err != nil {
		return Orderbook{}, err
	}
	var book Orderbook
	if _, err := s.client.Do(req, &book); err != nil {
		return Orderbook{}, err
	}

	return book, nil
}

type Ticker struct {
	Pair   string `json:"pair"`
	High   string `json:"high"`
	Low    string `json:"low"`
	Volume string `json:"vol"`
	Last   string `json:"last"`
	Buy    string `json:"buy"`
	Sell   string `json:"sell"`
	Open   string `json:"open"`
	Date   int64  `json:"date"`
}

type TickerParam struct {
	Symbols []string `url:"symbols,comma"`
}

func (s *PublicDataService) ListTickers(ctx context.Context, params TickerParam, opts ...RequestOpt) ([]Ticker, error) {
	var err error
	params.Symbols, err = s.normalizeInstrumentSymbols(params.Symbols)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, "/tickers", params, opts)
	if err != nil {
		return nil, err
	}
	var tickers []Ticker
	if resp, err := s.client.Do(req, &tickers); err != nil {
		fmt.Print(resp)
		return nil, err
	}

	return tickers, nil
}

type UnixTime struct {
	time.Time
}

func (u *UnixTime) UnmarshalJSON(bytes []byte) error {
	var unix int64
	if err := json.Unmarshal(bytes, &unix); err != nil {
		return err
	}
	u.Time = time.Unix(unix, 0)
	return nil
}

type Trade struct {
	Tid    int64    `json:"tid"`
	Date   UnixTime `json:"date"`
	Type   string   `json:"type"`
	Price  float64  `json:"price"`
	Amount float64  `json:"amount"`
}

func (t Trade) String() string {
	b, _ := json.Marshal(t)
	return string(b)
}

type TradeParams struct {
	Symbol string    `url:"-"`
	Tid    int64     `url:"tid,omitempty"`
	Since  int64     `url:"since,omitempty"`
	From   time.Time `url:"from,omitempty,unix"`
	To     time.Time `url:"to,omitempty,unix"`
}

func (s *PublicDataService) ListTrades(ctx context.Context, params TradeParams, opts ...RequestOpt) ([]Trade, error) {
	var err error
	params.Symbol, err = s.normalizeInstrumentSymbol(params.Symbol)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, fmt.Sprintf("/%s/trades", params.Symbol), params, opts)
	if err != nil {
		return nil, err
	}
	var data []Trade
	if resp, err := s.client.Do(req, &data); err != nil {
		fmt.Print(resp)
		return nil, err
	}

	return data, nil
}
