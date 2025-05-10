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
	fmt.Println(key)
	fmt.Println(secret)
	coincheck, err := coincheck.New(key, secret)
	if err != nil {
		log.Fatalln(err)
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()

	_, err = coincheck.WebSocketTrade(ctx)
	if err != nil {
		log.Fatalln(err)
	}

}
