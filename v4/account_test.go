package mbc_test

import (
	"context"
	"encoding/json"
	"github.com/lucaskatayama/mbc/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestAccountService_ListAccounts(t *testing.T) {
	assert := assert.New(t)

	mux, server, client := setup(t, "id", "secret")
	defer teardown(server)

	mux.HandleFunc("/api/v4/accounts",
		func(w http.ResponseWriter, r *http.Request) {
			testMethod(t, r, http.MethodGet)
			w.WriteHeader(http.StatusOK)
			data := []mbc.Account{
				{
					Currency:     "BRL",
					CurrencySign: "R$",
					Id:           "1",
					Name:         "Default Account",
					Type:         "live",
				},
			}
			resp, _ := json.Marshal(data)
			_, _ = w.Write(resp)
		},
	)

	got, resp, err := client.Account.ListAccounts(context.Background())
	if err != nil {
		t.Fatalf("Account.ListAccounts returned error: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("respnse is not OK/200: %d", resp.StatusCode)
	}
	assert.Len(got, 1, "should be 1")
}
