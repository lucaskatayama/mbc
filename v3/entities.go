package mbc

// =======================   PUBLIC   ============================

// Ticker represents the ticker
type Ticker struct {
	High string `json:"high"`
	Low  string `json:"low"`
	Vol  string `json:"vol"`
	Last string `json:"last"`
	Buy  string `json:"buy"`
	Sell string `json:"sell"`
	Open string `json:"open"`
	Date int64  `json:"date"`
}

type tickerD struct {
	Ticker Ticker `json:"ticker"`
}

// Orderbook represents the orderbook
type Orderbook struct {
	Asks      [][]float64 `json:"asks"`
	Bids      [][]float64 `json:"bids"`
	Timestamp int64       `json:"timestamp"`
}

// Trade represents a trade
type Trade struct {
	Tid    int64   `json:"tid"`
	Date   int64   `json:"date"`
	Type   string  `json:"type"`
	Price  float64 `json:"price"`
	Amount float64 `json:"amount"`
}

// =========================   PRIVATE  =================================

// Operation represents an operation
type Operation struct {
	ID                int    `json:"operation_id"`
	Qty               string `json:"quantity"`
	Price             string `json:"price"`
	FeeRate           string `json:"fee_rate"`
	ExecutedTimestamp string `json:"executed_timestamp"`
}

// Order represents an order
type Order struct {
	ID               int         `json:"order_id"`
	Pair       string      `json:"coin_pair"`
	OrderType  OrderType   `json:"order_type"`
	Status     int         `json:"status"`
	HasFills         bool        `json:"has_fills"`
	Qty              string      `json:"quantity"`
	LimitPrice       string      `json:"limit_price"`
	ExecutedQuantity string      `json:"executed_quantity"`
	ExecutedPriceAvg string      `json:"executed_price_avg"`
	Fee              string      `json:"fee"`
	CreatedAt        string      `json:"created_timestamp"`
	UpdatedAt  string      `json:"updated_timestamp"`
	Operations []Operation `json:"operations"`
}

type orderD struct {
	ResponseData struct {
		Orders []Order `json:"orders"`
	} `json:"response_data"`
	StatusCode          StatusCode `json:"status_code"`
	ServerUnixTimestamp string     `json:"server_unix_timestamp"`
	ErrorMessage        string     `json:"error_message"`
}
