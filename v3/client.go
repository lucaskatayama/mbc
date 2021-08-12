package mbc

const endpoint = "https://www.mercadobitcoin.net"

// Client represents a MercadoBitcoin API client
type Client struct {
	id     string
	secret string
}

// New creates a client for an user id/secret
func New(id string, secret string) Client {
	return Client{
		id,
		secret,
	}
}
