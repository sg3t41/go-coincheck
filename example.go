package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/sg3t41/go-coincheck/pkg/coincheck"
)

func main() {
	// 環境変数からAPIキー・シークレットを取得（または直書き可）
	key := os.Getenv("COINCHECK_API_KEY")
	secret := os.Getenv("COINCHECK_API_SECRET")

	client, err := coincheck.New(
		// APIキーとシークレットキーを指定
		coincheck.WithCredentials(key, secret),
		// REST API使用
		coincheck.WithREST(),
		// WebSocket API使用
		coincheck.WithWebSocket(),
	)
	if err != nil {
		log.Fatalf("[ERROR] coincheck.New: %v", err)
	}

	ctx := context.Background()

	/* --- REST API --- */

	// 1. ticker取得
	ticker, err := client.REST.Ticker(ctx, "btc_jpy")
	if err != nil {
		log.Printf("Ticker error: %v", err)
	} else {
		fmt.Printf("Ticker: %+v\n", ticker)
	}

	// 2. アカウント情報
	acct, err := client.REST.Accounts(ctx)
	if err != nil {
		log.Printf("Accounts error: %v", err)
	} else {
		fmt.Printf("Accounts: %+v\n", acct)
	}

	// 3. 残高
	bal, err := client.REST.Balance(ctx)
	if err != nil {
		log.Printf("Balance error: %v", err)
	} else {
		fmt.Printf("Balance: %+v\n", bal)
	}

	// 4. 取引所ステータス
	exst, err := client.REST.ExchangeStatus(ctx, "btc_jpy")
	if err != nil {
		log.Printf("ExchangeStatus error: %v", err)
	} else {
		fmt.Printf("ExchangeStatus: %+v\n", exst)
	}

	// 5. 参考レート
	ref, err := client.REST.ReferenceRate(ctx, "btc_jpy")
	if err != nil {
		log.Printf("ReferenceRate error: %v", err)
	} else {
		fmt.Printf("ReferenceRate: %+v\n", ref)
	}

	// 6. 注文レート
	ordersRate, err := client.REST.OrdersRate(ctx, "btc_jpy", "buy", 1000000, 0.01)
	if err != nil {
		log.Printf("OrdersRate error: %v", err)
	} else {
		fmt.Printf("OrdersRate: %+v\n", ordersRate)
	}

	// 7. 取引履歴
	trades, err := client.REST.Trades(ctx, "btc_jpy")
	if err != nil {
		log.Printf("Trades error: %v", err)
	} else {
		fmt.Printf("Trades: %+v\n", trades)
	}

	// 8. 板情報
	orderBooks, err := client.REST.OrderBooks(ctx, "btc_jpy")
	if err != nil {
		log.Printf("OrderBooks error: %v", err)
	} else {
		fmt.Printf("OrderBooks: %+v\n", orderBooks)
	}

	// 9. 取引履歴（簡易）
	transactions, err := client.REST.Transactions(ctx)
	if err != nil {
		log.Printf("Transactions error: %v", err)
	} else {
		fmt.Printf("Transactions: %+v\n", transactions)
	}

	// 10. 取引履歴（ページネーション）
	transactionsPag, err := client.REST.TransactionsPagination(ctx, 10, "desc", nil, nil)
	if err != nil {
		log.Printf("TransactionsPagination error: %v", err)
	} else {
		fmt.Printf("TransactionsPagination: %+v\n", transactionsPag)
	}

	// 11. 未約定注文一覧
	openOrders, err := client.REST.OpenOrders(ctx)
	if err != nil {
		log.Printf("OpenOrders error: %v", err)
	} else {
		fmt.Printf("OpenOrders: %+v\n", openOrders)
	}

	// 12. 指定IDの注文取得
	order, err := client.REST.GetOrder(ctx, 123456)
	if err != nil {
		log.Printf("GetOrder error: %v", err)
	} else {
		fmt.Printf("GetOrder: %+v\n", order)
	}

	// 13. 新規注文
	createOrder, err := client.REST.CreateOrder(ctx, "btc_jpy", "buy", 1000000, 0.01)
	if err != nil {
		log.Printf("CreateOrder error: %v", err)
	} else {
		fmt.Printf("CreateOrder: %+v\n", createOrder)
	}

	// 14. 注文キャンセル
	cancelOrder, err := client.REST.CancelOrder(ctx, 123456)
	if err != nil {
		log.Printf("CancelOrder error: %v", err)
	} else {
		fmt.Printf("CancelOrder: %+v\n", cancelOrder)
	}

	// 15. キャンセルステータス
	cancelStatus, err := client.REST.CancelStatus(ctx, 123456)
	if err != nil {
		log.Printf("CancelStatus error: %v", err)
	} else {
		fmt.Printf("CancelStatus: %+v\n", cancelStatus)
	}

	/* --- WebSocket API --- */

	// Ctrl+Cで終了できるように
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	wsCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func() {
		<-sigCh
		fmt.Println("\n[INFO] interrupt received, closing WebSocket...")
		cancel()
	}()

	// 板(OrderBook)の購読
	obCh, err := client.WS.Orderbook(wsCtx, "btc_jpy")
	if err != nil {
		log.Printf("WS.Orderbook error: %v", err)
	} else {
		go func() {
			for msg := range obCh {
				fmt.Printf("[WS Orderbook] message: %s\n", msg)
			}
			fmt.Println("[WS Orderbook] channel closed")
		}()
	}

	// 成行約定(Trades)の購読
	tradesCh, err := client.WS.Trades(wsCtx, "btc_jpy")
	if err != nil {
		log.Printf("WS.Trades error: %v", err)
	} else {
		go func() {
			for msg := range tradesCh {
				fmt.Printf("[WS Trades] message: %s\n", msg)
			}
			fmt.Println("[WS Trades] channel closed")
		}()
	}

	// サンプルなので10秒待って終了（Ctrl+Cでも止まる）
	fmt.Println("[INFO] WebSocketサンプルは10秒で自動終了します。")
	time.Sleep(10 * time.Second)
}
