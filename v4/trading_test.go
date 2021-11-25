package mbc_test

import (
	"context"
	"github.com/lucaskatayama/mbc/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestTradingService_PlaceOrder(t *testing.T) {
	assert := assert.New(t)

	mux, server, client := setup(t, "id", "secret")
	defer teardown(server)

	mux.HandleFunc("/api/v4/accounts/1/LTC-BRL/orders",
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}
			testMethod(t, r, http.MethodPost)
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"orderId": "123"}`))
		},
	)

	got, resp, err := client.Trading.PlaceOrder(context.Background(), "1", mbc.LTC_BRL, mbc.OrderRequest{
		Type:       mbc.Market,
		Side:       mbc.Buy,
		Cost:       0,
		LimitPrice: 0,
		Qty:        0,
	})
	if err != nil {
		t.Fatalf("Trading.PlaceOrder returned error: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("respnse is not OK/200: %d", resp.StatusCode)
	}
	assert.Equalf(got, mbc.OrderID("123"), "should be 123")
}
