package mbc

import (
	"context"
	"fmt"
	"os"
	"testing"
)

func TestClient_ListOrders(t *testing.T) {
	client := New(os.Getenv("MB_ID"), os.Getenv("MB_SECRET"))
	orders, err := client.ListOrders(context.Background(), "btc", "brl", WithStatuses(Cancelled, Open))
	if err != nil {
		panic(err)
	}
	fmt.Println(orders)
}
