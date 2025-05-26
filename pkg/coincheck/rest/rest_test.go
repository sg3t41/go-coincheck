package rest

import (
	"testing"
)

// client.Newの動作を模倣するためのダミークライアントを用意
type dummyClient struct{}

func TestNew_Normal(t *testing.T) {
	// 正しいkey/secretで生成できるか
	r, err := New("dummy_key", "dummy_secret")
	if err != nil {
		t.Fatalf("REST.New 正常系でエラー: %v", err)
	}
	if r == nil {
		t.Fatal("REST.New: 戻り値がnil")
	}
}

func TestNew_ClientError(t *testing.T) {
	// client.Newが失敗した場合にちゃんとエラーを返すかのテスト
	// ※ client.Newが引数で必ず失敗するケースを作れるならそれで
	// もし難しければ、ここはスキップしてもOK
	_, err := New("", "")
	if err == nil {
		t.Error("client.Newが失敗した場合にerrが返らない")
	}
}
