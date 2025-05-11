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

	ctx := context.Background()

	accounts, err := coincheck.REST.Accounts(ctx)
	fmt.Println(accounts)

	balance, err := coincheck.REST.Balance(ctx)
	fmt.Println(balance)

	//	go func() {
	//		msgs, err := coincheck.WS.Trades(ctx, "btc_jpy")
	//		if err != nil {
	//			log.Fatalln("[ERROR] WebSocketTrade failed:", err)
	//		}
	//
	//		for msg := range msgs {
	//			fmt.Printf("[MESSAGE] %s\n", msg)
	//		}
	//	}()

	log.Println("[INFO] Press Ctrl+C to exit...")
	select {}
}
