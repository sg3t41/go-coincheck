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

	ctx := context.Background()

	msgs, err := coincheck.WebSocketTrade(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	// メッセージを受け取る
	go func() {
		for msg := range msgs {
			fmt.Printf("%s\n", msg)
		}
	}()

	time.Sleep(100000 * time.Second)
}
