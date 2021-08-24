package mbc

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Creates a client for public and private requests
func ExampleNew() {
	client := New()
	client.Trades(context.Background(), "btc", "brl")
}

func TestNew(t *testing.T) {
	assert := assert.New(t)
	client := New()
	ticker, err := client.Ticker(context.Background(), "btc", "brl")
	if err != nil {
		t.Errorf("failed to fetch ticker: %+v", err)
		return
	}

	assert.NotEmpty(ticker.Buy, "buy should not be empty")
	assert.NotEmpty(ticker.Sell, "sell should not be empty")
}

func TestNew_WithIdSecretForPublicAPI(t *testing.T) {
	assert := assert.New(t)
	client := New(WithIdSecret("empty", "empty"))
	ticker, err := client.Ticker(context.Background(), "btc", "brl")
	if err != nil {
		t.Errorf("failed to fetch ticker: %+v", err)
		return
	}

	assert.NotEmpty(ticker.Buy, "buy should not be empty")
	assert.NotEmpty(ticker.Sell, "sell should not be empty")
}
