package mbc_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lucaskatayama/mbc/v4"
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

	got, resp, err := client.Trading.PlaceOrder(context.Background(), "1", mbc.LTCBRL, mbc.OrderRequest{
		Type: mbc.Market,
		Side: mbc.Buy,
	})
	if err != nil {
		t.Fatalf("Trading.PlaceOrder returned error: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("respnse is not OK/200: %d", resp.StatusCode)
	}
	assert.Equalf(got, mbc.OrderID("123"), "should be 123")
}
