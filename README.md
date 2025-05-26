# go-coincheck

[![GitHub release](https://img.shields.io/github/v/release/sg3t41/go-coincheck?include_prereleases)](https://github.com/sg3t41/go-coincheck/releases)
![Go version](https://img.shields.io/github/go-mod/go-version/sg3t41/go-coincheck?style=flat-square)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/db340dce37434e5bbef6b2261eb8fb8d)](https://app.codacy.com/gh/sg3t41/go-coincheck/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade)
[![Go Report Card](https://goreportcard.com/badge/github.com/sg3t41/go-coincheck)](https://goreportcard.com/report/github.com/sg3t41/go-coincheck)
[![Go Reference](https://pkg.go.dev/badge/github.com/sg3t41/go-coincheck/v2.svg)](https://pkg.go.dev/github.com/sg3t41/go-coincheck)


[Coincheck](https://coincheck.com/) REST & WebSocket API クライアントライブラリのGo実装です。

## Install

```sh
go get github.com/sg3t41/go-coincheck
```

## Usage
[`example.go`](https://github.com/sg3t41/go-coincheck/blob/main/example.go)
```go
package main

import (
	"context"
	"fmt"

	"github.com/sg3t41/go-coincheck/pkg/coincheck"
)

func main() {
	api, _ := coincheck.New(
		coincheck.WithCredentials("key", "secret"),
		coincheck.WithREST(),
		coincheck.WithWebSocket(),
	)

	ctx := context.Background()

	/* --- REST API --- */
	ticker, _ := api.REST.Ticker(ctx, "btc_jpy")
	fmt.Printf("Ticker: %+v\n", ticker)

	/* --- WebSocket API --- */
	obch, _ := api.WS.Orderbook(ctx, "btc_jpy")
	go func() {
		for msg := range obch {
			fmt.Printf("%s\n", msg)
		}
	}()
}
```
