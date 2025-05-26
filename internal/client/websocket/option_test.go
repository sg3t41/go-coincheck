package websocket

import (
	"strings"
	"testing"
)

func TestWithBaseURL(t *testing.T) {
	tests := map[string]struct {
		strURL     string
		wantErr    bool
		errMsgPart string
		wantURL    string
	}{
		"正常系: 正しいURL": {
			strURL:  "wss://ws.example.com",
			wantErr: false,
			wantURL: "wss://ws.example.com",
		},
		"異常系: 空文字列": {
			strURL:     "",
			wantErr:    true,
			errMsgPart: "ベースURLが空",
		},
		"異常系: 不正なURL": {
			strURL:     "://bad-url",
			wantErr:    true,
			errMsgPart: "parse",
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			wsc := &webSocketClient{}
			err := WithBaseURL(tt.strURL)(wsc)
			if (err != nil) != tt.wantErr {
				t.Fatalf("エラー期待値: wantErr=%v, got=%v (err=%v)", tt.wantErr, err != nil, err)
			}
			if tt.wantErr && tt.errMsgPart != "" && (err == nil || !strings.Contains(err.Error(), tt.errMsgPart)) {
				t.Errorf("エラーメッセージ部分一致失敗: 期待 %q, 実際 %v", tt.errMsgPart, err)
			}
			if !tt.wantErr && wsc.url != nil && wsc.url.String() != tt.wantURL {
				t.Errorf("URL: 期待=%q, 実際=%q", tt.wantURL, wsc.url.String())
			}
			if !tt.wantErr && wsc.url == nil {
				t.Errorf("URL: 期待がnilでないのに実際はnil")
			}
		})
	}
}
