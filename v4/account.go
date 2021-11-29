package mbc

import (
	"context"
	"fmt"
	"net/http"
)

// AccountService handles account operations
type AccountService struct {
	client *Client
}

type AccountID string

// Account represents an account
type Account struct {
	Currency     string    `json:"currency"`
	CurrencySign string    `json:"currencySign"`
	Id           AccountID `json:"id"`
	Name         string    `json:"name"`
	Type         string    `json:"type"`
}

// ListAccounts returns a list of available accounts
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

// ListBalances fecthes balances for each asset for an AccountID
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
