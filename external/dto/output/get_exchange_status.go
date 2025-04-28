package output

import (
	"fmt"
	"strings"
	"time"
)

// GetExchangeStatus は取引所のステータスのレスポンス
type GetExchangeStatus struct {
	ExchangeStatus []ExchangeStatus `json:"exchange_status"`
}

// ExchangeStatus は取引所のステータス情報
type ExchangeStatus struct {
	Pair         string       `json:"pair"`
	Status       string       `json:"status"`
	Timestamp    int64        `json:"timestamp"`
	Availability Availability `json:"availability"`
}

// Availability は注文の可用性情報
type Availability struct {
	Order       bool `json:"order"`
	MarketOrder bool `json:"market_order"`
	Cancel      bool `json:"cancel"`
}

// Print は取引所のステータスを一覧表示
func (e *GetExchangeStatus) Print() {
	fmt.Println("\n=== 取引所のステータス ===")
	fmt.Println(strings.Repeat("=", 80))
	fmt.Printf("%-10s %-12s %-25s %-8s %-12s %-12s\n",
		"Pair", "Status", "Timestamp (JST)", "Order", "MarketOrder", "Cancel")
	fmt.Println(strings.Repeat("=", 80))

	for _, status := range e.ExchangeStatus {
		// JST に変換
		timestampJST := ItoJST(status.Timestamp)

		fmt.Printf("%-10s %-12s %-25s %-8v %-12v %-12v\n",
			status.Pair, status.Status, timestampJST, status.Availability.Order,
			status.Availability.MarketOrder, status.Availability.Cancel)
	}
	fmt.Println(strings.Repeat("=", 80))
}

// toJST はUTCのタイムスタンプをJSTに変換
func ItoJST(utcTimestamp int64) string {
	t := time.Unix(utcTimestamp, 0)
	jst := t.In(time.FixedZone("JST", 9*60*60)) // UTC+9
	return jst.Format("2006-01-02 15:04:05")
}
