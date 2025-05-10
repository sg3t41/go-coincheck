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
		log.Fatalln(err)
	}

	ctx := context.Background()

	tradeChan := make(chan string)

	if err := coincheck.WebSocketTrade(ctx, "doge_jpy-orderbook", tradeChan); err != nil {
		log.Fatalln(err)
	}

	// メッセージを受け取る
	go func() {
		for msg := range tradeChan {
			fmt.Printf("%s\n", msg)
		}
	}()

}
