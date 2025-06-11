package balance

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

func TestBalanceGET(t *testing.T) {
	tests := map[string]struct {
		mockResponse string
		want         *GetResponse
		wantErr      bool
	}{
		"[正常系]全項目": {
			mockResponse: `{
				"success": true,
				"jpy": "123456.78",
				"btc": "0.12345678",
				"jpy_reserved": "1000",
				"btc_reserved": "0.01",
				"jpy_lending": "2000",
				"btc_lending": "0.02",
				"jpy_lend_in_use": "3000",
				"btc_lend_in_use": "0.03",
				"jpy_lent": "4000",
				"btc_lent": "0.04",
				"jpy_debt": "5000",
				"btc_debt": "0.05",
				"jpy_tsumitate": "6000",
				"btc_tsumitate": "0.06"
			}`,
			want: &GetResponse{
				Success:      true,
				JPY:          "123456.78",
				BTC:          "0.12345678",
				JPYReserved:  "1000",
				BTCReserved:  "0.01",
				JPYLending:   "2000",
				BTCLending:   "0.02",
				JPYLendInUse: "3000",
				BTCLendInUse: "0.03",
				JPYLent:      "4000",
				BTCLent:      "0.04",
				JPYDebt:      "5000",
				BTCDebt:      "0.05",
				JPYTsumitate: "6000",
				BTCTsumitate: "0.06",
			},
			wantErr: false,
		},
		"[正常系]最低限": {
			mockResponse: `{
				"success": true,
				"jpy": "1",
				"btc": "0",
				"jpy_reserved": "0",
				"btc_reserved": "0",
				"jpy_lending": "0",
				"btc_lending": "0",
				"jpy_lend_in_use": "0",
				"btc_lend_in_use": "0",
				"jpy_lent": "0",
				"btc_lent": "0",
				"jpy_debt": "0",
				"btc_debt": "0",
				"jpy_tsumitate": "0",
				"btc_tsumitate": "0"
			}`,
			want: &GetResponse{
				Success:      true,
				JPY:          "1",
				BTC:          "0",
				JPYReserved:  "0",
				BTCReserved:  "0",
				JPYLending:   "0",
				BTCLending:   "0",
				JPYLendInUse: "0",
				BTCLendInUse: "0",
				JPYLent:      "0",
				BTCLent:      "0",
				JPYDebt:      "0",
				BTCDebt:      "0",
				JPYTsumitate: "0",
				BTCTsumitate: "0",
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
			balance := New(client)

			ctx := context.Background()
			got, err := balance.GET(ctx)
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
