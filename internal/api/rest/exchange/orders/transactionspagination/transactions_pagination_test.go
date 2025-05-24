package transactionspagination

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/sg3t41/go-coincheck/internal/api/rest/exchange/orders/transactions"
	httpc "github.com/sg3t41/go-coincheck/internal/client/http"
)

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

func mustParseTime(t *testing.T, s string) time.Time {
	tm, err := time.Parse(time.RFC3339, s)
	if err != nil {
		t.Fatalf("invalid time: %v", err)
	}
	return tm
}

func intPtr(v int) *int { return &v }

func TestTransactionsPaginationGET(t *testing.T) {
	tests := map[string]struct {
		limit         int
		order         string
		startingAfter *int
		endingBefore  *int
		mockResponse  string
		want          *GetResponse
		wantErr       bool
	}{
		"[正常系]複数取引": {
			limit: 2, order: "desc", startingAfter: intPtr(10), endingBefore: intPtr(20),
			mockResponse: `{
				"success": true,
				"pagination": {
					"limit": 2,
					"order": "desc",
					"starting_after": 10,
					"ending_before": 20
				},
				"data": [
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
				Pagination: PaginationOrder{
					Limit:         2,
					Order:         "desc",
					StartingAfter: intPtr(10),
					EndingBefore:  intPtr(20),
				},
				Data: []transactions.Transaction{
					{
						ID:          1,
						OrderID:     101,
						CreatedAt:   mustParseTime(t, "2022-01-01T12:00:00Z"),
						Funds:       transactions.Funds{BTC: "0.01", JPY: "-50000"},
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
						Funds:       transactions.Funds{BTC: "-0.01", JPY: "51000"},
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
			limit: 1, order: "asc", startingAfter: nil, endingBefore: nil,
			mockResponse: `{
				"success": true,
				"pagination": {
					"limit": 1,
					"order": "asc",
					"starting_after": null,
					"ending_before": null
				},
				"data": []
			}`,
			want: &GetResponse{
				Success: true,
				Pagination: PaginationOrder{
					Limit:         1,
					Order:         "asc",
					StartingAfter: nil,
					EndingBefore:  nil,
				},
				Data: []transactions.Transaction{},
			},
			wantErr: false,
		},
		"[異常系]JSONの形式が不正": {
			limit: 1, order: "asc", startingAfter: nil, endingBefore: nil,
			mockResponse: `{invalid json}`,
			want:         nil,
			wantErr:      true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			client := &mockClient{mockResponse: tt.mockResponse}
			tp := New(client)

			ctx := context.Background()
			got, err := tp.GET(ctx, tt.limit, tt.order, tt.startingAfter, tt.endingBefore)
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
			if got.Pagination.Limit != tt.want.Pagination.Limit {
				t.Errorf("Pagination.Limit: want=%v, got=%v", tt.want.Pagination.Limit, got.Pagination.Limit)
			}
			if got.Pagination.Order != tt.want.Pagination.Order {
				t.Errorf("Pagination.Order: want=%v, got=%v", tt.want.Pagination.Order, got.Pagination.Order)
			}
			if !equalIntPtr(got.Pagination.StartingAfter, tt.want.Pagination.StartingAfter) {
				t.Errorf("Pagination.StartingAfter: want=%v, got=%v", tt.want.Pagination.StartingAfter, got.Pagination.StartingAfter)
			}
			if !equalIntPtr(got.Pagination.EndingBefore, tt.want.Pagination.EndingBefore) {
				t.Errorf("Pagination.EndingBefore: want=%v, got=%v", tt.want.Pagination.EndingBefore, got.Pagination.EndingBefore)
			}
			if len(got.Data) != len(tt.want.Data) {
				t.Errorf("Data length: want=%d, got=%d", len(tt.want.Data), len(got.Data))
			}
			for i := range got.Data {
				wantTx := tt.want.Data[i]
				gotTx := got.Data[i]
				if gotTx.ID != wantTx.ID {
					t.Errorf("Data[%d].ID: want=%v, got=%v", i, wantTx.ID, gotTx.ID)
				}
				if gotTx.OrderID != wantTx.OrderID {
					t.Errorf("Data[%d].OrderID: want=%v, got=%v", i, wantTx.OrderID, gotTx.OrderID)
				}
				if !gotTx.CreatedAt.Equal(wantTx.CreatedAt) {
					t.Errorf("Data[%d].CreatedAt: want=%v, got=%v", i, wantTx.CreatedAt, gotTx.CreatedAt)
				}
				if !reflect.DeepEqual(gotTx.Funds, wantTx.Funds) {
					t.Errorf("Data[%d].Funds: want=%v, got=%v", i, wantTx.Funds, gotTx.Funds)
				}
				if gotTx.Pair != wantTx.Pair {
					t.Errorf("Data[%d].Pair: want=%v, got=%v", i, wantTx.Pair, gotTx.Pair)
				}
				if gotTx.Rate != wantTx.Rate {
					t.Errorf("Data[%d].Rate: want=%v, got=%v", i, wantTx.Rate, gotTx.Rate)
				}
				if gotTx.FeeCurrency != wantTx.FeeCurrency {
					t.Errorf("Data[%d].FeeCurrency: want=%v, got=%v", i, wantTx.FeeCurrency, gotTx.FeeCurrency)
				}
				if gotTx.Fee != wantTx.Fee {
					t.Errorf("Data[%d].Fee: want=%v, got=%v", i, wantTx.Fee, gotTx.Fee)
				}
				if gotTx.Liquidity != wantTx.Liquidity {
					t.Errorf("Data[%d].Liquidity: want=%v, got=%v", i, wantTx.Liquidity, gotTx.Liquidity)
				}
				if gotTx.Side != wantTx.Side {
					t.Errorf("Data[%d].Side: want=%v, got=%v", i, wantTx.Side, gotTx.Side)
				}
			}
		})
	}
}

func equalIntPtr(a, b *int) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}
