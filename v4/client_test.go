package mbc_test

import (
	"context"
	"fmt"
	"github.com/lucaskatayama/mbc/v4"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetTickers(t *testing.T) {
	client := mbc.NewREST()

	ticker, err := client.GetTickers(context.Background(), "BTC-BRL")
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	t.Log(ticker)
}

func TestGetOrderbook(t *testing.T) {
	client := mbc.NewREST()

	ticker, err := client.GetOrderbook(context.Background(), "BTC-BRL")
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	t.Log(ticker)
}

func TestWebsocket(t *testing.T) {
	done := make(chan os.Signal)
	client := mbc.NewWS()
	client.HandleMsg(func(msg []byte) {
		t.Log(string(msg))
	})
	client.HandleErr(func(msg []byte) {
		t.Log(string(msg))
	})
	err := client.Subscribe(context.Background(), "ticker", "BRLBTC")
	if err != nil {
		t.Log(err)
	}
	<-done
}

func TestClient_GetTickers(t *testing.T) {
	type args struct {
		instruments []string
	}
	tests := []struct {
		name    string
		args    args
		wantT   []mbc.Ticker
		wantErr assert.ErrorAssertionFunc
	}{
		{
			"Simple Test",
			args{
				instruments: []string{"BTC-BRL"},
			},
			[]mbc.Ticker{},
			func(t assert.TestingT, err error, i ...interface{}) bool {
				return false
			},
		},
		{
			"MB pair",
			args{
				instruments: []string{"BRLBTC"},
			},
			[]mbc.Ticker{},
			func(t assert.TestingT, err error, i ...interface{}) bool {
				return false
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			ctx := context.Background()
			c := mbc.NewREST()
			gotT, err := c.GetTickers(ctx, tt.args.instruments...)
			if tt.wantErr(t, err, fmt.Sprintf("GetTickers(%v)", tt.args.instruments)) {
				return
			}
			assert.IsType(tt.wantT, gotT, "GetTickers(%v)", tt.args.instruments)
			t.Log(gotT)
		})
	}
}
