package mbc

const endpoint = "https://www.mercadobitcoin.net"

// Client represents a MercadoBitcoin API client
type Client struct {
	id     string
	secret string
}

// ClientOpt represents a client option
type ClientOpt func(client *Client)

// WithIdSecret adds id/secret auth option to client
func WithIdSecret(id, secret string) ClientOpt {
	return func(client *Client) {
		client.id = id
		client.secret = secret
	}
}

// New creates a client for an user id/secret
func New(opts ...ClientOpt) *Client {
	c := &Client{}

	for _, opt := range opts {
		opt(c)
	}

	return c
}
