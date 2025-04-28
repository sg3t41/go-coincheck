package output

import (
	"fmt"
	"time"
)

// CreateOrder は CreateOrder のレスポンス
type CreateOrder struct {
	// Success はリクエストの成功を示す
	Success bool `json:"success"`
	// ID は注文ID
	ID int `json:"id"`
	// Rate は注文レート
	Rate string `json:"rate"`
	// Amount は注文量
	Amount string `json:"amount"`
	// OrderType は注文タイプ
	OrderType string `json:"order_type"`
	// TimeInForce は注文の有効期間
	TimeInForce string `json:"time_in_force"`
	// StopLossRate はストップロスレート
	StopLossRate *string `json:"stop_loss_rate"`
	// Pair は通貨ペア
	Pair string `json:"pair"`
	// CreatedAt は作成日時
	CreatedAt time.Time `json:"created_at"`
}

// Print は Order を見やすくフォーマットして出力する
func (o CreateOrder) Print() {
	fmt.Println("=== Order Info ===")
	fmt.Printf(" Success        : %t\n", o.Success)
	fmt.Printf(" Order ID       : %d\n", o.ID)
	fmt.Printf(" Rate           : %s JPY\n", o.Rate)
	fmt.Printf(" Amount         : %s\n", o.Amount)
	fmt.Printf(" Order Type     : %s\n", o.OrderType)
	fmt.Printf(" Time In Force  : %s\n", o.TimeInForce)
	if o.StopLossRate != nil {
		fmt.Printf(" Stop Loss Rate : %s\n", *o.StopLossRate)
	} else {
		fmt.Println(" Stop Loss Rate : nil")
	}
	fmt.Printf(" Pair           : %s\n", o.Pair)
	fmt.Printf(" Created At     : %s\n", o.CreatedAt.Format("2006-01-02 15:04:05"))
	fmt.Println("===================")
}
