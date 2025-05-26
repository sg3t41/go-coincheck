package client

import (
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

func New(opts ...Option) (Client, error) {
	c := &client{}
	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}
	return c, nil
}
