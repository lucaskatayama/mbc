package mbc

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"net/url"
	"strings"
)

var (
	defaultOnErr Handler = func(msg []byte) {}
	defaultOnMsg Handler = func(msg []byte) {}
)

type WS struct {
	conn     *websocket.Conn
	endpoint string
	onMsg    Handler
	onErr    Handler
}

func NewWS() *WS {
	u := url.URL{Scheme: "wss", Host: prodWS, Path: "/ws"}
	return &WS{
		endpoint: u.String(),
		onMsg:    defaultOnMsg,
		onErr:    defaultOnErr,
	}
}

// Connect connects to the websocket server and starts listening to messages
func (ws *WS) Connect() error {
	conn, _, err := websocket.DefaultDialer.Dial(ws.endpoint, http.Header{
		"Origin": []string{"github.com/lucaskatayama/mbc"},
	})
	if err != nil {
		return err
	}
	ws.conn = conn

	go func() {
		defer conn.Close()
		for {
			var msg []byte
			_, msg, err := conn.ReadMessage()
			if err != nil {
				ws.onErr(msg)
				return
			}
			ws.onMsg(msg)

		}
	}()

	return nil
}

type Subscribe struct {
	Type         string       `json:"type"`
	Subscription Subscription `json:"subscription"`
}

type Subscription struct {
	Name  string `json:"name"`
	Id    string `json:"id"`
	Limit int64  `json:"limit"`
}

type Handler func(msg []byte)

func normalizeInstrument(i string) string {
	if strings.Contains(i, "-") {
		return i
	}
	return fmt.Sprintf("%s-%s", i[3:], i[:3])
}

// Subscribe creates a subscription
// Messages are handled by HandleMsg and HandleErr
func (ws *WS) Subscribe(ctx context.Context, t string, instrument string) error {
	if err := ws.conn.WriteJSON(Subscribe{
		Type: "subscribe",
		Subscription: Subscription{
			Name:  t,
			Id:    normalizeInstrument(instrument),
			Limit: 0,
		},
	}); err != nil {
		return err
	}
	return nil
}

func (ws *WS) HandleMsg(h Handler) {
	ws.onMsg = h
}

func (ws *WS) HandleErr(h Handler) {
	ws.onErr = h
}
