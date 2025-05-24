package client

import (
	"github.com/sg3t41/go-coincheck/external/e"
	"github.com/sg3t41/go-coincheck/internal/client/http"
	"github.com/sg3t41/go-coincheck/internal/client/websocket"
)

type client struct {
	http.HTTPClient
	websocket.WebSocketClient
}

type Client interface {
	http.HTTPClient
	websocket.WebSocketClient
}

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

func New(opts ...Option) (Client, error) {
	c := &client{}
	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}
	return c, nil
}
