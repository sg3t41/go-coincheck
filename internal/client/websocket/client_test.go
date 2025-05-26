package websocket

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

// よもや統合テスト

func TestWebSocketClient_統合テスト(t *testing.T) {
	tests := map[string]struct {
		serverAction      func(conn *websocket.Conn, t *testing.T)
		subscribeChannel  string
		wantMessages      []string
		expectSubscribeOK bool
	}{
		// 正常系: サーバーが3メッセージ送信→クライアントが3つ受信できる
		"正常系: 3メッセージを受信できる": {
			serverAction: func(conn *websocket.Conn, t *testing.T) {
				// サブスクライブメッセージ受信確認
				_, submsg, err := conn.ReadMessage()
				if err != nil {
					t.Fatalf("サブスクライブメッセージ受信失敗: %v", err)
				}
				expected := `{"type":"subscribe","channel":"test_channel"}`
				if string(submsg) != expected {
					t.Errorf("subscribe msg: 期待値 %v, 実際 %v", expected, string(submsg))
				}
				// クライアントへ3メッセージ送信
				for i := 0; i < 3; i++ {
					conn.WriteMessage(websocket.TextMessage, []byte(`{"data":`+strconv.Itoa(i)+`}`))
				}

				time.Sleep(100 * time.Millisecond)
				conn.Close()
			},
			subscribeChannel:  "test_channel",
			wantMessages:      []string{`{"data":0}`, `{"data":1}`, `{"data":2}`},
			expectSubscribeOK: true,
		},
		// 異常系: サーバーが即切断→クライアントは受信できない
		"異常系: サーバーが即切断": {
			serverAction: func(conn *websocket.Conn, t *testing.T) {
				conn.Close()
			},
			subscribeChannel:  "test_channel",
			wantMessages:      nil,
			expectSubscribeOK: false,
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			// テスト用WebSocketサーバ起動
			upgrader := websocket.Upgrader{}
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				conn, err := upgrader.Upgrade(w, r, nil)
				if err != nil {
					t.Fatalf("アップグレード失敗: %v", err)
					return
				}
				defer conn.Close()
				tt.serverAction(conn, t)
			}))
			defer server.Close()

			// クライアント作成
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			defer cancel()
			wsurl := "ws" + server.URL[len("http"):] // ws://... に変換

			client, err := NewClient(WithBaseURL(wsurl))
			if err != nil {
				t.Fatalf("NewClient エラー: %v", err)
			}
			in := make(chan string, 10)
			subErr := client.Subscribe(ctx, tt.subscribeChannel, in)
			got := []string{}
			for m := range in {
				got = append(got, m)
			}
			if tt.expectSubscribeOK && subErr != nil {
				t.Errorf("Subscribe: 想定外のエラー: %v", subErr)
			}
			// 異常系: サーバ切断時はSubscribe自体はエラーにならずメッセージも来ないことがある
			if !tt.expectSubscribeOK && subErr == nil && len(got) == 0 {
				t.Log("Subscribe: サーバ切断のためメッセージなし")
			}
			if len(tt.wantMessages) > 0 && len(got) != len(tt.wantMessages) {
				t.Errorf("受信メッセージ数: 期待 %d, 実際 %d: %v", len(tt.wantMessages), len(got), got)
			}
			for i := range tt.wantMessages {
				if i >= len(got) {
					t.Errorf("メッセージ不足: %d番目 期待 %q", i, tt.wantMessages[i])
					continue
				}
				if got[i] != tt.wantMessages[i] {
					t.Errorf("msg[%d]: 期待 %q, 実際 %q", i, tt.wantMessages[i], got[i])
				}
			}
			client.Close()
		})
	}
}
