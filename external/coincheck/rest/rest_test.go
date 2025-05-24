package rest

import (
	"context"
	"testing"
)

// TestNew はRESTクライアントの初期化をテスト
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
			client, err := New(tt.key, tt.secret)

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
			}
		})
	}
}

// TestREST_Interface はRESTインターフェースの実装をテスト
func TestREST_Interface(t *testing.T) {
	// クライアントの初期化
	client, err := New("valid-key", "valid-secret")
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}

	// インターフェースの実装チェック
	var _ REST = client
}

// TestREST_Methods はRESTクライアントの各メソッドをテスト
func TestREST_Methods(t *testing.T) {
	// クライアントの初期化
	client, err := New("valid-key", "valid-secret")
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
				_, err := client.Ticker(context.Background(), "btc_jpy")
				return err
			},
			wantErr: false,
		},
		{
			name: "Accounts",
			fn: func() error {
				_, err := client.Accounts(context.Background())
				return err
			},
			wantErr: false,
		},
		{
			name: "Balance",
			fn: func() error {
				_, err := client.Balance(context.Background())
				return err
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
