package ticker

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

func TestTickerGET(t *testing.T) {
	tests := map[string]struct {
		mockResponse string
		want         *GetResponse
		wantErr      bool
	}{
		"[正常系]": {
			mockResponse: `{
				"last": 1234567,
				"bid": 1234560,
				"ask": 1234570,
				"high": 1300000,
				"low": 1200000,
				"volume": 1.2345,
				"timestamp": 1620000000
			}`,
			want: &GetResponse{
				Last:      1234567,
				Bid:       1234560,
				Ask:       1234570,
				High:      1300000,
				Low:       1200000,
				Volume:    1.2345,
				Timestamp: 1620000000,
			},
			wantErr: false,
		},
		"[異常系]JSONの形式が不正": {
			mockResponse: `{invalid json}`,
			want:         nil,
			wantErr:      true,
		},
		"[正常系]0を含んだフィールド有": {
			mockResponse: `{
				"last": 111,
				"bid": 222,
				"ask": 333,
				"high": 0,
				"low": 0,
				"volume": 0,
				"timestamp": 0
			}`,
			want: &GetResponse{
				Last:      111,
				Bid:       222,
				Ask:       333,
				High:      0,
				Low:       0,
				Volume:    0,
				Timestamp: 0,
			},
			wantErr: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			client := &mockClient{mockResponse: tt.mockResponse}
			ticker := New(client)

			ctx := context.Background()
			got, err := ticker.GET(ctx, "btc_jpy")
			if (err != nil) != tt.wantErr {
				t.Fatalf("error expectation: wantErr=%v, got=%v (err=%v)", tt.wantErr, err != nil, err)
			}
			if tt.wantErr {
				return
			}
			if got == nil {
				t.Fatalf("got == nil, want %+v", tt.want)
			}
			if got.Last != tt.want.Last {
				t.Errorf("Last: want=%v, got=%v", tt.want.Last, got.Last)
			}
			if got.Bid != tt.want.Bid {
				t.Errorf("Bid: want=%v, got=%v", tt.want.Bid, got.Bid)
			}
			if got.Ask != tt.want.Ask {
				t.Errorf("Ask: want=%v, got=%v", tt.want.Ask, got.Ask)
			}
			if got.High != tt.want.High {
				t.Errorf("High: want=%v, got=%v", tt.want.High, got.High)
			}
			if got.Low != tt.want.Low {
				t.Errorf("Low: want=%v, got=%v", tt.want.Low, got.Low)
			}
			if got.Volume != tt.want.Volume {
				t.Errorf("Volume: want=%v, got=%v", tt.want.Volume, got.Volume)
			}
			if got.Timestamp != tt.want.Timestamp {
				t.Errorf("Timestamp: want=%v, got=%v", tt.want.Timestamp, got.Timestamp)
			}
		})
	}
}
