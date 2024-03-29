package mbc

import (
	"errors"
	"sync"

	"github.com/hashicorp/go-retryablehttp"

	"github.com/lucaskatayama/mbc/v4/utils"
)

type ClientOpt func(c *Client) error

// WithMaxRetries set a maximun retry limit
func WithMaxRetries(n int) ClientOpt {
	return func(c *Client) error {
		c.client.RetryMax = n
		return nil
	}
}

// WithErrorHandler sets a global error handler
func WithErrorHandler(h retryablehttp.ErrorHandler) ClientOpt {
	return func(c *Client) error {
		if h == nil {
			return errors.New("error handler is nil")
		}
		c.client.ErrorHandler = h
		return nil
	}
}

// WithBaseURL changes the base URL
func WithBaseURL(u string) ClientOpt {
	return func(c *Client) error {
		return c.setBaseURL(u)
	}
}

// WithWebsocket enables websocket on client
func WithWebsocket() ClientOpt {
	return func(c *Client) error {
		c.Websocket = &WebSocketService{
			subHandler: subHandlerMap{
				RWMutex: &sync.RWMutex{},
				m:       map[string]WebSocketHandler{},
			},
		}
		if err := c.Websocket.setBaseURL(defaultWSBaseURL); err != nil {
			return err
		}
		return nil
	}
}

// WithLogger sets a logger to client
func WithLogger(log utils.Log) ClientOpt {
	return func(c *Client) error {
		c.log = log
		if log == nil {
			return errors.New("log should not be nil")
		}
		return nil
	}
}
