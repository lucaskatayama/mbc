
package mbc

import (
	"context"

)

// Creates a client for public and private requests
func ExampleNew() {
	client := New("<ID>", "<SECRET>")
	client.Trades(context.Background(), "btc", "brl")
}
