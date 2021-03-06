package mbc

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClient_Ticker(t *testing.T) {
	assert := assert.New(t)
	client := New()
	ticker, err := client.Ticker(context.Background(), "btc", "brl")
	if err != nil {
		assert.Failf("test failed", "error response: %+v", err)
		return
	}
	assert.NotEmpty(ticker.Sell, "sell should contain data")
	assert.NotEmpty(ticker.Buy, "buy should contain data")
}

func TestClient_Orderbook(t *testing.T) {
	assert := assert.New(t)
	client := New()
	o, err := client.Orderbook(context.Background(), "btc", "brl", WithLimit(1))
	if err != nil {
		assert.Failf("test failed", "error response: %+v", err)
		return
	}
	assert.NotEmpty(o.Asks, "asks should contain data")
	assert.NotEmpty(o.Bids, "bids should contain data")
}

func ExampleClient_Orderbook() {
	client := New()
	o, _ := client.Orderbook(context.Background(), "btc", "brl", WithLimit(1))
	b, _ := json.MarshalIndent(o, "", "  ")
	fmt.Println(string(b))
	// Output:
	// {
	//    "asks": [
	//       [241671.49, 0.01923905],
	//       [241831.86, 0.00205555]
	//    ],
	//    "bids": [
	//       [241664.97001, 0.05772251],
	//       [241664.95, 0.14321774]
	//    ],
	//    "timestamp":1628561698332025456
	// }
}

func ExampleClient_Ticker() {
	client := New()
	o, _ := client.Ticker(context.Background(), "btc", "brl")
	b, _ := json.MarshalIndent(o, "", "  ")
	fmt.Println(string(b))
	// Output:
	// {
	//  "ticker": {
	//    "high": "244900.00000000",
	//    "low": "227539.84013000",
	//    "vol": "158.76695397",
	//    "last": "241671.93011000",
	//    "buy": "241671.93011000",
	//    "sell": "242021.99998000",
	//    "open": "228450.00000000",
	//    "date": 1628563357
	//  }
	// }
}

func ExampleClient_Trades() {
	client := New()
	o, _ := client.Trades(context.Background(), "btc", "brl", FromTid(90000))
	b, _ := json.MarshalIndent(o, "", "  ")
	fmt.Println(string(b))
	// Output:
	// [
	//  {
	//    "tid": 90001,
	//    "date": 1414274533,
	//    "type": "buy",
	//    "price": 919.7789,
	//    "amount": 0.41666409
	//  },
	//  {
	//    "tid": 90002,
	//    "date": 1414274533,
	//    "type": "buy",
	//    "price": 919.9913,
	//    "amount": 0.66149665
	//  }
	// ]
}
