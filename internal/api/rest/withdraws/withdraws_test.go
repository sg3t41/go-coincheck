package withdraws

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

func TestWithdrawsGET(t *testing.T) {
	tests := map[string]struct {
		mockResponse string
		want         *GetResponse
		wantErr      bool
	}{
		"[正常系]出金履歴取得": {
			mockResponse: `{
				"success": true,
				"withdraws": [
					{
						"id": 12345,
						"status": "completed",
						"amount": "100000",
						"currency": "jpy",
						"fee": "400",
						"is_fast": false,
						"bank_account_id": 1,
						"created_at": "2023-01-01T00:00:00Z",
						"updated_at": "2023-01-01T01:00:00Z"
					}
				]
			}`,
			want: &GetResponse{
				Success: true,
				Withdraws: []WithdrawTransaction{
					{
						ID:            12345,
						Status:        "completed",
						Amount:        "100000",
						Currency:      "jpy",
						Fee:           "400",
						IsFast:        false,
						BankAccountID: 1,
						CreatedAt:     time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
						UpdatedAt:     time.Date(2023, 1, 1, 1, 0, 0, 0, time.UTC),
					},
				},
			},
			wantErr: false,
		},
		"[正常系]空の履歴": {
			mockResponse: `{
				"success": true,
				"withdraws": []
			}`,
			want: &GetResponse{
				Success:   true,
				Withdraws: []WithdrawTransaction{},
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
			if len(mockResp.Withdraws) != len(tt.want.Withdraws) {
				t.Errorf("Withdraws length = %v, want %v", len(mockResp.Withdraws), len(tt.want.Withdraws))
			}
		})
	}
}

func TestWithdrawsPOST(t *testing.T) {
	tests := map[string]struct {
		mockResponse string
		params       CreateWithdrawParams
		want         *PostResponse
		wantErr      bool
	}{
		"[正常系]通常出金申請": {
			mockResponse: `{
				"success": true,
				"id": 12345,
				"amount": "100000",
				"currency": "jpy",
				"fee": "400",
				"is_fast": false,
				"bank_account_id": 1,
				"status": "pending"
			}`,
			params: CreateWithdrawParams{
				BankAccountID: 1,
				Amount:        "100000",
				Currency:      "jpy",
				IsFast:        false,
			},
			want: &PostResponse{
				Success:       true,
				ID:            12345,
				Amount:        "100000",
				Currency:      "jpy",
				Fee:           "400",
				IsFast:        false,
				BankAccountID: 1,
				Status:        "pending",
			},
			wantErr: false,
		},
		"[正常系]高速出金申請": {
			mockResponse: `{
				"success": true,
				"id": 12346,
				"amount": "50000",
				"currency": "jpy",
				"fee": "770",
				"is_fast": true,
				"bank_account_id": 1,
				"status": "pending"
			}`,
			params: CreateWithdrawParams{
				BankAccountID: 1,
				Amount:        "50000",
				Currency:      "jpy",
				IsFast:        true,
			},
			want: &PostResponse{
				Success:       true,
				ID:            12346,
				Amount:        "50000",
				Currency:      "jpy",
				Fee:           "770",
				IsFast:        true,
				BankAccountID: 1,
				Status:        "pending",
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

func TestWithdrawsDELETE(t *testing.T) {
	tests := map[string]struct {
		mockResponse string
		withdrawID   int
		want         *DeleteResponse
		wantErr      bool
	}{
		"[正常系]出金申請キャンセル": {
			mockResponse: `{
				"success": true,
				"id": 12345
			}`,
			withdrawID: 12345,
			want: &DeleteResponse{
				Success: true,
				ID:      12345,
			},
			wantErr: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			// モックレスポンスを解析
			var mockResp DeleteResponse
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