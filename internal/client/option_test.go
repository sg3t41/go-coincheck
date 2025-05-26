package client

import (
	"errors"
	"testing"

	"github.com/sg3t41/go-coincheck/internal/client/http"
	"github.com/sg3t41/go-coincheck/internal/client/websocket"
)

type mockHTTP2 struct{ http.HTTPClient }
type mockWS2 struct{ websocket.WebSocketClient }

func restore() {
	httpNewClient = http.NewClient
	websocketNewClient = websocket.NewClient
}

func TestWithREST_And_WithWebSocket(t *testing.T) {
	tests := map[string]struct {
		setupHTTPNewClient      func()
		setupWebSocketNewClient func()
		option                  Option
		wantHTTP                bool
		wantWS                  bool
		wantErr                 bool
	}{
		"WithREST: 正常系": {
			setupHTTPNewClient: func() {
				httpNewClient = func(...http.Option) (http.HTTPClient, error) {
					return &mockHTTP2{}, nil
				}
			},
			option:   WithREST("key", "secret", "url"),
			wantHTTP: true,
			wantWS:   false,
			wantErr:  false,
		},
		"WithREST: 異常系": {
			setupHTTPNewClient: func() {
				httpNewClient = func(...http.Option) (http.HTTPClient, error) {
					return nil, errors.New("fail http")
				}
			},
			option:   WithREST("key", "secret", "url"),
			wantHTTP: false,
			wantWS:   false,
			wantErr:  true,
		},
		"WithWebSocket: 正常系": {
			setupWebSocketNewClient: func() {
				websocketNewClient = func(...websocket.Option) (websocket.WebSocketClient, error) {
					return &mockWS2{}, nil
				}
			},
			option:   WithWebSocket("url"),
			wantHTTP: false,
			wantWS:   true,
			wantErr:  false,
		},
		"WithWebSocket: 異常系": {
			setupWebSocketNewClient: func() {
				websocketNewClient = func(...websocket.Option) (websocket.WebSocketClient, error) {
					return nil, errors.New("fail ws")
				}
			},
			option:   WithWebSocket("url"),
			wantHTTP: false,
			wantWS:   false,
			wantErr:  true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			defer restore()
			if tt.setupHTTPNewClient != nil {
				tt.setupHTTPNewClient()
			}
			if tt.setupWebSocketNewClient != nil {
				tt.setupWebSocketNewClient()
			}
			c := &client{}
			err := tt.option(c)
			if (err != nil) != tt.wantErr {
				t.Fatalf("error expectation: wantErr=%v, got=%v (err=%v)", tt.wantErr, err != nil, err)
			}
			if tt.wantErr {
				return
			}
			if tt.wantHTTP && c.HTTPClient == nil {
				t.Errorf("HTTPClient is nil, want not nil")
			}
			if !tt.wantHTTP && c.HTTPClient != nil {
				t.Errorf("HTTPClient is not nil, want nil")
			}
			if tt.wantWS && c.WebSocketClient == nil {
				t.Errorf("WebSocketClient is nil, want not nil")
			}
			if !tt.wantWS && c.WebSocketClient != nil {
				t.Errorf("WebSocketClient is not nil, want nil")
			}
		})
	}
}
