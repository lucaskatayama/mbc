# MercadoBitcoin API v4

## Installation

```
$ go get -u github.com/lucaskatayama/mbc/v4
```

## Usage

### Initialize client

For public data only:

```go
package main 

import "github.com/lucaskatayama/mbc/v4"

func main() {
	client, err := mbc.NewPublicOnlyClient()
	if err != nil {
		panic(err)
    }
}
```


### Public data

```go
package main

import (
  "context"
  "github.com/lucaskatayama/mbc/v4"
)

func main() {
	client, err := mbc.NewPublicOnlyClient()
  // ticker
  tickers, err := client.ListTickers(context.Background(), []string{"BTC-BRL", "ltc-brl"})
  if err != nil {
	  log.Panic(err)
  }

  // orderbook
  ordebook, err := client.GetOrderbook(context.Background(), "BTC-BRL")
  if err != nil {
    log.Panic(err)
  }

  // trades
  trades, err := client.ListTrades(context.Background(), "BTC-BRL")
  if err != nil {
    log.Panic(err)
  }
}
```


## References

- [API v4](https://api.mercadobitcoin.net)
