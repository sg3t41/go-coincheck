package coincheck

import (
	"testing"
)

// テスト用ダミーrest/wsパッケージを使う場合は適宜importパスを変えてください。

func TestWithCredentials(t *testing.T) {
	tests := map[string]struct {
		key     string
		secret  string
		wantKey string
		wantSec string
	}{
		"正常系: key,secretをセット": {
			key:     "testkey",
			secret:  "testsecret",
			wantKey: "testkey",
			wantSec: "testsecret",
		},
		"正常系: 空文字列も通る": {
			key:     "",
			secret:  "",
			wantKey: "",
			wantSec: "",
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			c := &Coincheck{}
			err := WithCredentials(tt.key, tt.secret)(c)
			if err != nil {
				t.Fatalf("WithCredentials エラー: %v", err)
			}
			if c.credentials.key != tt.wantKey || c.credentials.secret != tt.wantSec {
				t.Errorf("credentials: 期待 key=%q, secret=%q, 実際 key=%q, secret=%q",
					tt.wantKey, tt.wantSec, c.credentials.key, c.credentials.secret)
			}
		})
	}
}

func TestWithREST(t *testing.T) {
	tests := map[string]struct {
		key     string
		secret  string
		wantErr bool
	}{
		"正常系: 通常のkey/secretでREST生成": {
			key:     "k",
			secret:  "s",
			wantErr: false,
		},
		// 異常系はrest.New側でテストするのが普通なので省略可
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			c := &Coincheck{credentials: credentials{key: tt.key, secret: tt.secret}}
			err := WithREST()(c)
			if (err != nil) != tt.wantErr {
				t.Fatalf("エラー期待値: wantErr=%v, got=%v (err=%v)", tt.wantErr, err != nil, err)
			}
			if !tt.wantErr && c.REST == nil {
				t.Errorf("RESTクライアントがセットされていません")
			}
		})
	}
}

func TestWithWebSocket(t *testing.T) {
	tests := map[string]struct {
		wantErr bool
	}{
		"正常系: WebSocket生成": {
			wantErr: false,
		},
		// 異常系はws.New側でテストするのが普通なので省略可
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			c := &Coincheck{}
			err := WithWebSocket()(c)
			if (err != nil) != tt.wantErr {
				t.Fatalf("エラー期待値: wantErr=%v, got=%v (err=%v)", tt.wantErr, err != nil, err)
			}
			if !tt.wantErr && c.WS == nil {
				t.Errorf("WSクライアントがセットされていません")
			}
		})
	}
}
