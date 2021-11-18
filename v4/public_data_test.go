package mbc_test

import (
	"context"
	"fmt"
	"github.com/lucaskatayama/mbc/v4"
	"testing"
	"time"
)

func TestOrderbook_Simple(t *testing.T) {
	c, err := mbc.NewPublicOnlyClient()
	if err != nil {
		t.Fail()
		return
	}

	o, err := c.PublicData.GetOrderbook(context.Background(), mbc.OrderbookParams{Symbol: "BTC-brl"})
	if err != nil {
		t.Fail()
		return
	}
	fmt.Println(o)
}

func TestTicker_Simple(t *testing.T) {
	c, err := mbc.NewPublicOnlyClient()
	if err != nil {
		t.Fail()
		return
	}

	o, err := c.PublicData.ListTickers(context.Background(), mbc.TickerParam{Symbols: []string{"btc-brl", "ltc-BRL"}})
	if err != nil {
		t.Fail()
		return
	}
	fmt.Println(o)
}

func TestTrades_Simple(t *testing.T) {
	c, err := mbc.NewPublicOnlyClient()
	if err != nil {
		t.Fail()
		return
	}

	o, err := c.PublicData.ListTrades(context.Background(), mbc.TradeParams{Symbol: "btc-brl"})
	if err != nil {
		t.Fail()
		return
	}
	fmt.Println(o)
}

func TestTrades_Since(t *testing.T) {
	c, err := mbc.NewPublicOnlyClient()
	if err != nil {
		t.Fail()
		return
	}

	o, err := c.PublicData.ListTrades(context.Background(), mbc.TradeParams{
		Symbol: "btc-brl",
		From:   time.Now().Add(-2 * time.Hour),
		To:     time.Now().Add(-time.Hour),
	})
	if err != nil {
		t.Fail()
		return
	}
	fmt.Println(o)
}
