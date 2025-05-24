package ordersrate

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

func TestOrdersRateGET(t *testing.T) {
	tests := map[string]struct {
		pair         string
		orderType    string
		price        float64
		amount       float64
		mockResponse string
		want         *GetResponse
		wantErr      bool
	}{
		"[正常系]sell": {
			pair:      "btc_jpy",
			orderType: "sell",
			price:     1234567.89,
			amount:    0, // unused for sell
			mockResponse: `{
				"success": true,
				"rate": "1234567.89",
				"price": "1234567.89",
				"amount": "0.5"
			}`,
			want: &GetResponse{
				Success: true,
				Rate:    "1234567.89",
				Price:   "1234567.89",
				Amount:  "0.5",
			},
			wantErr: false,
		},
		"[正常系]buy": {
			pair:      "btc_jpy",
			orderType: "buy",
			price:     0, // unused for buy
			amount:    0.1,
			mockResponse: `{
				"success": true,
				"rate": "1234000",
				"price": "123400",
				"amount": "0.1"
			}`,
			want: &GetResponse{
				Success: true,
				Rate:    "1234000",
				Price:   "123400",
				Amount:  "0.1",
			},
			wantErr: false,
		},
		"[異常系]JSONの形式が不正": {
			pair:         "btc_jpy",
			orderType:    "buy",
			price:        0,
			amount:       0.1,
			mockResponse: `{invalid json}`,
			want:         nil,
			wantErr:      true,
		},
		"[正常系]値が空": {
			pair:      "btc_jpy",
			orderType: "sell",
			price:     0,
			amount:    0,
			mockResponse: `{
				"success": true,
				"rate": "",
				"price": "",
				"amount": ""
			}`,
			want: &GetResponse{
				Success: true,
				Rate:    "",
				Price:   "",
				Amount:  "",
			},
			wantErr: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			client := &mockClient{mockResponse: tt.mockResponse}
			ordersRate := New(client)

			ctx := context.Background()
			got, err := ordersRate.GET(ctx, tt.pair, tt.orderType, tt.price, tt.amount)
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
			if got.Rate != tt.want.Rate {
				t.Errorf("Rate: want=%v, got=%v", tt.want.Rate, got.Rate)
			}
			if got.Price != tt.want.Price {
				t.Errorf("Price: want=%v, got=%v", tt.want.Price, got.Price)
			}
			if got.Amount != tt.want.Amount {
				t.Errorf("Amount: want=%v, got=%v", tt.want.Amount, got.Amount)
			}
		})
	}
}
