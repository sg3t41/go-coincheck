package trades

import (
	"context"
	"errors"
	"net/http"
	"reflect"
	"testing"
	"time"

	httpc "github.com/sg3t41/go-coincheck/internal/client/http"
)

/* MOCK ここから */

type mockClient struct {
	connectErr   error
	subscribeErr error
	subscribedCh string
}

func (m *mockClient) Connect(ctx context.Context) error {
	return m.connectErr
}

func (m *mockClient) Subscribe(ctx context.Context, channel string, in chan<- string) error {
	m.subscribedCh = channel
	if m.subscribeErr != nil {
		return m.subscribeErr
	}
	// 疑似的にメッセージを送信しチャネルを閉じる
	go func() {
		in <- "msg1"
		in <- "msg2"
		close(in)
	}()
	return nil
}

func (m *mockClient) Close() error { return nil }

func (m *mockClient) CreateRequest(ctx context.Context, input httpc.RequestInput) (*http.Request, error) {
	req, _ := http.NewRequest("GET", "http://example.com", nil)
	return req, nil
}
func (m *mockClient) Do(req *http.Request, output any) error { return nil }

/* MOCK ここまで */

func TestTradesSubscribe(t *testing.T) {
	tests := map[string]struct {
		connectErr   error
		subscribeErr error
		wantMsgs     []string
		wantErr      bool
		wantChannel  string
	}{
		"[正常系]": {
			connectErr:   nil,
			subscribeErr: nil,
			wantMsgs:     []string{"msg1", "msg2"},
			wantErr:      false,
			wantChannel:  "btc_jpy-trades",
		},
		"[異常系]Connect失敗": {
			connectErr:   errors.New("connect error"),
			subscribeErr: nil,
			wantMsgs:     nil,
			wantErr:      true,
			wantChannel:  "",
		},
		"[異常系]Subscribe失敗": {
			connectErr:   nil,
			subscribeErr: errors.New("subscribe error"),
			wantMsgs:     nil,
			wantErr:      true,
			wantChannel:  "btc_jpy-trades",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			mock := &mockClient{
				connectErr:   tt.connectErr,
				subscribeErr: tt.subscribeErr,
			}
			tr := New(mock)

			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			ch, err := tr.Subscribe(ctx, "btc_jpy")
			if (err != nil) != tt.wantErr {
				t.Fatalf("error: wantErr=%v, got=%v, err=%v", tt.wantErr, err != nil, err)
			}
			if tt.wantErr {
				return
			}
			// メッセージを受信して検証
			got := []string{}
			for msg := range ch {
				got = append(got, msg)
			}
			if !reflect.DeepEqual(got, tt.wantMsgs) {
				t.Errorf("got messages = %v, want %v", got, tt.wantMsgs)
			}
			// Subscribeで渡されたチャネル名の検証
			if mock.subscribedCh != tt.wantChannel {
				t.Errorf("Subscribe channel: got=%v, want=%v", mock.subscribedCh, tt.wantChannel)
			}
		})
	}
}
