package mbc_test

import (
	"context"
	"fmt"
	"github.com/lucaskatayama/mbc/v4"
	"testing"
	"time"
)

func TestWebsocket_Simple(t *testing.T) {
	c, err := mbc.NewPublicOnlyClient(mbc.WithWebsocket())
	if err != nil {
		t.Fail()
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	c.Websocket.OnError(func(err error) {
		fmt.Println("ERROR")
		fmt.Println(err)
	})

	c.Websocket.Connect(context.Background())

	handler := func(name string) mbc.WebSocketHandler {
		return func(msg mbc.WebSocketMessage) {
			fmt.Printf("[%s] %v\n", name, string(msg.Data))
		}
	}
	c.Websocket.SubscribeOrderbook("BTC-BRl", mbc.Orderbook10, handler("ordebook"))
	c.Websocket.SubscribeTicker("BTC-BRl", handler("ticker"))
	<-ctx.Done()
	c.Websocket.Close()
}
