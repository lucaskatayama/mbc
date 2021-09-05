# MercadoBitcoin API v3

## Installation

```
$ go get -u github.com/lucaskatayama/mbc/v3
```

## Usage

### Initialize client

For public data:

```go
package main 

import "github.com/lucaskatayama/mbc/v3"

func main() {
    client := mbc.New()
}
```

For private data, use your api key and secret

```go
package main 

import "github.com/lucaskatayama/mbc/v3"

func main() {
    client := mbc.New(mbc.WithIdSecret("<id>", "<secret>"))
}
```

### Public data

```go
package main

import (
  "context"
  "github.com/lucaskatayama/mbc/v3"
)

func main() {
  client := mbc.New()
  // ticker
  ticker, err := client.Ticker(context.Background(), "BTC", "BRL")
  if err != nil {
	  log.Panic(err)
  }

  // orderbook
  ordebook, err := client.Orderbook(context.Background(), "BTC", "BRL")
  if err != nil {
    log.Panic(err)
  }

  // trades
  trades, err := client.Trades(context.Background(), "BTC", "BRL")
  if err != nil {
    log.Panic(err)
  }
}
```

## API

- Public API
    - [x] ~~orderbook~~ Orderbook
    - [x] ~~ticker~~ Ticker
    - [x] ~~trades~~ Trades
- Trade API
    - [ ] list_system_messages
    - [x] ~~get_account_info~~ GetBalances/GetWithdrawalLimits
    - [ ] get_order
    - [x] ~~list_orders~~ ListOrders
    - [ ] list_orderbook
    - [ ] place_buy_order
    - [ ] place_sell_order
    - [ ] place_postonly_buy_order
    - [ ] place_postonly_sell_order
    - [ ] place_market_buy_order
    - [ ] place_market_sell_order
    - [ ] cancel_order
    - [ ] get_withdrawal
    - [ ] withdraw_coin

## CLI

[Under Construction]


## References

- [Public API](https://www.mercadobitcoin.com.br/api-doc/)
- [Trading API](https://www.mercadobitcoin.com.br/trade-api/)
