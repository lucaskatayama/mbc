package mbc

import (
	"errors"
	"github.com/hashicorp/go-retryablehttp"
	"sync"
)

type ClientOpt func(c *Client) error

func WithMaxRetries(n int) ClientOpt {
	return func(c *Client) error {
		c.client.RetryMax = n
		return nil
	}
}

func WithErrorHandler(h retryablehttp.ErrorHandler) ClientOpt {
	return func(c *Client) error {
		if h == nil {
			return errors.New("error handler is nil")
		}
		c.client.ErrorHandler = h
		return nil
	}
}

func WithWebsocket() ClientOpt {
	return func(c *Client) error {
		c.Websocket = &WebsocketService{
			subHandler: subHandlerMap{
				RWMutex: &sync.RWMutex{},
				m:       map[string]func(msg []byte){},
			},
		}
		if err := c.Websocket.setBaseURL(defaultWSBaseURL); err != nil {
			return err
		}
		return nil
	}
}
