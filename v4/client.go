package mbc

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/hashicorp/go-cleanhttp"
	"github.com/hashicorp/go-retryablehttp"

	"github.com/lucaskatayama/mbc/v4/utils"
)

const (
	defaultBaseURL    = "https://api.mercadobitcoin.net"
	apiVersionPath    = "/api/v4"
	userAgent         = "github.com/lucaskatayama/mbc"
	defaultMaxRetries = 5
	defaultWSBaseURL  = "wss://ws.mercadobitcoin.net/ws"
)

type AuthType int

const (
	BasicAuth AuthType = iota
	OAuth
)

type Client struct {
	// HTTP client used to communicate with the API.
	client *retryablehttp.Client

	baseURL *url.URL

	// disableRetries is used to disable the default retry logic.
	disableRetries bool

	// id and secret used for basic authentication.
	id, secret string

	userAgent string

	PublicData *PublicDataService
	Account    *AccountService
	Trading    *TradingService
	Websocket  *WebSocketService
	log        utils.Log
}

// RequestOpt changes the retryablehttp.Request
type RequestOpt func(r *retryablehttp.Request) error

// NewClient returns a client with id/secret authentication
func NewClient(id, secret string, opts ...ClientOpt) (*Client, error) {
	client, err := newClient(opts...)
	if err != nil {
		return nil, err
	}
	client.id = id
	client.secret = secret
	return client, nil
}

// NewPublicOnlyClient returns a client for PublicData access only
func NewPublicOnlyClient(opts ...ClientOpt) (*Client, error) {
	client, err := newClient(opts...)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func newClient(opts ...ClientOpt) (*Client, error) {
	c := &Client{
		userAgent: userAgent,
		log:       utils.DummyLog{},
	}
	// Configure the HTTP client.
	c.client = &retryablehttp.Client{
		CheckRetry:   retryablehttp.DefaultRetryPolicy,
		Backoff:      retryablehttp.DefaultBackoff,
		ErrorHandler: retryablehttp.PassthroughErrorHandler,
		HTTPClient:   cleanhttp.DefaultPooledClient(),
		RetryWaitMin: 100 * time.Millisecond,
		RetryWaitMax: 400 * time.Millisecond,
		RetryMax:     defaultMaxRetries,
	}

	_ = c.setBaseURL(defaultBaseURL)

	c.PublicData = &PublicDataService{c}
	c.Account = &AccountService{c}
	c.Trading = &TradingService{c}

	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}
	return c, nil
}

func (c *Client) setBaseURL(baseURL string) error {
	u, err := url.Parse(baseURL)
	if err != nil {
		return err
	}

	u.Path += apiVersionPath

	c.baseURL = u
	return nil
}

// newRequest creates a request
func (c *Client) newRequest(ctx context.Context, method string, path string, body interface{}, options []RequestOpt) (*retryablehttp.Request, error) {
	u := *c.baseURL
	p, err := url.PathUnescape(path)
	if err != nil {
		return nil, err
	}
	u.RawPath = c.baseURL.Path + path
	u.Path = c.baseURL.Path + p

	// Create a request specific headers map.
	reqHeaders := make(http.Header)

	if c.userAgent != "" {
		reqHeaders.Set("User-Agent", c.userAgent)
	}

	var rBody interface{}
	switch {
	case method == http.MethodPost || method == http.MethodPut:
		reqHeaders.Set("Content-Type", "application/json")

		if body != nil {
			rBody, err = json.Marshal(body)
			if err != nil {
				return nil, err
			}
		}
	case body != nil:
		q, err := query.Values(body)
		if err != nil {
			return nil, err
		}
		u.RawQuery = q.Encode()
	}

	req, err := retryablehttp.NewRequest(method, u.String(), rBody)

	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

	for _, fn := range options {
		if fn == nil {
			continue
		}
		if err := fn(req); err != nil {
			return nil, err
		}
	}

	// Set the request specific headers.
	for k, v := range reqHeaders {
		req.Header[k] = v
	}

	return req, nil

}

// do executes a request
func (c *Client) do(req *retryablehttp.Request, ptr interface{}) (*http.Response, error) {
	c.log.Debugf("%s %s", req.Method, req.URL.String())
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	c.log.Debugf("%d", resp.StatusCode)
	if resp.StatusCode > 400 {
		return nil, errors.New("request error")
	}

	if resp.StatusCode == http.StatusOK && ptr != nil {
		if err := json.NewDecoder(resp.Body).Decode(ptr); err != nil {
			return nil, err
		}
	}
	return resp, nil
}
