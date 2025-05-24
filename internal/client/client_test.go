package client

import (
	"testing"
)

// TestNew_WithREST はRESTクライアントのみを初期化するテスト
func TestNew_WithREST(t *testing.T) {
	// テストケース
	tests := []struct {
		TEST_NAME  string
		KEY        string
		SECRET     string
		WANT_ERROR bool
	}{
		{
			TEST_NAME:  "正常系: 有効な認証情報",
			KEY:        "valid-key",
			SECRET:     "valid-secret",
			WANT_ERROR: false,
		},
		{
			TEST_NAME:  "異常系: 空の認証情報",
			KEY:        "",
			SECRET:     "",
			WANT_ERROR: true,
		},
	}

	// テスト実行
	for _, tt := range tests {
		t.Run(tt.TEST_NAME, func(t *testing.T) {
			// クライアントの初期化
			client, err := New(
				WithREST(tt.KEY, tt.SECRET),
			)

			// エラーチェック
			if (err != nil) != tt.WANT_ERROR {
				t.Errorf("New() error = %v, wantErr %v", err, tt.WANT_ERROR)
				return
			}

			// 正常系の場合、クライアントが正しく初期化されているか確認
			if !tt.WANT_ERROR {
				if client == nil {
					t.Error("New() returned nil client")
				}
			}
		})
	}
}

// TestNew_WithWebSocket はWebSocketクライアントのみを初期化するテスト
func TestNew_WithWebSocket(t *testing.T) {
	// クライアントの初期化
	client, err := New(
		WithWebSocket(),
	)

	// エラーチェック
	if err != nil {
		t.Errorf("New() error = %v", err)
		return
	}

	// クライアントが正しく初期化されているか確認
	if client == nil {
		t.Error("New() returned nil client")
	}
}

// TestNew_WithBoth はRESTとWebSocketの両方のクライアントを初期化するテスト
func TestNew_WithBoth(t *testing.T) {
	// クライアントの初期化
	client, err := New(
		WithREST("valid-key", "valid-secret"),
		WithWebSocket(),
	)

	// エラーチェック
	if err != nil {
		t.Errorf("New() error = %v", err)
		return
	}

	// クライアントが正しく初期化されているか確認
	if client == nil {
		t.Error("New() returned nil client")
	}
}

// TestNew_WithNoOptions はオプションなしでクライアントを初期化するテスト
func TestNew_WithNoOptions(t *testing.T) {
	// クライアントの初期化
	client, err := New()

	// エラーチェック
	if err != nil {
		t.Errorf("New() error = %v", err)
		return
	}

	// クライアントが正しく初期化されているか確認
	if client == nil {
		t.Error("New() returned nil client")
	}
}
