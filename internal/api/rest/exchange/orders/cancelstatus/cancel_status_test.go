package cancelstatus

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"

	httpc "github.com/sg3t41/go-coincheck/internal/client/http"
)

/* MOCK ここから */

type mockClient struct {
	mockResponse string
}

func (m *mockClient) CreateRequest(ctx context.Context, input httpc.RequestInput) (*http.Request, error) {
	req, _ := http.NewRequest("GET", "http://example.com", nil)
	return req, nil
}
func (m *mockClient) Do(req *http.Request, output any) error {
	return json.Unmarshal([]byte(m.mockResponse), output)
}
func (m *mockClient) Close() error                      { return nil }
func (m *mockClient) Connect(ctx context.Context) error { panic("not implemented") }
func (m *mockClient) Subscribe(ctx context.Context, channel string, in chan<- string) error {
	panic("not implemented")
}

/* MOCK ここまで */

func TestCancelStatusGET(t *testing.T) {
	tests := map[string]struct {
		mockResponse string
		want         *GetResponse
		wantErr      bool
	}{
		"[正常系]キャンセル済み": {
			mockResponse: `{
				"success": true,
				"id": 12345,
				"cancel": true,
				"created_at": "2022-01-01T12:00:00Z"
			}`,
			want: &GetResponse{
				Success:   true,
				ID:        12345,
				Cancel:    true,
				CreatedAt: "2022-01-01T12:00:00Z",
			},
			wantErr: false,
		},
		"[正常系]未キャンセル": {
			mockResponse: `{
				"success": true,
				"id": 67890,
				"cancel": false,
				"created_at": "2022-01-01T13:00:00Z"
			}`,
			want: &GetResponse{
				Success:   true,
				ID:        67890,
				Cancel:    false,
				CreatedAt: "2022-01-01T13:00:00Z",
			},
			wantErr: false,
		},
		"[異常系]JSONの形式が不正": {
			mockResponse: `{invalid json}`,
			want:         nil,
			wantErr:      true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			client := &mockClient{mockResponse: tt.mockResponse}
			cstatus := New(client)

			ctx := context.Background()
			got, err := cstatus.GET(ctx, 12345)
			if (err != nil) != tt.wantErr {
				t.Fatalf("error expectation: wantErr=%v, got=%v (err=%v)", tt.wantErr, err != nil, err)
			}
			if tt.wantErr {
				return
			}
			if got == nil {
				t.Fatalf("got == nil, want %+v", tt.want)
			}
			if got.Success != tt.want.Success {
				t.Errorf("Success: want=%v, got=%v", tt.want.Success, got.Success)
			}
			if got.ID != tt.want.ID {
				t.Errorf("ID: want=%v, got=%v", tt.want.ID, got.ID)
			}
			if got.Cancel != tt.want.Cancel {
				t.Errorf("Cancel: want=%v, got=%v", tt.want.Cancel, got.Cancel)
			}
			if got.CreatedAt != tt.want.CreatedAt {
				t.Errorf("CreatedAt: want=%v, got=%v", tt.want.CreatedAt, got.CreatedAt)
			}
		})
	}
}
