package main

import (
	"context"
	"fmt"
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
	coincheck, err := coincheck.New(key, secret)
	if err != nil {
		log.Fatalln("[ERROR] Failed to initialize Coincheck client:", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tradeChan := make(chan string, 100)

	go func() {
		err := coincheck.WS.Trades(ctx, "btc_jpy-trades", tradeChan)
		if err != nil {
			log.Fatalln("[ERROR] WebSocketTrade failed:", err)
		}
	}()

	go func() {
		for msg := range tradeChan {
			fmt.Printf("[MESSAGE] %s\n", msg)
		}
	}()

	log.Println("[INFO] Press Ctrl+C to exit...")
	select {}
}
