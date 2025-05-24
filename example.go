package main

import (
	"context"
	"log"
	"os"

	"github.com/sg3t41/go-coincheck/pkg/coincheck"
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
		coincheck.WithCredentials(key, secret),
		coincheck.WithREST(),
		coincheck.WithWebSocket(),
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
