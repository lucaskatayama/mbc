package mbc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/go-querystring/query"
	"github.com/hashicorp/go-cleanhttp"
	"github.com/hashicorp/go-retryablehttp"
	"net/http"
	"net/url"
	"time"
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
)

type Client struct {
	// HTTP client used to communicate with the API.
	client *retryablehttp.Client

	baseURL *url.URL

	// disableRetries is used to disable the default retry logic.
	disableRetries bool

	// username and password used for basic authentication.
	username, password string

	UserAgent string

	PublicData *PublicDataService
	Websocket  *WebsocketService
}

func newClient(opts ...ClientOpt) (*Client, error) {
	c := &Client{UserAgent: userAgent}
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

// NewClient returns a client with username/password authentication
func NewClient(username, password string, opts ...ClientOpt) (*Client, error) {
	client, err := newClient(opts...)
	if err != nil {
		return nil, err
	}
	client.username = username
	client.password = password
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

type RequestOpt func(r *retryablehttp.Request) error

func (c *Client) NewRequest(ctx context.Context, method string, path string, body interface{}, options []RequestOpt) (*retryablehttp.Request, error) {
	u := *c.baseURL
	p, err := url.PathUnescape(path)
	if err != nil {
		return nil, err
	}
	u.RawPath = c.baseURL.Path + path
	u.Path = c.baseURL.Path + p

	// Create a request specific headers map.
	reqHeaders := make(http.Header)
	reqHeaders.Set("Accept", "application/json")

	if c.UserAgent != "" {
		reqHeaders.Set("User-Agent", c.UserAgent)
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
	fmt.Println(req.URL.String())

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

type Response struct {
	*http.Response
}

func (c *Client) Do(req *retryablehttp.Request, ptr interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

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
