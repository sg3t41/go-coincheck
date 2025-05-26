package client

import (
	"github.com/sg3t41/go-coincheck/internal/client/http"
	"github.com/sg3t41/go-coincheck/internal/client/websocket"
	"github.com/sg3t41/go-coincheck/internal/e"
)

type Option func(*client) error

func WithREST(key, secret, baseURL string) Option {
	return func(c *client) error {
		httpClient, err := http.NewClient(
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
		websocketClient, err := websocket.NewClient(
			websocket.WithBaseURL(baseURL),
		)
		if err != nil {
			return e.WithPrefixError(err)
		}
		c.WebSocketClient = websocketClient
		return nil
	}
}
