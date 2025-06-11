package transactions

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
	"time"

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

func mustParseTime(t *testing.T, s string) time.Time {
	tm, err := time.Parse(time.RFC3339, s)
	if err != nil {
		t.Fatalf("invalid time: %v", err)
	}
	return tm
}

func TestTransactionsGET(t *testing.T) {
	tests := map[string]struct {
		mockResponse string
		want         *GetResponse
		wantErr      bool
	}{
		"[正常系]複数取引": {
			mockResponse: `{
				"success": true,
				"transactions": [
					{
						"id": 1,
						"order_id": 101,
						"created_at": "2022-01-01T12:00:00Z",
						"funds": { "btc": "0.01", "jpy": "-50000" },
						"pair": "btc_jpy",
						"rate": "5000000",
						"fee_currency": "jpy",
						"fee": "10",
						"liquidity": "T",
						"side": "buy"
					},
					{
						"id": 2,
						"order_id": 102,
						"created_at": "2022-01-01T13:00:00Z",
						"funds": { "btc": "-0.01", "jpy": "51000" },
						"pair": "btc_jpy",
						"rate": "5100000",
						"fee_currency": "btc",
						"fee": "0.0001",
						"liquidity": "M",
						"side": "sell"
					}
				]
			}`,
			want: &GetResponse{
				Success: true,
				Transactions: []Transaction{
					{
						ID:          1,
						OrderID:     101,
						CreatedAt:   mustParseTime(t, "2022-01-01T12:00:00Z"),
						Funds:       Funds{BTC: "0.01", JPY: "-50000"},
						Pair:        "btc_jpy",
						Rate:        "5000000",
						FeeCurrency: "jpy",
						Fee:         "10",
						Liquidity:   "T",
						Side:        "buy",
					},
					{
						ID:          2,
						OrderID:     102,
						CreatedAt:   mustParseTime(t, "2022-01-01T13:00:00Z"),
						Funds:       Funds{BTC: "-0.01", JPY: "51000"},
						Pair:        "btc_jpy",
						Rate:        "5100000",
						FeeCurrency: "btc",
						Fee:         "0.0001",
						Liquidity:   "M",
						Side:        "sell",
					},
				},
			},
			wantErr: false,
		},
		"[正常系]空配列": {
			mockResponse: `{
				"success": true,
				"transactions": []
			}`,
			want: &GetResponse{
				Success:      true,
				Transactions: []Transaction{},
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
			transactions := New(client)

			ctx := context.Background()
			got, err := transactions.GET(ctx)
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
			if len(got.Transactions) != len(tt.want.Transactions) {
				t.Errorf("Transactions length: want=%d, got=%d", len(tt.want.Transactions), len(got.Transactions))
			}
			for i := range got.Transactions {
				wantTx := tt.want.Transactions[i]
				gotTx := got.Transactions[i]
				if gotTx.ID != wantTx.ID {
					t.Errorf("Transactions[%d].ID: want=%v, got=%v", i, wantTx.ID, gotTx.ID)
				}
				if gotTx.OrderID != wantTx.OrderID {
					t.Errorf("Transactions[%d].OrderID: want=%v, got=%v", i, wantTx.OrderID, gotTx.OrderID)
				}
				if !gotTx.CreatedAt.Equal(wantTx.CreatedAt) {
					t.Errorf("Transactions[%d].CreatedAt: want=%v, got=%v", i, wantTx.CreatedAt, gotTx.CreatedAt)
				}
				if !reflect.DeepEqual(gotTx.Funds, wantTx.Funds) {
					t.Errorf("Transactions[%d].Funds: want=%v, got=%v", i, wantTx.Funds, gotTx.Funds)
				}
				if gotTx.Pair != wantTx.Pair {
					t.Errorf("Transactions[%d].Pair: want=%v, got=%v", i, wantTx.Pair, gotTx.Pair)
				}
				if gotTx.Rate != wantTx.Rate {
					t.Errorf("Transactions[%d].Rate: want=%v, got=%v", i, wantTx.Rate, gotTx.Rate)
				}
				if gotTx.FeeCurrency != wantTx.FeeCurrency {
					t.Errorf("Transactions[%d].FeeCurrency: want=%v, got=%v", i, wantTx.FeeCurrency, gotTx.FeeCurrency)
				}
				if gotTx.Fee != wantTx.Fee {
					t.Errorf("Transactions[%d].Fee: want=%v, got=%v", i, wantTx.Fee, gotTx.Fee)
				}
				if gotTx.Liquidity != wantTx.Liquidity {
					t.Errorf("Transactions[%d].Liquidity: want=%v, got=%v", i, wantTx.Liquidity, gotTx.Liquidity)
				}
				if gotTx.Side != wantTx.Side {
					t.Errorf("Transactions[%d].Side: want=%v, got=%v", i, wantTx.Side, gotTx.Side)
				}
			}
		})
	}
}
