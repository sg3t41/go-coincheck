package client

import (
	"errors"
	"testing"

	"github.com/sg3t41/go-coincheck/internal/client/http"
	"github.com/sg3t41/go-coincheck/internal/client/websocket"
)

// モックのHTTPClient, WebSocketClient
type mockHTTP struct{ http.HTTPClient }
type mockWS struct{ websocket.WebSocketClient }

func withHTTP(h http.HTTPClient) Option {
	return func(c *client) error {
		c.HTTPClient = h
		return nil
	}
}
func withWS(w websocket.WebSocketClient) Option {
	return func(c *client) error {
		c.WebSocketClient = w
		return nil
	}
}
func withError() Option {
	return func(c *client) error {
		return errors.New("option error")
	}
}

func TestNewClient(t *testing.T) {
	tests := map[string]struct {
		opts     []Option
		wantErr  bool
		wantHTTP bool
		wantWS   bool
	}{
		"オプション無し": {
			opts:     nil,
			wantErr:  false,
			wantHTTP: false,
			wantWS:   false,
		},
		"HTTP だけセット": {
			opts:     []Option{withHTTP(&mockHTTP{})},
			wantErr:  false,
			wantHTTP: true,
			wantWS:   false,
		},
		"WS だけセット": {
			opts:     []Option{withWS(&mockWS{})},
			wantErr:  false,
			wantHTTP: false,
			wantWS:   true,
		},
		"両方セット": {
			opts:     []Option{withHTTP(&mockHTTP{}), withWS(&mockWS{})},
			wantErr:  false,
			wantHTTP: true,
			wantWS:   true,
		},
		"エラーオプション": {
			opts:     []Option{withError()},
			wantErr:  true,
			wantHTTP: false,
			wantWS:   false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			c, err := New(tt.opts...)
			if (err != nil) != tt.wantErr {
				t.Fatalf("error expectation: wantErr=%v, got=%v (err=%v)", tt.wantErr, err != nil, err)
			}
			if tt.wantErr {
				if c != nil {
					t.Errorf("got client != nil, want nil")
				}
				return
			}
			if c == nil {
				t.Fatalf("got client == nil, want not nil")
			}

			// 型アサーションして中身をチェック
			cli := c.(*client)
			if tt.wantHTTP && cli.HTTPClient == nil {
				t.Errorf("HTTPClient is nil, want not nil")
			}
			if !tt.wantHTTP && cli.HTTPClient != nil {
				t.Errorf("HTTPClient is not nil, want nil")
			}
			if tt.wantWS && cli.WebSocketClient == nil {
				t.Errorf("WebSocketClient is nil, want not nil")
			}
			if !tt.wantWS && cli.WebSocketClient != nil {
				t.Errorf("WebSocketClient is not nil, want nil")
			}
		})
	}
}
