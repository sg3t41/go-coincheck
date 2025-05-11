package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

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

	ctx, cancel := context.WithTimeout(context.Background(), 10000*time.Second)
	defer cancel()

	tradeChan := make(chan string, 100)

	err = coincheck.WebSocketTrade(ctx, "btc_jpy-trades", tradeChan)
	if err != nil {
		log.Fatalln(err)
	}

	// メッセージを受け取る
	go func() {
		for msg := range tradeChan {
			fmt.Printf("%s\n", msg)
		}
	}()

	time.Sleep(10000 * time.Second)
}
