package client

import (
	"github.com/sg3t41/go-coincheck/internal/client/http"
	"github.com/sg3t41/go-coincheck/internal/client/websocket"
	"github.com/sg3t41/go-coincheck/internal/e"
)

type client struct {
	http.HTTPClient
	websocket.WebSocketClient
}

type Client interface {
	http.HTTPClient
	websocket.WebSocketClient
}

func New(key, secret string) (Client, error) {
	httpClient, err := http.NewHTTPClient(key, secret)
	if err != nil {
		return nil, e.WithPrefixError(err)
	}

	websocketClient, err := websocket.NewWebSocketClient()
	if err != nil {
		return nil, e.WithPrefixError(err)
	}

	return &client{
		httpClient,
		websocketClient,
	}, nil
}
