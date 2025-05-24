package orders

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
	req, _ := http.NewRequest(input.Method, "http://example.com", nil)
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

func TestOrdersGET(t *testing.T) {
	tests := map[string]struct {
		mockResponse string
		want         *GetResponse
		wantErr      bool
	}{
		"[正常系]": {
			mockResponse: `{
				"success": true,
				"id": 1,
				"pair": "btc_jpy",
				"status": "open",
				"order_type": "buy",
				"rate": "1234567",
				"stop_loss_rate": null,
				"maker_fee_rate": "0.001",
				"taker_fee_rate": "0.002",
				"amount": "0.05",
				"market_buy_amount": null,
				"executed_amount": 0.01,
				"executed_market_buy_amount": null,
				"expired_type": "",
				"prevented_match_id": 0,
				"expired_amount": "0",
				"expired_market_buy_amount": null,
				"time_in_force": "GTC",
				"created_at": "2025-01-01T12:34:56Z"
			}`,
			want: &GetResponse{
				Success:                 true,
				ID:                      1,
				Pair:                    "btc_jpy",
				Status:                  "open",
				OrderType:               "buy",
				Rate:                    strPtr("1234567"),
				StopLossRate:            nil,
				MakerFeeRate:            "0.001",
				TakerFeeRate:            "0.002",
				Amount:                  "0.05",
				MarketBuyAmount:         nil,
				ExecutedAmount:          0.01,
				ExecutedMarketBuyAmount: nil,
				ExpiredType:             "",
				PreventedMatchID:        0,
				ExpiredAmount:           "0",
				ExpiredMarketBuyAmount:  nil,
				TimeInForce:             "GTC",
				CreatedAt:               "2025-01-01T12:34:56Z",
			},
			wantErr: false,
		},
		"[異常系]JSONの形式が不正": {
			mockResponse: `{invalid json}`,
			want:         nil,
			wantErr:      true,
		},
		"[正常系]null多め": {
			mockResponse: `{
				"success": true,
				"id": 2,
				"pair": "btc_jpy",
				"status": "closed",
				"order_type": "sell",
				"rate": null,
				"stop_loss_rate": null,
				"maker_fee_rate": "0.001",
				"taker_fee_rate": "0.002",
				"amount": "0.00",
				"market_buy_amount": null,
				"executed_amount": 0,
				"executed_market_buy_amount": null,
				"expired_type": "",
				"prevented_match_id": 0,
				"expired_amount": "0",
				"expired_market_buy_amount": null,
				"time_in_force": "GTC",
				"created_at": "2025-01-02T12:34:56Z"
			}`,
			want: &GetResponse{
				Success:                 true,
				ID:                      2,
				Pair:                    "btc_jpy",
				Status:                  "closed",
				OrderType:               "sell",
				Rate:                    nil,
				StopLossRate:            nil,
				MakerFeeRate:            "0.001",
				TakerFeeRate:            "0.002",
				Amount:                  "0.00",
				MarketBuyAmount:         nil,
				ExecutedAmount:          0,
				ExecutedMarketBuyAmount: nil,
				ExpiredType:             "",
				PreventedMatchID:        0,
				ExpiredAmount:           "0",
				ExpiredMarketBuyAmount:  nil,
				TimeInForce:             "GTC",
				CreatedAt:               "2025-01-02T12:34:56Z",
			},
			wantErr: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			client := &mockClient{mockResponse: tt.mockResponse}
			orders := New(client)

			ctx := context.Background()
			got, err := orders.GET(ctx, 1)
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

func mustParseTime(t *testing.T, s string) time.Time {
	tm, err := time.Parse(time.RFC3339, s)
	if err != nil {
		t.Fatalf("invalid time: %v", err)
	}
	return tm
}

func TestOrdersPOST(t *testing.T) {
	tests := map[string]struct {
		pair         string
		orderType    string
		rate         float64
		amount       float64
		mockResponse string
		want         *PostResponse
		wantErr      bool
	}{
		"[正常系]": {
			pair:      "btc_jpy",
			orderType: "buy",
			rate:      1234567,
			amount:    0.01,
			mockResponse: `{
				"success": true,
				"id": 10,
				"rate": "1234567",
				"amount": "0.01",
				"order_type": "buy",
				"time_in_force": "GTC",
				"stop_loss_rate": null,
				"pair": "btc_jpy",
				"created_at": "2025-01-01T12:34:56Z"
			}`,
			want: &PostResponse{
				Success:      true,
				ID:           10,
				Rate:         "1234567",
				Amount:       "0.01",
				OrderType:    "buy",
				TimeInForce:  "GTC",
				StopLossRate: nil,
				Pair:         "btc_jpy",
				CreatedAt:    mustParseTime(t, "2025-01-01T12:34:56Z"),
			},
			wantErr: false,
		},
		"[異常系]JSONの形式が不正": {
			pair:         "btc_jpy",
			orderType:    "sell",
			rate:         1200000,
			amount:       0.1,
			mockResponse: `{invalid json}`,
			want:         nil,
			wantErr:      true,
		},
		"[正常系]stop_loss_rate有": {
			pair:      "btc_jpy",
			orderType: "sell",
			rate:      1000000,
			amount:    0.2,
			mockResponse: `{
				"success": true,
				"id": 11,
				"rate": "1000000",
				"amount": "0.2",
				"order_type": "sell",
				"time_in_force": "FOK",
				"stop_loss_rate": "900000",
				"pair": "btc_jpy",
				"created_at": "2025-01-03T10:00:00Z"
			}`,
			want: &PostResponse{
				Success:      true,
				ID:           11,
				Rate:         "1000000",
				Amount:       "0.2",
				OrderType:    "sell",
				TimeInForce:  "FOK",
				StopLossRate: strPtr("900000"),
				Pair:         "btc_jpy",
				CreatedAt:    mustParseTime(t, "2025-01-03T10:00:00Z"),
			},
			wantErr: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			client := &mockClient{mockResponse: tt.mockResponse}
			orders := New(client)

			ctx := context.Background()
			got, err := orders.POST(ctx, tt.pair, tt.orderType, tt.rate, tt.amount)
			if (err != nil) != tt.wantErr {
				t.Fatalf("error expectation: wantErr=%v, got=%v (err=%v)", tt.wantErr, err != nil, err)
			}
			if tt.wantErr {
				return
			}
			if got == nil {
				t.Fatalf("got == nil, want %+v", tt.want)
			}
			// compare except CreatedAt
			if got.Success != tt.want.Success {
				t.Errorf("Success: want=%v, got=%v", tt.want.Success, got.Success)
			}
			if got.ID != tt.want.ID {
				t.Errorf("ID: want=%v, got=%v", tt.want.ID, got.ID)
			}
			if got.Rate != tt.want.Rate {
				t.Errorf("Rate: want=%v, got=%v", tt.want.Rate, got.Rate)
			}
			if got.Amount != tt.want.Amount {
				t.Errorf("Amount: want=%v, got=%v", tt.want.Amount, got.Amount)
			}
			if got.OrderType != tt.want.OrderType {
				t.Errorf("OrderType: want=%v, got=%v", tt.want.OrderType, got.OrderType)
			}
			if got.TimeInForce != tt.want.TimeInForce {
				t.Errorf("TimeInForce: want=%v, got=%v", tt.want.TimeInForce, got.TimeInForce)
			}
			if !equalStrPtr(got.StopLossRate, tt.want.StopLossRate) {
				t.Errorf("StopLossRate: want=%v, got=%v", tt.want.StopLossRate, got.StopLossRate)
			}
			if got.Pair != tt.want.Pair {
				t.Errorf("Pair: want=%v, got=%v", tt.want.Pair, got.Pair)
			}
			if !got.CreatedAt.Equal(tt.want.CreatedAt) {
				t.Errorf("CreatedAt: want=%v, got=%v", tt.want.CreatedAt, got.CreatedAt)
			}
		})
	}
}

func TestOrdersDELETE(t *testing.T) {
	tests := map[string]struct {
		id           int
		mockResponse string
		want         *DeleteResponse
		wantErr      bool
	}{
		"[正常系]": {
			id: 42,
			mockResponse: `{
				"success": true,
				"id": 42
			}`,
			want: &DeleteResponse{
				Success: true,
				ID:      42,
			},
			wantErr: false,
		},
		"[異常系]JSONの形式が不正": {
			id:           99,
			mockResponse: `{invalid json}`,
			want:         nil,
			wantErr:      true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			client := &mockClient{mockResponse: tt.mockResponse}
			orders := New(client)

			ctx := context.Background()
			got, err := orders.DELETE(ctx, tt.id)
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

func equalStrPtr(a, b *string) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}
