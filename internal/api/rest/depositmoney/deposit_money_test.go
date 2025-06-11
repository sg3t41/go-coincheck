package depositmoney

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
	"time"
)

/* MOCK ここから */

type mockClient struct{}

func (m *mockClient) CreateRequest(ctx context.Context, input interface{}) (*http.Request, error) {
	panic("not implemented")
}

func (m *mockClient) Do(req *http.Request, v interface{}) error {
	panic("not implemented")
}

/* MOCK ここまで */

func TestDepositMoneyGET(t *testing.T) {
	tests := map[string]struct {
		mockResponse string
		want         *GetResponse
		wantErr      bool
	}{
		"[正常系]入金履歴取得": {
			mockResponse: `{
				"success": true,
				"deposits": [
					{
						"id": 12345,
						"amount": "0.01",
						"currency": "btc",
						"address": "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa",
						"txhash": "abc123def456",
						"status": "confirmed",
						"created_at": "2023-01-01T00:00:00Z",
						"updated_at": "2023-01-01T00:01:00Z"
					}
				]
			}`,
			want: &GetResponse{
				Success: true,
				Deposits: []DepositTransaction{
					{
						ID:        12345,
						Amount:    "0.01",
						Currency:  "btc",
						Address:   "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa",
						TxHash:    stringPtr("abc123def456"),
						Status:    "confirmed",
						CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
						UpdatedAt: time.Date(2023, 1, 1, 0, 1, 0, 0, time.UTC),
					},
				},
			},
			wantErr: false,
		},
		"[正常系]空の履歴": {
			mockResponse: `{
				"success": true,
				"deposits": []
			}`,
			want: &GetResponse{
				Success:  true,
				Deposits: []DepositTransaction{},
			},
			wantErr: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			// モックレスポンスを解析
			var mockResp GetResponse
			if err := json.Unmarshal([]byte(tt.mockResponse), &mockResp); err != nil {
				t.Fatalf("Mock response unmarshal failed: %v", err)
			}

			// 基本的な検証
			if mockResp.Success != tt.want.Success {
				t.Errorf("Success = %v, want %v", mockResp.Success, tt.want.Success)
			}
			if len(mockResp.Deposits) != len(tt.want.Deposits) {
				t.Errorf("Deposits length = %v, want %v", len(mockResp.Deposits), len(tt.want.Deposits))
			}
		})
	}
}

func TestDepositMoneyPostFast(t *testing.T) {
	tests := map[string]struct {
		mockResponse string
		depositID    int
		want         *PostFastResponse
		wantErr      bool
	}{
		"[正常系]高速入金成功": {
			mockResponse: `{
				"success": true,
				"id": 12345,
				"status": "processing"
			}`,
			depositID: 12345,
			want: &PostFastResponse{
				Success: true,
				ID:      12345,
				Status:  "processing",
			},
			wantErr: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			// モックレスポンスを解析
			var mockResp PostFastResponse
			if err := json.Unmarshal([]byte(tt.mockResponse), &mockResp); err != nil {
				t.Fatalf("Mock response unmarshal failed: %v", err)
			}

			// 基本的な検証
			if mockResp.Success != tt.want.Success {
				t.Errorf("Success = %v, want %v", mockResp.Success, tt.want.Success)
			}
			if mockResp.ID != tt.want.ID {
				t.Errorf("ID = %v, want %v", mockResp.ID, tt.want.ID)
			}
		})
	}
}

func stringPtr(s string) *string {
	return &s
}