package http

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

// モックサーバーのレスポンスデータを定義
type mockResponse struct {
	Message string `json:"message"`
}

func TestCreateRequest(t *testing.T) {
	client := &httpClient{
		baseURL: &url.URL{Scheme: "https", Host: "coincheck.com"},
	}

	input := RequestInput{
		Method: "GET",
		Path:   "/api/test",
		QueryParam: map[string]string{
			"key": "value",
		},
	}

	ctx := context.Background()
	req, err := client.CreateRequest(ctx, input)
	if err != nil {
		t.Fatalf("CreateRequest failed: %v", err)
	}

	if req.URL.String() != "https://coincheck.com/api/test?key=value" {
		t.Errorf("unexpected URL: %s", req.URL.String())
	}

	if req.Method != "GET" {
		t.Errorf("unexpected method: %s", req.Method)
	}
}

func TestDo(t *testing.T) {
	// モックサーバーを作成
	mockResp := mockResponse{Message: "Success"}
	mockRespBytes, _ := json.Marshal(mockResp)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(mockRespBytes)
	}))
	defer server.Close()

	client := &httpClient{
		httpClient: http.DefaultClient,
		baseURL:    &url.URL{Scheme: "https", Host: server.URL},
	}

	req, err := http.NewRequest("GET", server.URL, nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	var output mockResponse
	err = client.Do(req, &output)
	if err != nil {
		t.Fatalf("Do failed: %v", err)
	}

	if output.Message != "Success" {
		t.Errorf("unexpected response: %v", output.Message)
	}
}

func TestSetAuthHeaders(t *testing.T) {
	creds := &credentials{
		key:    "test-key",
		secret: "test-secret",
	}

	client := &httpClient{
		credentials: creds,
	}

	req, err := http.NewRequest("GET", "https://example.com", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	err = client.setAuthHeaders(req, "")
	if err != nil {
		t.Fatalf("setAuthHeaders failed: %v", err)
	}

	if req.Header.Get("ACCESS-KEY") != "test-key" {
		t.Errorf("unexpected ACCESS-KEY: %s", req.Header.Get("ACCESS-KEY"))
	}

	// 他のヘッダーも同様に検証可能
}
