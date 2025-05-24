package orderbooks

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

func TestOrderBooksGET(t *testing.T) {
	tests := map[string]struct {
		mockResponse string
		want         *GetResponse
		wantErr      bool
	}{
		"[正常系]": {
			mockResponse: `{
				"asks": [["1234567", "0.1"], ["1234570", "0.2"]],
				"bids": [["1234500", "0.5"], ["1234450", "1.0"]]
			}`,
			want: &GetResponse{
				Asks: [][]string{
					{"1234567", "0.1"},
					{"1234570", "0.2"},
				},
				Bids: [][]string{
					{"1234500", "0.5"},
					{"1234450", "1.0"},
				},
			},
			wantErr: false,
		},
		"[異常系]JSONの形式が不正": {
			mockResponse: `{invalid json}`,
			want:         nil,
			wantErr:      true,
		},
		"[正常系]空配列": {
			mockResponse: `{
				"asks": [],
				"bids": []
			}`,
			want: &GetResponse{
				Asks: [][]string{},
				Bids: [][]string{},
			},
			wantErr: false,
		},
		"[正常系]片方のみデータ有": {
			mockResponse: `{
				"asks": [["100", "0.1"]],
				"bids": []
			}`,
			want: &GetResponse{
				Asks: [][]string{
					{"100", "0.1"},
				},
				Bids: [][]string{},
			},
			wantErr: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			client := &mockClient{mockResponse: tt.mockResponse}
			orderBooks := New(client)

			ctx := context.Background()
			got, err := orderBooks.GET(ctx, "btc_jpy")
			if (err != nil) != tt.wantErr {
				t.Fatalf("error expectation: wantErr=%v, got=%v (err=%v)", tt.wantErr, err != nil, err)
			}
			if tt.wantErr {
				return
			}
			if got == nil {
				t.Fatalf("got == nil, want %+v", tt.want)
			}
			if !reflect.DeepEqual(got.Asks, tt.want.Asks) {
				t.Errorf("Asks: want=%v, got=%v", tt.want.Asks, got.Asks)
			}
			if !reflect.DeepEqual(got.Bids, tt.want.Bids) {
				t.Errorf("Bids: want=%v, got=%v", tt.want.Bids, got.Bids)
			}
		})
	}
}
