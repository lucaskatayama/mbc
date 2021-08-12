package mbc

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// OrderbookOption represents orderbook request option
type OrderbookOption func(values url.Values)

// WithLimit adds limit to orderbook request
func WithLimit(limit int) OrderbookOption {
	return func(values url.Values) {
		values.Add("limit", fmt.Sprintf("%d", limit))
	}
}

// TradesOption represents a trades request option
type TradesOption func(values url.Values)

// FromTid filters trades starting from given tid
func FromTid(tid int64) TradesOption {
	return func(values url.Values) {
		values.Add("tid", fmt.Sprintf("%d", tid))
	}
}

// ToID adds ID upper limit
func ToID(id string) ListOrdersOption {
	return func(values url.Values) {
		values.Add("to_id", id)
	}
}

// FromID adds ID lower limit
func FromID(id string) ListOrdersOption {
	return func(values url.Values) {
		values.Add("from_id", id)
	}
}

// WithOrderTypes adds order type filter
func WithOrderTypes(t OrderType) ListOrdersOption {
	return func(values url.Values) {
		values.Add("order_type", fmt.Sprintf("%d", t))
	}
}

// WithStatuses adds status filter
func WithStatuses(statuses ...OrderStatus) ListOrdersOption {
	return func(values url.Values) {
		b, _ := json.Marshal(statuses)
		values.Add("status_list", string(b))
	}
}

// HasFills filter if order has fills
func HasFills(has bool) ListOrdersOption {
	return func(values url.Values) {
		fill := "0"
		if has {
			fill = "1"
		}
		values.Add("has_fills", fill)
	}
}

// FromTimestamp adds lower limit for timestamp
func FromTimestamp(ts int64) ListOrdersOption {
	return func(values url.Values) {
		values.Add("from_timestamp", fmt.Sprintf("%d", ts))
	}
}

// ToTimestamp adds lower limit for timestamp
func ToTimestamp(ts int64) ListOrdersOption {
	return func(values url.Values) {
		values.Add("to_timestamp", fmt.Sprintf("%d", ts))
	}
}

// ListOrdersOption represents a request option
type ListOrdersOption func(values url.Values)
