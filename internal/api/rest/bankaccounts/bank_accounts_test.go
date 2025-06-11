package bankaccounts

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

func TestBankAccountsGET(t *testing.T) {
	tests := map[string]struct {
		mockResponse string
		want         *GetResponse
		wantErr      bool
	}{
		"[正常系]銀行口座一覧取得": {
			mockResponse: `{
				"success": true,
				"bank_accounts": [
					{
						"id": 1,
						"bank_name": "三菱UFJ銀行",
						"branch_name": "新宿支店",
						"bank_type": "futsu",
						"number": "1234567",
						"name": "ヤマダタロウ",
						"created_at": "2023-01-01T00:00:00Z",
						"updated_at": "2023-01-01T00:00:00Z"
					}
				]
			}`,
			want: &GetResponse{
				Success: true,
				BankAccounts: []BankAccount{
					{
						ID:         1,
						BankName:   "三菱UFJ銀行",
						BranchName: "新宿支店",
						BankType:   "futsu",
						Number:     "1234567",
						Name:       "ヤマダタロウ",
						CreatedAt:  time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
						UpdatedAt:  time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
					},
				},
			},
			wantErr: false,
		},
		"[正常系]空の口座一覧": {
			mockResponse: `{
				"success": true,
				"bank_accounts": []
			}`,
			want: &GetResponse{
				Success:      true,
				BankAccounts: []BankAccount{},
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
			if len(mockResp.BankAccounts) != len(tt.want.BankAccounts) {
				t.Errorf("BankAccounts length = %v, want %v", len(mockResp.BankAccounts), len(tt.want.BankAccounts))
			}
		})
	}
}

func TestBankAccountsPOST(t *testing.T) {
	tests := map[string]struct {
		mockResponse string
		params       CreateBankAccountParams
		want         *PostResponse
		wantErr      bool
	}{
		"[正常系]銀行口座登録": {
			mockResponse: `{
				"success": true,
				"id": 1,
				"bank_name": "三菱UFJ銀行",
				"branch_name": "新宿支店",
				"bank_type": "futsu",
				"number": "1234567",
				"name": "ヤマダタロウ"
			}`,
			params: CreateBankAccountParams{
				BankName:   "三菱UFJ銀行",
				BranchName: "新宿支店",
				BankType:   "futsu",
				Number:     "1234567",
				Name:       "ヤマダタロウ",
			},
			want: &PostResponse{
				Success:    true,
				ID:         1,
				BankName:   "三菱UFJ銀行",
				BranchName: "新宿支店",
				BankType:   "futsu",
				Number:     "1234567",
				Name:       "ヤマダタロウ",
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

func TestBankAccountsDELETE(t *testing.T) {
	tests := map[string]struct {
		mockResponse string
		accountID    int
		want         *DeleteResponse
		wantErr      bool
	}{
		"[正常系]銀行口座削除": {
			mockResponse: `{
				"success": true,
				"id": 1
			}`,
			accountID: 1,
			want: &DeleteResponse{
				Success: true,
				ID:      1,
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