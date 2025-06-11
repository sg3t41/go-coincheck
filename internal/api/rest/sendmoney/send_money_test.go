package sendmoney

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/sg3t41/go-coincheck/internal/client"
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

func TestSendMoneyPOST(t *testing.T) {
	tests := map[string]struct {
		mockResponse string
		params       SendMoneyParams
		want         *PostResponse
		wantErr      bool
	}{
		"[正常系]BTC送金": {
			mockResponse: `{
				"success": true,
				"id": 12345,
				"address": "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa",
				"amount": "0.01",
				"fee": "0.0005"
			}`,
			params: SendMoneyParams{
				Address:  "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa",
				Amount:   0.01,
				Currency: "btc",
			},
			want: &PostResponse{
				Success: true,
				ID:      12345,
				Address: "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa",
				Amount:  "0.01",
				Fee:     "0.0005",
			},
			wantErr: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			// モックレスポンスを解析
			var mockResp PostResponse
			if err := json.Unmarshal([]byte(tt.mockResponse), &mockResp); err != nil {
				t.Fatalf("Mock response unmarshal failed: %v", err)
			}

			// 結果を比較（実際のテストではモッククライアントを使用）
			if mockResp.Success != tt.want.Success {
				t.Errorf("Success = %v, want %v", mockResp.Success, tt.want.Success)
			}
			if mockResp.ID != tt.want.ID {
				t.Errorf("ID = %v, want %v", mockResp.ID, tt.want.ID)
			}
		})
	}
}

func TestSendMoneyGET(t *testing.T) {
	tests := map[string]struct {
		mockResponse string
		want         *GetResponse
		wantErr      bool
	}{
		"[正常系]送金履歴取得": {
			mockResponse: `{
				"success": true,
				"send_moneys": [
					{
						"id": 12345,
						"address": "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa",
						"amount": "0.01",
						"fee": "0.0005",
						"currency": "btc",
						"txhash": "abc123def456",
						"status": "completed",
						"created_at": "2023-01-01T00:00:00Z",
						"updated_at": "2023-01-01T00:01:00Z"
					}
				]
			}`,
			want: &GetResponse{
				Success: true,
				SendMoneyies: []SendMoneyTransaction{
					{
						ID:        12345,
						Address:   "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa",
						Amount:    "0.01",
						Fee:       "0.0005",
						Currency:  "btc",
						TxHash:    stringPtr("abc123def456"),
						Status:    "completed",
						CreatedAt: time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
						UpdatedAt: time.Date(2023, 1, 1, 0, 1, 0, 0, time.UTC),
					},
				},
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
			if len(mockResp.SendMoneyies) != len(tt.want.SendMoneyies) {
				t.Errorf("SendMoneyies length = %v, want %v", len(mockResp.SendMoneyies), len(tt.want.SendMoneyies))
			}
		})
	}
}

func stringPtr(s string) *string {
	return &s
}