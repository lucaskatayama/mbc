package mbc

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Client struct {
	host string
}

type ClientOpt func(c *Client)

func NewREST(opts ...ClientOpt) *Client {
	c := &Client{
		host: prodREST,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
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

// GetTickers retrieves tickers
func (c *Client) GetTickers(ctx context.Context, instruments ...string) (t []Ticker, err error) {
	q := url.Values{
		"symbols": instruments,
	}
	u := url.URL{Scheme: "https", Host: c.host, Path: "/api/v4/tickers", RawQuery: q.Encode()}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&t); err != nil {
		return t, err
	}
	return
}

type Orderbook struct {
	Asks [][]float64 `json:"asks"`
	Bids [][]float64 `json:"bids"`
}

func (c *Client) GetOrderbook(ctx context.Context, i string) (o Orderbook, err error) {
	u := url.URL{Scheme: "https", Host: c.host, Path: fmt.Sprintf("/api/v4/%s/orderbook", normalizeInstrument(i))}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return o, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return o, err
	}
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(&o); err != nil {
		return o, err
	}

	return o, err
}
