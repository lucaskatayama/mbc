package mbc

import (
	"context"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var (
	privateEndpoint = fmt.Sprintf("%s/tapi/v3/", endpoint)
)

func mac(secret string, body string) string {
	h := hmac.New(sha512.New, []byte(secret))
	data := fmt.Sprintf("/tapi/v3/?%s", body)
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func (c *Client) formFor(base, quote, method string) url.Values {
	form := url.Values{}
	form.Add("tapi_method", method)
	form.Add("coin_pair", fmt.Sprintf("%s%s", strings.ToUpper(quote), strings.ToUpper(base)))

	form.Add("tapi_nonce", fmt.Sprintf("%d", time.Now().UnixNano()))

	return form
}

// ListOrders lists user orders
func (c *Client) ListOrders(ctx context.Context, base string, quote string, opts ...ListOrdersOption) ([]Order, error) {
	form := c.formFor(base, quote, "list_orders")
	for _, opt := range opts {
		opt(form)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, privateEndpoint, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("TAPI-MAC", mac(c.secret, form.Encode()))
	req.Header.Add("TAPI-ID", c.id)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var orderD orderD
	if err := json.NewDecoder(res.Body).Decode(&orderD); err != nil {
		return nil, err
	}

	if orderD.StatusCode != SuccessCode {
		return nil, errors.New(orderD.ErrorMessage)
	}

	return orderD.ResponseData.Orders, nil
}
