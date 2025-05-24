package opens

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

func strPtr(s string) *string { return &s }

func TestOpensGET(t *testing.T) {
	tests := map[string]struct {
		mockResponse string
		want         *GetResponse
		wantErr      bool
	}{
		"[正常系]複数注文": {
			mockResponse: `{
				"success": true,
				"orders": [
					{
						"id": 1,
						"order_type": "buy",
						"rate": "1234567",
						"pair": "btc_jpy",
						"pending_amount": "0.01",
						"pending_market_buy_amount": null,
						"stop_loss_rate": null,
						"created_at": "2022-01-01T12:00:00Z"
					},
					{
						"id": 2,
						"order_type": "sell",
						"rate": null,
						"pair": "btc_jpy",
						"pending_amount": null,
						"pending_market_buy_amount": "10000",
						"stop_loss_rate": "1200000",
						"created_at": "2022-01-01T13:00:00Z"
					}
				]
			}`,
			want: &GetResponse{
				Success: true,
				Orders: []Order{
					{
						ID:                     1,
						OrderType:              "buy",
						Rate:                   strPtr("1234567"),
						Pair:                   "btc_jpy",
						PendingAmount:          strPtr("0.01"),
						PendingMarketBuyAmount: nil,
						StopLossRate:           nil,
						CreatedAt:              "2022-01-01T12:00:00Z",
					},
					{
						ID:                     2,
						OrderType:              "sell",
						Rate:                   nil,
						Pair:                   "btc_jpy",
						PendingAmount:          nil,
						PendingMarketBuyAmount: strPtr("10000"),
						StopLossRate:           strPtr("1200000"),
						CreatedAt:              "2022-01-01T13:00:00Z",
					},
				},
			},
			wantErr: false,
		},
		"[正常系]空配列": {
			mockResponse: `{
				"success": true,
				"orders": []
			}`,
			want: &GetResponse{
				Success: true,
				Orders:  []Order{},
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
			opens := New(client)

			ctx := context.Background()
			got, err := opens.GET(ctx)
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
			if len(got.Orders) != len(tt.want.Orders) {
				t.Errorf("Orders length: want=%d, got=%d", len(tt.want.Orders), len(got.Orders))
			}
			for i := range got.Orders {
				if got.Orders[i].ID != tt.want.Orders[i].ID {
					t.Errorf("Orders[%d].ID: want=%v, got=%v", i, tt.want.Orders[i].ID, got.Orders[i].ID)
				}
				if got.Orders[i].OrderType != tt.want.Orders[i].OrderType {
					t.Errorf("Orders[%d].OrderType: want=%v, got=%v", i, tt.want.Orders[i].OrderType, got.Orders[i].OrderType)
				}
				if !equalStrPtr(got.Orders[i].Rate, tt.want.Orders[i].Rate) {
					t.Errorf("Orders[%d].Rate: want=%v, got=%v", i, tt.want.Orders[i].Rate, got.Orders[i].Rate)
				}
				if got.Orders[i].Pair != tt.want.Orders[i].Pair {
					t.Errorf("Orders[%d].Pair: want=%v, got=%v", i, tt.want.Orders[i].Pair, got.Orders[i].Pair)
				}
				if !equalStrPtr(got.Orders[i].PendingAmount, tt.want.Orders[i].PendingAmount) {
					t.Errorf("Orders[%d].PendingAmount: want=%v, got=%v", i, tt.want.Orders[i].PendingAmount, got.Orders[i].PendingAmount)
				}
				if !equalStrPtr(got.Orders[i].PendingMarketBuyAmount, tt.want.Orders[i].PendingMarketBuyAmount) {
					t.Errorf("Orders[%d].PendingMarketBuyAmount: want=%v, got=%v", i, tt.want.Orders[i].PendingMarketBuyAmount, got.Orders[i].PendingMarketBuyAmount)
				}
				if !equalStrPtr(got.Orders[i].StopLossRate, tt.want.Orders[i].StopLossRate) {
					t.Errorf("Orders[%d].StopLossRate: want=%v, got=%v", i, tt.want.Orders[i].StopLossRate, got.Orders[i].StopLossRate)
				}
				if got.Orders[i].CreatedAt != tt.want.Orders[i].CreatedAt {
					t.Errorf("Orders[%d].CreatedAt: want=%v, got=%v", i, tt.want.Orders[i].CreatedAt, got.Orders[i].CreatedAt)
				}
			}
		})
	}
}

func equalStrPtr(a, b *string) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}
