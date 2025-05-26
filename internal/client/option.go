package client

import (
	"github.com/sg3t41/go-coincheck/internal/client/http"
	"github.com/sg3t41/go-coincheck/internal/client/websocket"
	"github.com/sg3t41/go-coincheck/internal/e"
)

type Option func(*client) error

// --- 単体テスト用に差し替え可能にする ---
var httpNewClient = http.NewClient
var websocketNewClient = websocket.NewClient

func WithREST(key, secret, baseURL string) Option {
	return func(c *client) error {
		httpClient, err := httpNewClient(
			http.WithCredentials(key, secret),
			http.WithBaseURL(baseURL),
		)
		if err != nil {
			return e.WithPrefixError(err)
		}
		c.HTTPClient = httpClient
		return nil
	}
}

func WithWebSocket(baseURL string) Option {
	return func(c *client) error {
		websocketClient, err := websocketNewClient(
			websocket.WithBaseURL(baseURL),
		)
		if err != nil {
			return e.WithPrefixError(err)
		}
		c.WebSocketClient = websocketClient
		return nil
	}
}
