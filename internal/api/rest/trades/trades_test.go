package trades

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

func intPtr(v int) *int { return &v }

func TestTradesGET(t *testing.T) {
	tests := map[string]struct {
		mockResponse string
		want         *GetResponse
		wantErr      bool
	}{
		"[正常系]": {
			mockResponse: `{
				"success": true,
				"pagination": {
					"limit": 3,
					"order": "desc",
					"starting_after": 100,
					"ending_before": 200
				},
				"data": [
					{
						"id": 1,
						"amount": "0.01",
						"rate": "1234567",
						"pair": "btc_jpy",
						"order_type": "buy",
						"created_at": "2021-01-01T10:00:00Z"
					},
					{
						"id": 2,
						"amount": "0.02",
						"rate": "1234570",
						"pair": "btc_jpy",
						"order_type": "sell",
						"created_at": "2021-01-01T11:00:00Z"
					}
				]
			}`,
			want: &GetResponse{
				Success: true,
				Pagination: Pagination{
					Limit:         3,
					Order:         "desc",
					StartingAfter: intPtr(100),
					EndingBefore:  intPtr(200),
				},
				Data: []Trade{
					{
						ID:        1,
						Amount:    "0.01",
						Rate:      "1234567",
						Pair:      "btc_jpy",
						OrderType: "buy",
						CreatedAt: "2021-01-01T10:00:00Z",
					},
					{
						ID:        2,
						Amount:    "0.02",
						Rate:      "1234570",
						Pair:      "btc_jpy",
						OrderType: "sell",
						CreatedAt: "2021-01-01T11:00:00Z",
					},
				},
			},
			wantErr: false,
		},
		"[異常系]JSONの形式が不正": {
			mockResponse: `{invalid json}`,
			want:         nil,
			wantErr:      true,
		},
		"[正常系]空データ": {
			mockResponse: `{
				"success": true,
				"pagination": {
					"limit": 0,
					"order": "asc",
					"starting_after": null,
					"ending_before": null
				},
				"data": []
			}`,
			want: &GetResponse{
				Success: true,
				Pagination: Pagination{
					Limit:         0,
					Order:         "asc",
					StartingAfter: nil,
					EndingBefore:  nil,
				},
				Data: []Trade{},
			},
			wantErr: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			client := &mockClient{mockResponse: tt.mockResponse}
			trades := New(client)

			ctx := context.Background()
			got, err := trades.GET(ctx, "btc_jpy")
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
			if (got.Pagination.StartingAfter == nil) != (tt.want.Pagination.StartingAfter == nil) ||
				(got.Pagination.StartingAfter != nil && tt.want.Pagination.StartingAfter != nil && *got.Pagination.StartingAfter != *tt.want.Pagination.StartingAfter) {
				t.Errorf("Pagination.StartingAfter: want=%v, got=%v", tt.want.Pagination.StartingAfter, got.Pagination.StartingAfter)
			}
			if (got.Pagination.EndingBefore == nil) != (tt.want.Pagination.EndingBefore == nil) ||
				(got.Pagination.EndingBefore != nil && tt.want.Pagination.EndingBefore != nil && *got.Pagination.EndingBefore != *tt.want.Pagination.EndingBefore) {
				t.Errorf("Pagination.EndingBefore: want=%v, got=%v", tt.want.Pagination.EndingBefore, got.Pagination.EndingBefore)
			}
			if len(got.Data) != len(tt.want.Data) {
				t.Errorf("Data length: want=%d, got=%d", len(tt.want.Data), len(got.Data))
			} else {
				for i := range got.Data {
					if got.Data[i] != tt.want.Data[i] {
						t.Errorf("Data[%d]: want=%+v, got=%+v", i, tt.want.Data[i], got.Data[i])
					}
				}
			}
		})
	}
}
