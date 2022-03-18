package mbc_test

import (
	"github.com/lucaskatayama/mbc/v4"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupPublicOnly(t *testing.T) (*http.ServeMux, *httptest.Server, *mbc.Client) {
	mux := http.NewServeMux()

	server := httptest.NewServer(mux)

	client, err := mbc.NewPublicOnlyClient(mbc.WithBaseURL(server.URL))
	if err != nil {
		server.Close()
		t.Fatalf("Failed to create client: %v", err)
	}

	return mux, server, client
}

func setup(t *testing.T, id, secret string) (*http.ServeMux, *httptest.Server, *mbc.Client) {
	mux := http.NewServeMux()

	server := httptest.NewServer(mux)

	client, err := mbc.NewClient(id, secret, mbc.WithBaseURL(server.URL))
	if err != nil {
		server.Close()
		t.Fatalf("Failed to create client: %v", err)
	}

	return mux, server, client
}

func teardown(server *httptest.Server) {
	server.Close()
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %s, want %s", got, want)
	}
}
