package exchangestatus

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

func TestExchangeStatusGET(t *testing.T) {
	tests := map[string]struct {
		mockResponse string
		want         *GetReponse
		wantErr      bool
	}{
		"[正常系]": {
			mockResponse: `{
				"pair": "btc_jpy",
				"status": "OK",
				"timestamp": 1620000000,
				"availability": {
					"order": true,
					"market_order": false,
					"cancel": true
				}
			}`,
			want: &GetReponse{
				Pair:      "btc_jpy",
				Status:    "OK",
				Timestamp: 1620000000,
				Availability: Availability{
					Order:       true,
					MarketOrder: false,
					Cancel:      true,
				},
			},
			wantErr: false,
		},
		"[異常系]JSONの形式が不正": {
			mockResponse: `{invalid json}`,
			want:         nil,
			wantErr:      true,
		},
		"[正常系]全てfalse": {
			mockResponse: `{
				"pair": "eth_jpy",
				"status": "NO",
				"timestamp": 1000000000,
				"availability": {
					"order": false,
					"market_order": false,
					"cancel": false
				}
			}`,
			want: &GetReponse{
				Pair:      "eth_jpy",
				Status:    "NO",
				Timestamp: 1000000000,
				Availability: Availability{
					Order:       false,
					MarketOrder: false,
					Cancel:      false,
				},
			},
			wantErr: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			client := &mockClient{mockResponse: tt.mockResponse}
			es := New(client)

			ctx := context.Background()
			got, err := es.GET(ctx, "btc_jpy")
			if (err != nil) != tt.wantErr {
				t.Fatalf("error expectation: wantErr=%v, got=%v (err=%v)", tt.wantErr, err != nil, err)
			}
			if tt.wantErr {
				return
			}
			if got == nil {
				t.Fatalf("got == nil, want %+v", tt.want)
			}
			if got.Pair != tt.want.Pair {
				t.Errorf("Pair: want=%v, got=%v", tt.want.Pair, got.Pair)
			}
			if got.Status != tt.want.Status {
				t.Errorf("Status: want=%v, got=%v", tt.want.Status, got.Status)
			}
			if got.Timestamp != tt.want.Timestamp {
				t.Errorf("Timestamp: want=%v, got=%v", tt.want.Timestamp, got.Timestamp)
			}
			if got.Availability.Order != tt.want.Availability.Order {
				t.Errorf("Availability.Order: want=%v, got=%v", tt.want.Availability.Order, got.Availability.Order)
			}
			if got.Availability.MarketOrder != tt.want.Availability.MarketOrder {
				t.Errorf("Availability.MarketOrder: want=%v, got=%v", tt.want.Availability.MarketOrder, got.Availability.MarketOrder)
			}
			if got.Availability.Cancel != tt.want.Availability.Cancel {
				t.Errorf("Availability.Cancel: want=%v, got=%v", tt.want.Availability.Cancel, got.Availability.Cancel)
			}
		})
	}
}
