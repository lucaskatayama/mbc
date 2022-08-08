package mbc_test

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/lucaskatayama/mbc/v4"
	"github.com/lucaskatayama/mbc/v4/utils"

	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestPublicDataService_GetTickers(t *testing.T) {
	assert := assert.New(t)

	mux, server, client := setupPublicOnly(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/tickers",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			w.WriteHeader(http.StatusOK)

			var data []mbc.Ticker
			symbols := strings.Split(r.URL.Query().Get("symbols"), ",")
			for _, s := range symbols {
				data = append(data, mbc.Ticker{
					Pair:   mbc.InstrumentSymbol(s),
					High:   "1",
					Low:    "0",
					Volume: "10",
					Last:   "10",
					Buy:    "10",
					Sell:   "10",
					Open:   "10",
					Ts:     utils.UnixTime(time.Now()),
				})
			}

			resp, _ := json.Marshal(data)
			_, _ = w.Write(resp)
		},
	)

	params := mbc.TickerParams{Symbols: []mbc.InstrumentSymbol{"BTC-BRL", "LTC-BRL"}}
	got, _, err := client.PublicData.ListTickers(context.Background(), params)
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

	mux, server, client := setupPublicOnly(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/BTC-BRL/trades",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			w.WriteHeader(http.StatusOK)
			var data []mbc.Trade

			for i := 0; i < 1000; i++ {
				data = append(data, mbc.Trade{
					Tid:    int64(i),
					Ts:     utils.UnixTime(time.Now()),
					Type:   "buy",
					Price:  "10",
					Amount: "10",
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

func ExampleNewClient() {
	client, err := mbc.NewClient("id", "secret")
	if err != nil {
		panic(err)
	}

	fmt.Println(client)
}

func ExampleNewPublicOnlyClient() {
	client, err := mbc.NewPublicOnlyClient()
	if err != nil {
		panic(err)
	}
	client.PublicData.GetOrderbook(context.Background(), mbc.OrderbookParams{Symbol: mbc.BTCBRL})
}

func ExampleNewPublicOnlyClient_getOrderbook() {
	client, err := mbc.NewPublicOnlyClient()
	if err != nil {
		panic(err)
	}
	o, resp, err := client.PublicData.GetOrderbook(context.Background(), mbc.OrderbookParams{Symbol: mbc.BTCBRL})
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusOK {
		panic("not ok")
	}
	fmt.Println(o)
}

func ExampleNewPublicOnlyClient_listTickers() {
	client, err := mbc.NewPublicOnlyClient()
	if err != nil {
		panic(err)
	}
	o, resp, err := client.PublicData.ListTickers(context.Background(), mbc.TickerParams{Symbols: []mbc.InstrumentSymbol{mbc.BTCBRL}})
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusOK {
		panic("not ok")
	}
	fmt.Println(o)
}

func ExampleNewPublicOnlyClient_listTrades() {
	client, err := mbc.NewPublicOnlyClient()
	if err != nil {
		panic(err)
	}
	o, resp, err := client.PublicData.ListTrades(context.Background(), mbc.TradeParams{Symbol: mbc.BTCBRL})
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusOK {
		panic("not ok")
	}
	fmt.Println(o)
}

func ExampleNewPublicOnlyClient_withLogrusLogger() {
	log := logrus.New()
	log.SetLevel(logrus.FatalLevel)
	client, _ := mbc.NewPublicOnlyClient(mbc.WithLogger(log))

	client.PublicData.GetOrderbook(context.Background(), mbc.OrderbookParams{Symbol: mbc.BTCBRL})
}
