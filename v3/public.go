package mbc

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// Ticker retrieves ticker for base/quote
func (c Client) Ticker(ctx context.Context, base string, quote string) (Ticker, error) {
	u := fmt.Sprintf("%s/api/%s/ticker", endpoint, strings.ToUpper(base))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return Ticker{}, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return Ticker{}, err
	}
	defer res.Body.Close()

	var tickerD tickerD
	if err := json.NewDecoder(res.Body).Decode(&tickerD); err != nil {
		return Ticker{}, err
	}
	return tickerD.Ticker, err
}

// Orderbook retrieves ticker for base/quote
func (c Client) Orderbook(ctx context.Context, base string, quote string, opts ...OrderbookOption) (Orderbook, error) {

	bla := url.Values{}
	for _, opt := range opts {
		opt(bla)
	}

	endp := fmt.Sprintf("%s/api/%s/orderbook?%s", endpoint, strings.ToUpper(base), bla.Encode())

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endp, nil)
	if err != nil {
		return Orderbook{}, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return Orderbook{}, err
	}
	defer res.Body.Close()

	var orderbookD Orderbook
	if err := json.NewDecoder(res.Body).Decode(&orderbookD); err != nil {
		return Orderbook{}, err
	}
	return orderbookD, err
}

// Trades lists trades
func (c Client) Trades(ctx context.Context, base string, quote string, opts ...TradesOption) ([]Trade, error) {

	bla := url.Values{}
	for _, opt := range opts {
		opt(bla)
	}

	endp := fmt.Sprintf("%s/api/%s/trades?%s", endpoint, strings.ToUpper(base), bla.Encode())

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endp, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var trades []Trade
	if err := json.NewDecoder(res.Body).Decode(&trades); err != nil {
		return nil, err
	}
	return trades, err
}

// Coins list coins
func (c Client) Coins(ctx context.Context) ([]string, error) {
	endp := fmt.Sprintf("%s/api/coins", endpoint)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endp, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var d []string
	if err := json.NewDecoder(res.Body).Decode(&d); err != nil {
		return nil, err
	}
	return d, err
}
