package mbc

const endpoint = "https://www.mercadobitcoin.net"

// Client represents a MercadoBitcoin API client
type Client struct {
	id     string
	secret string
}

type ClientOpt func(client *Client)

func WithIdSecret(id, secret string) ClientOpt{
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
