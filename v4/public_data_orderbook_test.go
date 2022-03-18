package mbc_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lucaskatayama/mbc/v4"
)

func TestOrderbook_Orderbook(t *testing.T) {

	client, err := mbc.NewPublicOnlyClient()
	if err != nil {
		panic(err)
	}

	params := []mbc.OrderbookParams{
		{Symbol: mbc.BTCBRL},
		{Symbol: mbc.BTCBRL, Limit: 2},
	}

	for _, param := range params {
		t.Run(fmt.Sprintf("%s/%d", param.Symbol, param.Limit), func(t *testing.T) {
			assert := assert.New(t)
			o, resp, err := client.PublicData.GetOrderbook(context.Background(), param)
			if err != nil {
				panic(err)
			}

			assert.Equalf(http.StatusOK, resp.StatusCode, "status should be OK 200")
			assert.NotEmptyf(o.Asks, "asks should not be empty")
			assert.NotEmptyf(o.Bids, "bids should not be empty")
			if param.Limit > 0 {
				assert.Len(o.Asks, int(param.Limit))
				assert.Len(o.Bids, int(param.Limit))
			}
		})
	}
}

func TestPublicDataService_GetOrderbook(t *testing.T) {
	assert := assert.New(t)

	mux, server, client := setupPublicOnly(t)
	defer teardown(server)

	mux.HandleFunc("/api/v4/BTC-BRL/orderbook",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			w.WriteHeader(http.StatusOK)
			data := mbc.Orderbook{
				Asks: [][]string{
					{"0.1", "100"},
				},
				Bids: [][]string{
					{"0.05", "100"},
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
