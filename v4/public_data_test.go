package mbc_test

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/lucaskatayama/mbc/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestPublicDataService_GetOrderbook(t *testing.T) {
	assert := assert.New(t)

	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/BTC-BRL/orderbook",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			w.WriteHeader(http.StatusOK)
			data := mbc.Orderbook{
				Asks: [][]float64{
					{0.1, 100},
				},
				Bids: [][]float64{
					{0.05, 100},
				},
			}
			resp, _ := json.Marshal(data)
			_, _ = w.Write(resp)
		},
	)

	got, resp, err := client.PublicData.GetOrderbook(context.Background(), mbc.OrderbookParams{Symbol: "BTC-BRL"})
	if err != nil {
		t.Fatalf("PublicData.GetOrderbook returned error: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("respnse is not OK/200: %d", resp.StatusCode)
	}
	assert.NotNil(got.Asks, "asks should not be nil")
	assert.NotNil(got.Bids, "bids should not be nil")
}

func TestPublicDataService_GetTickers(t *testing.T) {
	assert := assert.New(t)

	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/tickers",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			w.WriteHeader(http.StatusOK)

			var data []mbc.Ticker
			symbols := strings.Split(r.URL.Query().Get("symbols"), ",")
			for _, s := range symbols {
				data = append(data, mbc.Ticker{
					Pair:   s,
					High:   "1",
					Low:    "0",
					Volume: "10",
					Last:   "10",
					Buy:    "10",
					Sell:   "10",
					Open:   "10",
					Ts:     mbc.UnixTime{time.Now()},
				})
			}

			resp, _ := json.Marshal(data)
			_, _ = w.Write(resp)
		},
	)

	params := mbc.TickerParams{Symbols: []string{"BTC-BRL", "LTC-BRL"}}
	got, err := client.PublicData.ListTickers(context.Background(), params)
	if err != nil {
		t.Fatalf("PublicData.ListTickers returned error: %v", err)
	}
	assert.Equalf(len(params.Symbols), len(got), "should return same length of requested symbols")
	for i, t := range got {
		assert.Equalf(params.Symbols[i], t.Pair, "pair should be equal on index: %d", i)
	}
}

func TestPublicDataService_GetTrades(t *testing.T) {
	assert := assert.New(t)

	mux, server, client := setup(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/BTC-BRL/trades",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			w.WriteHeader(http.StatusOK)
			var data []mbc.Trade

			for i := 0; i < 1000; i++ {
				data = append(data, mbc.Trade{
					Tid:    int64(i),
					Ts:     mbc.UnixTime{Time: time.Now()},
					Type:   "buy",
					Price:  10,
					Amount: 10,
				})
			}
			resp, _ := json.Marshal(data)
			_, _ = w.Write(resp)
		},
	)

	params := mbc.TradeParams{
		Symbol: "BTC-BRL",
	}
	got, resp, err := client.PublicData.ListTrades(context.Background(), params)
	if err != nil {
		t.Fatalf("PublicData.ListTrades returned error: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("response is not OK/200: %d", resp.StatusCode)
	}
	assert.NotNil(got)
}

func TestName(t *testing.T) {
	client, err := mbc.NewPublicOnlyClient()
	if err != nil {
		t.Fatalf("not ok")
	}

	o, resp, err := client.PublicData.GetOrderbook(context.Background(), mbc.OrderbookParams{Symbol: "BTC-BRL"})
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("response is not 200/OK: %d", resp.StatusCode)
	}

	fmt.Println(o)
}
