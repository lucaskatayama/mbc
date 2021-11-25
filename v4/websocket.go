package mbc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

// WebsocketHandler represents a websocket message handler (WebsocketMessage)
type WebsocketHandler func(msg WebsocketMessage)

type subHandlerMap struct {
	*sync.RWMutex
	m map[string]WebsocketHandler
}

func (s *subHandlerMap) add(key string, h WebsocketHandler) {
	s.Lock()
	s.m[key] = h
	s.Unlock()
}

func (s *subHandlerMap) get(key string) (WebsocketHandler, bool) {
	s.RLock()
	h, ok := s.m[key]
	s.RUnlock()
	return h, ok
}

type WebsocketService struct {
	conn       *websocket.Conn
	baseURL    *url.URL
	subHandler subHandlerMap
	msgHandler func(msg []byte)
	errHandler func(err error)
}

func (s *WebsocketService) setBaseURL(base string) error {
	u, err := url.Parse(base)
	if err != nil {
		return err
	}
	s.baseURL = u
	return nil
}

// WebsocketMessage represents a websocket message
// Data can be unmarshalled
type WebsocketMessage struct {
	Type  string          `json:"type"`
	ID    string          `json:"id"`
	Limit int64           `json:"limit"`
	Data  json.RawMessage `json:"data"`
}

// Connect connects to the websocket server and starts listening to messages
func (s *WebsocketService) Connect(ctx context.Context) error {
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
			var m WebsocketMessage
			if err := json.Unmarshal(message, &m); err != nil {
				s.errHandler(err)
				continue
			}
			i := fmt.Sprintf("%s-%s", m.ID[3:], m.ID[:3])
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

func (s *WebsocketService) OnErr(h func(err error)) {
	s.errHandler = h
}

func (s *WebsocketService) Send(msg interface{}) error {
	if s.conn == nil {
		return errors.New("not connected")
	}
	return s.conn.WriteJSON(msg)
}

type SubscriptionType string

const (
	TickerType    SubscriptionType = "ticker"
	OrderbookType SubscriptionType = "orderbook"
	TradeType     SubscriptionType = "trade"
)

func (s *WebsocketService) normalizeInstrument(in string) string {
	if strings.Contains(in, "-") {
		parts := strings.Split(strings.ToUpper(in), "-")
		return fmt.Sprintf("%s%s", parts[1], parts[0])
	}
	return in
}

func (s *WebsocketService) subscribe(t SubscriptionType, instrument string, limit OrderbookSize, h func(msg WebsocketMessage)) error {
	instrument = strings.ToUpper(instrument)
	if s.conn == nil {
		return errors.New("not connected")
	}
	key := fmt.Sprintf("%s.%s.%d", instrument, t, limit)
	s.subHandler.add(key, h)

	return s.conn.WriteJSON(Message{
		Type: "subscribe",
		Subscription: Sub{
			Name:  string(t),
			Id:    s.normalizeInstrument(instrument),
			Limit: limit,
		},
	})
}

func (s *WebsocketService) Ticker(instrument string, h WebsocketHandler) error {
	return s.subscribe(TickerType, instrument, 0, h)
}

type OrderbookSize int64

const (
	Orderbook10  OrderbookSize = 10
	Orderbook20  OrderbookSize = 20
	Orderbook50  OrderbookSize = 50
	Orderbook100 OrderbookSize = 100
	Orderbook200 OrderbookSize = 200
)

func (s *WebsocketService) Orderbook(instrument string, size OrderbookSize, h WebsocketHandler) error {
	return s.subscribe(OrderbookType, instrument, size, h)
}

func (s *WebsocketService) Trade(instrument string, h WebsocketHandler) error {
	return s.subscribe(TradeType, instrument, 0, h)
}

func (s *WebsocketService) Close() error {
	if s.conn != nil {
		return s.conn.Close()
	}
	return nil
}

type Sub struct {
	Name  string        `json:"name"`
	Id    string        `json:"id"`
	Limit OrderbookSize `json:"limit,omitempty"`
}

type Message struct {
	Type         string `json:"type"`
	Subscription Sub    `json:"subscription"`
}
