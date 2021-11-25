package mbc

import (
	"context"
	"fmt"
	"net/http"
)

type AccountService struct {
	client *Client
}

type AccountID string

type Account struct {
	Currency     string    `json:"currency"`
	CurrencySign string    `json:"currencySign"`
	Id           AccountID `json:"id"`
	Name         string    `json:"name"`
	Type         string    `json:"type"`
}

func (s *AccountService) ListAccounts(ctx context.Context, opts ...RequestOpt) ([]Account, *http.Response, error) {
	p := "/accounts"
	req, err := s.client.newRequest(ctx, http.MethodGet, p, nil, opts)
	if err != nil {
		return nil, nil, err
	}
	var accounts []Account
	resp, err := s.client.do(req, &accounts)

	if err != nil {
		return nil, nil, err
	}

	return accounts, resp, nil
}

// AccountBalance represents an asset balance for account
type AccountBalance struct {
	Available float64     `json:"available"`
	Symbol    AssetSymbol `json:"symbol"`
	Total     float64     `json:"total"`
}

func (s *AccountService) ListBalances(ctx context.Context, accountID AccountID, opts ...RequestOpt) ([]AccountBalance, *http.Response, error) {
	p := fmt.Sprintf("/accounts/%s/balances", accountID)
	req, err := s.client.newRequest(ctx, http.MethodGet, p, nil, opts)
	if err != nil {
		return nil, nil, err
	}
	var balances []AccountBalance
	resp, err := s.client.do(req, &balances)

	if err != nil {
		return nil, nil, err
	}

	return balances, resp, nil
}
