package main

import (
	"context"
	"log"
	"os"

	"github.com/sg3t41/go-coincheck/external/coincheck"
)

var (
	key    string
	secret string
)

func init() {
	key = os.Getenv("COINCHECK_ACCESS_KEY")
	secret = os.Getenv("COINCHECK_SECRET_ACCESS_KEY")
}

func main() {
	client, err := coincheck.New(
		// 認証情報を使用
		coincheck.UseCredentials(key, secret),
		// HTTP REST APIを使用
		coincheck.UseHTTP(),
		// WebSocket APIを使用
		coincheck.UseWebSocket(),
	)
	if err != nil {
		log.Fatalf("[ERROR] %v", err)
	}

	ctx := context.Background()
	ticker, err := client.REST.Ticker(ctx, "btc_jpy")
	if err != nil {
		log.Printf("[ERROR] %v", err)
		return
	}

	log.Printf("%+v\n", ticker)
}
