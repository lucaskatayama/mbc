package mbc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"net/url"
	"sync"
)

// WebSocketHandler represents a websocket message handler (WebSocketMessage)
type WebSocketHandler func(msg WebSocketMessage)

type Subscription struct {
	Name  SubscriptionType `json:"name"`
	Id    InstrumentSymbol `json:"id"`
	Limit OrderbookSize    `json:"limit,omitempty"`
}

type Message struct {
	Type         string       `json:"type"`
	Subscription Subscription `json:"subscription"`
}

// OrderbookSize represetns an orderbook size limit on orderbook websocket subscription
type OrderbookSize int64

const (
	Orderbook10  OrderbookSize = 10
	Orderbook20  OrderbookSize = 20
	Orderbook50  OrderbookSize = 50
	Orderbook100 OrderbookSize = 100
	Orderbook200 OrderbookSize = 200
)

// WebSocketMessage represents a websocket message
// Data can be unmarshalled
type WebSocketMessage struct {
	Type  SubscriptionType `json:"type"`
	ID    InstrumentSymbol `json:"id"`
	Limit OrderbookSize    `json:"limit,omitempty"`
	Data  json.RawMessage  `json:"data"`
}

type subHandlerMap struct {
	*sync.RWMutex
	m map[string]WebSocketHandler
}

func (s *subHandlerMap) add(key string, h WebSocketHandler) {
	s.Lock()
	s.m[key] = h
	s.Unlock()
}

func (s *subHandlerMap) get(key string) (WebSocketHandler, bool) {
	s.RLock()
	h, ok := s.m[key]
	s.RUnlock()
	return h, ok
}

// WebSocketService contains Websocket operations
type WebSocketService struct {
	conn       *websocket.Conn
	baseURL    *url.URL
	subHandler subHandlerMap
	msgHandler func(msg []byte)
	errHandler func(err error)
}

func (s *WebSocketService) setBaseURL(base string) error {
	u, err := url.Parse(base)
	if err != nil {
		return err
	}
	s.baseURL = u
	return nil
}

// Connect connects to the websocket server and starts listening to messages
func (s *WebSocketService) Connect(ctx context.Context) error {
	conn, _, err := websocket.DefaultDialer.DialContext(ctx, s.baseURL.String(), http.Header{
		"Origin": []string{"github.com/lucaskatayama/mbc"},
	})
	if err != nil {
		return err
	}
	s.conn = conn

	go func() {
		for {
			_, message, err := s.conn.ReadMessage()
			if err != nil && s.errHandler != nil {
				s.errHandler(err)
				return
			}
			var m WebSocketMessage
			if err := json.Unmarshal(message, &m); err != nil {
				s.errHandler(err)
				continue
			}
			i := m.ID.normalize()
			m.ID = i
			if m.Data == nil {
				continue
			}
			key := fmt.Sprintf("%s.%s.%d", i, m.Type, m.Limit)
			if h, ok := s.subHandler.get(key); ok {
				h(m)
			}
		}
	}()

	return nil
}

// Close closes a websocket connection
func (s *WebSocketService) Close() error {
	if s.conn != nil {
		return s.conn.Close()
	}
	return nil
}

// OnError sets an error handler
func (s *WebSocketService) OnError(h func(err error)) {
	s.errHandler = h
}

// Send sends a generic message
func (s *WebSocketService) Send(msg interface{}) error {
	if s.conn == nil {
		return errors.New("not connected")
	}
	return s.conn.WriteJSON(msg)
}

// SubscriptionType represents a websocket subscription type
type SubscriptionType string

const (
	TickerType    SubscriptionType = "ticker"
	OrderbookType SubscriptionType = "orderbook"
	TradeType     SubscriptionType = "trade"
)

func (s *WebSocketService) subscribe(t SubscriptionType, instrument InstrumentSymbol, limit OrderbookSize, h func(msg WebSocketMessage)) error {
	if s.conn == nil {
		return errors.New("not connected")
	}
	key := fmt.Sprintf("%s.%s.%d", instrument, t, limit)
	s.subHandler.add(key, h)

	return s.conn.WriteJSON(Message{
		Type: "subscribe",
		Subscription: Subscription{
			Name:  t,
			Id:    instrument.toMB(),
			Limit: limit,
		},
	})
}

func (s *WebSocketService) SubscribeOrderbook(instrument InstrumentSymbol, size OrderbookSize, h WebSocketHandler) error {
	return s.subscribe(OrderbookType, instrument, size, h)
}

func (s *WebSocketService) SubscribeTrade(instrument InstrumentSymbol, h WebSocketHandler) error {
	return s.subscribe(TradeType, instrument, 0, h)
}

func (s *WebSocketService) SubscribeTicker(instrument InstrumentSymbol, h WebSocketHandler) error {
	return s.subscribe(TickerType, instrument, 0, h)
}
