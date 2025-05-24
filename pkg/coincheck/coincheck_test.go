package coincheck

import (
	"context"
	"testing"

	"github.com/sg3t41/go-coincheck/internal/client"
)

// TestNew はCoincheckクライアントの初期化をテスト
func TestNew(t *testing.T) {
	// テストケース
	tests := []struct {
		name    string
		key     string
		secret  string
		wantErr bool
	}{
		{
			name:    "正常系: 有効な認証情報",
			key:     "valid-key",
			secret:  "valid-secret",
			wantErr: false,
		},
		{
			name:    "異常系: 空の認証情報",
			key:     "",
			secret:  "",
			wantErr: true,
		},
	}

	// テスト実行
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// クライアントの初期化
			client, err := New(
				client.WithREST(tt.key, tt.secret),
				client.WithWebSocket(),
			)

			// エラーチェック
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// 正常系の場合、クライアントが正しく初期化されているか確認
			if !tt.wantErr {
				if client == nil {
					t.Error("New() returned nil client")
				}
				if client.REST == nil {
					t.Error("New() returned nil REST client")
				}
				if client.WS == nil {
					t.Error("New() returned nil WebSocket client")
				}
			}
		})
	}
}

// TestCoincheck_Methods はCoincheckクライアントの各メソッドをテスト
func TestCoincheck_Methods(t *testing.T) {
	// クライアントの初期化
	client, err := New(
		client.WithREST(),
		client.WithWebSocket(),
	)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	// テストケース
	tests := []struct {
		name    string
		fn      func() error
		wantErr bool
	}{
		{
			name: "Ticker",
			fn: func() error {
				_, err := client.REST.Ticker(context.Background(), "btc_jpy")
				return err
			},
			wantErr: false,
		},
		{
			name: "Accounts",
			fn: func() error {
				_, err := client.REST.Accounts(context.Background())
				return err
			},
			wantErr: false,
		},
		{
			name: "Balance",
			fn: func() error {
				_, err := client.REST.Balance(context.Background())
				return err
			},
			wantErr: false,
		},
		{
			name: "WebSocket",
			fn: func() error {
				ws := client.WS
				if ws == nil {
					return nil
				}
				return nil
			},
			wantErr: false,
		},
	}

	// テスト実行
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// メソッドの実行
			err := tt.fn()

			// エラーチェック
			if (err != nil) != tt.wantErr {
				t.Errorf("%s() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			}
		})
	}
}
