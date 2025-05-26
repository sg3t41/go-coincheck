package websocket

import (
	"errors"
	"net/url"

	"github.com/sg3t41/go-coincheck/internal/e"
)

type Option func(*webSocketClient) error

func WithBaseURL(strURL string) Option {
	return func(wsc *webSocketClient) error {
		if strURL == "" {
			return e.WithPrefixError(errors.New("WebSocketのベースURLが空です"))
		}
		url, err := url.Parse(strURL)
		if err != nil {
			return e.WithPrefixError(err)
		}

		wsc.url = url
		return nil
	}
}
