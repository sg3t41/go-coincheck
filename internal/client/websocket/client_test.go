package websocket

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

func TestWebSocketClient_Connect_Subscribe_ReadMessages(t *testing.T) {
	// テスト用WSサーバ
	upgrader := websocket.Upgrader{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			t.Fatalf("upgrade error: %v", err)
			return
		}
		defer conn.Close()
		// サブスクライブメッセージ受信を待つ
		_, submsg, err := conn.ReadMessage()
		if err != nil {
			t.Fatalf("failed to read subscribe: %v", err)
		}
		expected := `{"type":"subscribe","channel":"test_channel"}`
		if string(submsg) != expected {
			t.Errorf("subscribe msg want %v, got %v", expected, string(submsg))
		}
		// いくつか送信
		for i := 0; i < 3; i++ {
			conn.WriteMessage(websocket.TextMessage, []byte(`{"data":`+string('0'+i)+`}`))
		}
		time.Sleep(100 * time.Millisecond)
		conn.Close() // クライアントの再接続処理も見たいなら
	}))
	defer server.Close()

	// テスト対象
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	wsurl := "ws" + server.URL[len("http"):] // ws://...
	client, err := NewClient(WithBaseURL(wsurl))
	if err != nil {
		t.Fatalf("NewClient error: %v", err)
	}
	in := make(chan string, 10)
	if err := client.Subscribe(ctx, "test_channel", in); err != nil {
		t.Fatalf("Subscribe error: %v", err)
	}
	got := []string{}
	for m := range in {
		got = append(got, m)
	}
	if len(got) != 3 {
		t.Errorf("want 3 messages, got %d", len(got))
	}
	client.Close()
}
