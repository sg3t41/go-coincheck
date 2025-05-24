package accounts

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"
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

func TestAccountsGet(t *testing.T) {
	tests := map[string]struct {
		mockResponse string
		want         *GetResponse
		wantErr      bool
	}{
		"[正常系]全部あり": {
			mockResponse: `{
				"success": true,
				"id": 123,
				"email": "user@example.com",
				"identity_status": "verified",
				"bitcoin_address": "1A1zP1...",
				"taker_fee": "0.001",
				"maker_fee": "0.0005",
				"exchange_fees": {
					"btc_jpy": {"maker_fee_rate": "0.0005", "taker_fee_rate": "0.001"},
					"eth_jpy": {"maker_fee_rate": "0.0007", "taker_fee_rate": "0.0012"}
				}
			}`,
			want: &GetResponse{
				Success:        true,
				ID:             123,
				Email:          "user@example.com",
				IdentityStatus: "verified",
				BitcoinAddress: "1A1zP1...",
				TakerFee:       "0.001",
				MakerFee:       "0.0005",
				ExchangeFees: map[string]ExchangeFeeRate{
					"btc_jpy": {MakerFeeRate: "0.0005", TakerFeeRate: "0.001"},
					"eth_jpy": {MakerFeeRate: "0.0007", TakerFeeRate: "0.0012"},
				},
			},
			wantErr: false,
		},
		"[正常系]exchange_fees空": {
			mockResponse: `{
				"success": true,
				"id": 456,
				"email": "user2@example.com",
				"identity_status": "pending",
				"bitcoin_address": "1Q2w3E...",
				"taker_fee": "0.002",
				"maker_fee": "0.001",
				"exchange_fees": {}
			}`,
			want: &GetResponse{
				Success:        true,
				ID:             456,
				Email:          "user2@example.com",
				IdentityStatus: "pending",
				BitcoinAddress: "1Q2w3E...",
				TakerFee:       "0.002",
				MakerFee:       "0.001",
				ExchangeFees:   map[string]ExchangeFeeRate{},
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
			ac := New(client)

			ctx := context.Background()
			got, err := ac.Get(ctx)
			if (err != nil) != tt.wantErr {
				t.Fatalf("error expectation: wantErr=%v, got=%v (err=%v)", tt.wantErr, err != nil, err)
			}
			if tt.wantErr {
				return
			}
			if got == nil {
				t.Fatalf("got == nil, want %+v", tt.want)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("want=%+v, got=%+v", tt.want, got)
			}
		})
	}
}
