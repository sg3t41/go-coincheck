package http

import (
	"net/url"
	"regexp"
	"strings"
	"testing"
)

func TestGenerateRequestHeaders(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		creds   *Credentials
		rawURL  string
		body    string
		wantKey string
		wantErr bool
		errMsg  string
	}{
		{
			name:    "正常系: 正しいキーとシークレット",
			creds:   &Credentials{key: "testkey", secret: "testsecret"},
			rawURL:  "https://api.example.com/foo",
			body:    `{"foo":"bar"}`,
			wantKey: "testkey",
			wantErr: false,
		},
		{
			name:    "異常系: シークレットが空",
			creds:   &Credentials{key: "any", secret: ""},
			rawURL:  "https://api.example.com/foo",
			body:    "body",
			wantKey: "any",
			wantErr: false, // hmac.Writeには空secretでもエラーにならず署名生成はされる
		},
	}

	for _, tt := range tests {
		tt := tt // for t.Parallel
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			u, _ := url.Parse(tt.rawURL)
			headers, err := tt.creds.GenerateRequestHeaders(u, tt.body)
			if (err != nil) != tt.wantErr {
				t.Fatalf("wantErr=%v, got=%v, err=%v", tt.wantErr, err != nil, err)
			}
			if tt.wantErr {
				if err == nil || (tt.errMsg != "" && !strings.Contains(err.Error(), tt.errMsg)) {
					t.Errorf("error message mismatch: want contains %q, got %v", tt.errMsg, err)
				}
				return
			}
			if headers.AccessKey != tt.wantKey {
				t.Errorf("AccessKey: want %q, got %q", tt.wantKey, headers.AccessKey)
			}
			// nonceは数字のみの文字列
			matched, _ := regexp.MatchString(`^\d+$`, headers.AccessNonce)
			if !matched {
				t.Errorf("AccessNonce: want digits, got %q", headers.AccessNonce)
			}
			// 署名はsha256の16進文字列
			matched, _ = regexp.MatchString(`^[0-9a-f]{64}$`, headers.AccessSignature)
			if !matched {
				t.Errorf("AccessSignature: want 64 hex chars, got %q", headers.AccessSignature)
			}
		})
	}
}
