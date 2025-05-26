package http

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

type mockRoundTripper struct {
	resp *http.Response
	err  error
}

func (m *mockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.resp, m.err
}

// --- テスト用のエラーを返すio.Reader ---
type errReader struct{}

func (errReader) Read([]byte) (int, error) { return 0, fmt.Errorf("テスト用エラー") }
func (errReader) Close() error             { return nil }

func TestCreateRequest(t *testing.T) {
	base, _ := url.Parse("https://api.example.com")
	c := &httpClient{
		baseURL: base,
	}

	tests := []struct {
		name       string
		input      RequestInput
		wantMethod string
		wantURL    string
		wantBody   string
		wantErr    bool
	}{
		{
			name:       "GETリクエスト",
			input:      RequestInput{Method: "GET", Path: "/api/hello"},
			wantMethod: "GET",
			wantURL:    "https://api.example.com/api/hello",
			wantBody:   "",
			wantErr:    false,
		},
		{
			name: "POSTリクエスト with body and query",
			input: RequestInput{
				Method:     "POST",
				Path:       "/api/post",
				Body:       strings.NewReader(`{"foo":"bar"}`),
				QueryParam: map[string]string{"x": "y"},
			},
			wantMethod: "POST",
			wantURL:    "https://api.example.com/api/post?x=y",
			wantBody:   `{"foo":"bar"}`,
			wantErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := c.CreateRequest(context.Background(), tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("error: wantErr=%v, got=%v (err=%v)", tt.wantErr, err != nil, err)
			}
			if tt.wantErr {
				return
			}
			if req.Method != tt.wantMethod {
				t.Errorf("Method: got=%s, want=%s", req.Method, tt.wantMethod)
			}
			if req.URL.String() != tt.wantURL {
				t.Errorf("URL: got=%s, want=%s", req.URL.String(), tt.wantURL)
			}
			if req.Body != nil {
				b, _ := io.ReadAll(req.Body)
				if string(b) != tt.wantBody {
					t.Errorf("Body: got=%q, want=%q", string(b), tt.wantBody)
				}
			}
		})
	}
}

func TestDo(t *testing.T) {
	type testCase struct {
		name       string
		resp       *http.Response
		roundErr   error
		wantOutput map[string]any
		wantErr    bool
		errMsg     string
	}

	tests := []testCase{
		{
			name: "正常系: 200 OK",
			resp: &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader(`{"success":true,"msg":"ok"}`)),
			},
			wantOutput: map[string]any{"success": true, "msg": "ok"},
			wantErr:    false,
		},
		{
			name: "異常系: ステータスコード500",
			resp: &http.Response{
				StatusCode: 500,
				Body:       io.NopCloser(strings.NewReader(`{"error":"fail"}`)),
			},
			wantErr: true,
			errMsg:  "予期しないステータスコード",
		},
		{
			name:     "異常系: RoundTripエラー",
			resp:     nil,
			roundErr: fmt.Errorf("transport fail"),
			wantErr:  true,
			errMsg:   "transport fail",
		},
		{
			name: "異常系: レスポンスボディが壊れている",
			resp: &http.Response{
				StatusCode: 200,
				Body:       errReader{}, // io.ReadAllでエラーを返す
			},
			wantErr: true,
			errMsg:  "テスト用エラー",
		},
		{
			name: "異常系: JSONデコードエラー",
			resp: &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(strings.NewReader(`invalid json`)),
			},
			wantErr: true,
			errMsg:  "invalid character",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &httpClient{
				httpClient: &http.Client{
					Transport: &mockRoundTripper{resp: tt.resp, err: tt.roundErr},
				},
			}
			req, _ := http.NewRequest("GET", "http://example.com", nil)
			var got map[string]any
			err := c.Do(req, &got)
			if (err != nil) != tt.wantErr {
				t.Fatalf("wantErr=%v, got=%v, err=%v", tt.wantErr, err != nil, err)
			}
			if tt.wantErr && err != nil && tt.errMsg != "" && !strings.Contains(err.Error(), tt.errMsg) {
				t.Errorf("error message mismatch, want contains %q, got %v", tt.errMsg, err)
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.wantOutput) {
				t.Errorf("output mismatch, want=%v, got=%v", tt.wantOutput, got)
			}
		})
	}
}
