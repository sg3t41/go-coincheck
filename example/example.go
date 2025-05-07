package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/sg3t41/go-coincheck/external/coincheck"
)

const (
	key    = ""
	secret = ""
)

func main() {
	coincheck, err := coincheck.New(key, secret)
	if err != nil {
		log.Fatalln(err)
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()

	ac, err := coincheck.Balance(ctx)
	if err != nil {
		fmt.Println(err)
	}
	ac.Print()
}
