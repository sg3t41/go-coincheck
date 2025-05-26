package http

import (
	"strings"
	"testing"
)

func TestWithCredentials(t *testing.T) {
	tests := []struct {
		name      string
		key       string
		secret    string
		wantErr   bool
		errMsg    string
		wantCreds bool
	}{
		{
			name:      "正常系: key, secret両方あり",
			key:       "k",
			secret:    "s",
			wantErr:   false,
			wantCreds: true,
		},
		{
			name:    "エラー: keyもsecretも空",
			key:     "",
			secret:  "",
			wantErr: true,
			errMsg:  "keyとsecretが空",
		},
		{
			name:    "エラー: keyが空",
			key:     "",
			secret:  "s",
			wantErr: true,
			errMsg:  "keyが空",
		},
		{
			name:    "エラー: secretが空",
			key:     "k",
			secret:  "",
			wantErr: true,
			errMsg:  "secretが空",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			hc := &httpClient{}
			err := WithCredentials(tt.key, tt.secret)(hc)
			if (err != nil) != tt.wantErr {
				t.Fatalf("wantErr=%v, got=%v, err=%v", tt.wantErr, err != nil, err)
			}
			if tt.wantErr && tt.errMsg != "" && (err == nil || !strings.Contains(err.Error(), tt.errMsg)) {
				t.Errorf("error message mismatch: want contains %q, got %v", tt.errMsg, err)
			}
			if tt.wantCreds {
				if hc.credentials == nil {
					t.Errorf("credentials: want not nil, got nil")
				} else if hc.credentials.key != tt.key || hc.credentials.secret != tt.secret {
					t.Errorf("credentials: want key=%q, secret=%q, got key=%q, secret=%q",
						tt.key, tt.secret, hc.credentials.key, hc.credentials.secret)
				}
			} else if hc.credentials != nil {
				t.Errorf("credentials: want nil, got not nil")
			}
		})
	}
}

func TestWithBaseURL(t *testing.T) {
	tests := []struct {
		name    string
		baseURL string
		wantErr bool
		errMsg  string
	}{
		{
			name:    "正常系: 正しいURL",
			baseURL: "https://api.example.com",
			wantErr: false,
		},
		{
			name:    "エラー: 空文字列",
			baseURL: "",
			wantErr: true,
			errMsg:  "ベースURLが空",
		},
		{
			name:    "エラー: 不正なURL",
			baseURL: "://bad-url",
			wantErr: true,
			errMsg:  "parse",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			hc := &httpClient{}
			err := WithBaseURL(tt.baseURL)(hc)
			if (err != nil) != tt.wantErr {
				t.Fatalf("wantErr=%v, got=%v, err=%v", tt.wantErr, err != nil, err)
			}
			if tt.wantErr && tt.errMsg != "" && (err == nil || !strings.Contains(err.Error(), tt.errMsg)) {
				t.Errorf("error message mismatch: want contains %q, got %v", tt.errMsg, err)
			}
			if !tt.wantErr {
				if hc.baseURL == nil {
					t.Errorf("baseURL: want not nil, got nil")
				} else if hc.baseURL.String() != tt.baseURL {
					t.Errorf("baseURL: want %q, got %q", tt.baseURL, hc.baseURL.String())
				}
			}
		})
	}
}
