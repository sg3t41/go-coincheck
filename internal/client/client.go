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

func New(key, secret string) (Client, error) {
	httpClient, err := http.NewClient(key, secret)
	if err != nil {
		return nil, e.WithPrefixError(err)
	}

	websocketClient, err := websocket.NewClient()
	if err != nil {
		return nil, e.WithPrefixError(err)
	}

	return &client{
		httpClient,
		websocketClient,
	}, nil
}
