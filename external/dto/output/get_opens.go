package output

import (
	"fmt"
	"strings"
	"time"
)

// Opens は未決済注文のレスポンス
type Opens struct {
	Success bool    `json:"success"`
	Orders  []Order `json:"orders"`
}

// Order は未決済注文の情報
type Order struct {
	ID                     int     `json:"id"`
	OrderType              string  `json:"order_type"`
	Rate                   *string `json:"rate"` // null の場合があるためポインタ型
	Pair                   string  `json:"pair"`
	PendingAmount          *string `json:"pending_amount"`
	PendingMarketBuyAmount *string `json:"pending_market_buy_amount"`
	StopLossRate           *string `json:"stop_loss_rate"`
	CreatedAt              string  `json:"created_at"`
}

// Print は未決済注文を一覧表示
func (o *Opens) Print() {
	fmt.Println("\n=== 未決済注文一覧 ===")
	fmt.Println(strings.Repeat("-", 80))
	fmt.Printf("%-10s %-6s %-10s %-8s %-12s %-12s %-12s %-25s\n",
		"ID", "Type", "Rate", "Pair", "Pending", "MarketBuy", "StopLoss", "CreatedAt (JST)")
	fmt.Println(strings.Repeat("-", 80))

	for _, order := range o.Orders {
		// Null の場合は "N/A" にする
		rate := getOrNA(order.Rate)
		pending := getOrNA(order.PendingAmount)
		marketBuy := getOrNA(order.PendingMarketBuyAmount)
		stopLoss := getOrNA(order.StopLossRate)

		// JST に変換
		createdAtJST := StrToJST(order.CreatedAt)

		fmt.Printf("%-10d %-6s %-10s %-8s %-12s %-12s %-12s %-25s\n",
			order.ID, order.OrderType, rate, order.Pair, pending, marketBuy, stopLoss, createdAtJST)
	}
	fmt.Println(strings.Repeat("-", 80))
}

// getOrNA は nil の場合に "N/A" を返す
func getOrNA(value *string) string {
	if value == nil {
		return "N/A"
	}
	return *value
}

// toJST はUTCの時間をJSTに変換
func StrToJST(utcTime string) string {
	t, err := time.Parse(time.RFC3339, utcTime)
	if err != nil {
		return "Invalid Time"
	}
	jst := t.In(time.FixedZone("JST", 9*60*60)) // UTC+9
	return jst.Format("2006-01-02 15:04:05")
}
