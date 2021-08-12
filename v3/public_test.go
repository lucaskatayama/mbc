package mbc

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestClient_Ticker(t *testing.T) {
	type args struct {
		ctx   context.Context
		base  string
		quote string
	}
	tests := []struct {
		name    string
		args    args
		want    Ticker
		wantErr bool
	}{
		{
			"Simple test",
			args{
				base:  "BTC",
				quote: "BRL",
			},
			Ticker{
				High: "",
				Low:  "",
				Vol:  "",
				Last: "",
				Buy:  "",
				Sell: "",
				Open: "",
				Date: 0,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New(os.Getenv("MB_ID"), os.Getenv("MB_SECRET"))
			got, err := c.Ticker(context.Background(), tt.args.base, tt.args.quote)
			if (err != nil) != tt.wantErr {
				t.Errorf("Ticker() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Ticker() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Orderbook(t *testing.T) {
	client := New(os.Getenv("MB_ID"), os.Getenv("MB_SECRET"))
	o, _ := client.Orderbook(context.Background(), "btc", "brl", WithLimit(1))
	fmt.Println(o)
}

func ExampleClient_Orderbook() {
	client := New("<ID>", "<SECRET>")
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
	client := New("<ID>", "<SECRET>")
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
	client := New("<ID>", "<SECRET>")
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
