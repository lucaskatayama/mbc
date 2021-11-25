package mbc

import (
	"context"
	"fmt"
	"net/http"
)

type TradingService struct {
	client *Client
}

type OrderType string

const (
	Market   OrderType = "market"
	Limit    OrderType = "limit"
	PostOnly OrderType = "post-only"
)

type OrderSide string

const (
	Buy  OrderSide = "buy"
	Sell OrderSide = "sell"
)

type OrderRequest struct {
	Type       OrderType `json:"type"`
	Side       OrderSide `json:"side"`
	Async      bool      `json:"async"`
	Cost       int       `json:"cost,omitempty"`
	LimitPrice int       `json:"limitPrice,omitempty"`
	Qty        float64   `json:"qty,omitempty"`
}

type Order struct {
	ID         OrderID          `json:"id"`
	Instrument InstrumentSymbol `json:"instrument"`
	Side       OrderSide        `json:"side"`
	Type       OrderType        `json:"type"`
	Status     string           `json:"status"`
	CreatedAt  UnixTime         `json:"created_at"`
	UpdatedAt  UnixTime         `json:"updated_at"`
	Qty        float64          `json:"qty"`
	LimitPrice int              `json:"limitPrice"`
	AvgPrice   UnixTime         `json:"avgPrice"`
	FilledQty  float64          `json:"filledQty"`
	Operations []Operations     `json:"executions"`
}

type OrderID string

type placeOrderResponse struct {
	OrderID OrderID `json:"orderId"`
}

type Operations struct{}

// ListOrders lists orders for AccountID and InstrumentSymbol
func (s *TradingService) ListOrders(ctx context.Context, accountID AccountID, instrument InstrumentSymbol, opts ...RequestOpt) ([]Order, *http.Response, error) {
	panic("not implemented")
}

// PlaceOrder places an order for AccountID with InstrumentSymbol
func (s *TradingService) PlaceOrder(ctx context.Context, accountID AccountID, instrument InstrumentSymbol, o OrderRequest, opts ...RequestOpt) (OrderID, *http.Response, error) {
	p := fmt.Sprintf("/accounts/%s/%s/orders", accountID, instrument)
	req, err := s.client.newRequest(ctx, http.MethodPost, p, o, opts)
	if err != nil {
		return "", nil, err
	}
	var order placeOrderResponse
	resp, err := s.client.do(req, &order)
	if err != nil {
		return "", nil, err
	}
	return order.OrderID, resp, nil
}

// CancelOrder cancels an order by OrderID for AccountID and InstrumentSymbol
func (s *TradingService) CancelOrder(ctx context.Context, id AccountID, symbol InstrumentSymbol, orderID OrderID, opts ...RequestOpt) (bool, *http.Response, error) {
	panic("not implemented")
}

// GetOrder fetches an order by OrderID for AccountID and InstrumentSymbol
func (s *TradingService) GetOrder(ctx context.Context, id AccountID, symbol InstrumentSymbol, orderID OrderID, opts ...RequestOpt) (Order, *http.Response, error) {
	panic("not implemented")
}
