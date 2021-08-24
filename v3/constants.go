package mbc

// OrderType represents the order side
type OrderType int

var (
	// Buy order type
	Buy  OrderType = 1
	// Sell order type
	Sell OrderType = 2
)

// OrderStatus represents the order status
type OrderStatus int

var (
	// Pending pending order status
	Pending   OrderStatus = 1
	// Open open order status
	Open      OrderStatus = 2
	// Cancelled cancelled order status
	Cancelled OrderStatus = 3
	// Filled filled order status
	Filled    OrderStatus = 4
)

// StatusCode represents the response status
type StatusCode int

const (
	// ErrorCode general error code
	ErrorCode   StatusCode = 201
	// SuccessCode success error code
	SuccessCode StatusCode = 100
)
