package mbc

// OrderType represents the order side
type OrderType int

var (
	Buy  OrderType = 1
	Sell OrderType = 2
)

// OrderStatus represents the order status
type OrderStatus int

var (
	Pending   OrderStatus = 1
	Open      OrderStatus = 2
	Cancelled OrderStatus = 3
	Filled    OrderStatus = 4
)

// StatusCode represents the response status
type StatusCode int

const (
	ErrorCode   StatusCode = 201
	SuccessCode StatusCode = 100
)
