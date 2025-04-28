package output

import (
	"fmt"
	"time"
)

// GetOrder は注文のステータスを持つ構造体
type GetOrder struct {
	Success                 bool    `json:"success"`                    // 成功フラグ
	ID                      int     `json:"id"`                         // 注文のID
	Pair                    string  `json:"pair"`                       // 取引ペア
	Status                  string  `json:"status"`                     // 注文のステータス
	OrderType               string  `json:"order_type"`                 // 注文のタイプ
	Rate                    *string `json:"rate"`                       // 注文のレート (null の場合は成り行き注文)
	StopLossRate            *string `json:"stop_loss_rate"`             // 逆指値レート (null の場合もあり)
	MakerFeeRate            string  `json:"maker_fee_rate"`             // Makerとして注文を行った場合の手数料
	TakerFeeRate            string  `json:"taker_fee_rate"`             // Takerとして注文を行った場合の手数料
	Amount                  string  `json:"amount"`                     // 注文の量
	MarketBuyAmount         *string `json:"market_buy_amount"`          // 成行買で注文した日本円の金額 (null の場合もあり)
	ExecutedAmount          float64 `json:"executed_amount"`            // 約定した量
	ExecutedMarketBuyAmount *string `json:"executed_market_buy_amount"` // 成行買で約定した日本円の金額 (null の場合もあり)
	ExpiredType             string  `json:"expired_type"`               // 失効した理由
	PreventedMatchID        int     `json:"prevented_match_id"`         // 対当した注文のID
	ExpiredAmount           string  `json:"expired_amount"`             // 失効した量
	ExpiredMarketBuyAmount  *string `json:"expired_market_buy_amount"`  // 成行買で失効した日本円の金額 (null の場合もあり)
	TimeInForce             string  `json:"time_in_force"`              // 注文有効期間
	CreatedAt               string  `json:"created_at"`                 // 注文の作成日時
}

// Print は注文のステータスを見やすくフォーマットして出力する
func (o GetOrder) Print() {
	fmt.Println("=== 注文のステータス ===")
	fmt.Printf(" 成功             : %t\n", o.Success)
	fmt.Printf(" 注文ID           : %d\n", o.ID)
	fmt.Printf(" 取引ペア         : %s\n", o.Pair)
	fmt.Printf(" ステータス       : %s\n", o.Status)
	fmt.Printf(" 注文タイプ       : %s\n", o.OrderType)
	fmt.Printf(" レート           : %s\n", getOrNA(o.Rate))
	fmt.Printf(" 逆指値レート     : %s\n", getOrNA(o.StopLossRate))
	fmt.Printf(" Maker手数料      : %s\n", o.MakerFeeRate)
	fmt.Printf(" Taker手数料      : %s\n", o.TakerFeeRate)
	fmt.Printf(" 注文量           : %s\n", o.Amount)
	fmt.Printf(" 成行買金額       : %s\n", getOrNA(o.MarketBuyAmount))
	fmt.Printf(" 約定量           : %f\n", o.ExecutedAmount)
	fmt.Printf(" 約定成行買金額   : %s\n", getOrNA(o.ExecutedMarketBuyAmount))
	fmt.Printf(" 失効理由         : %s\n", o.ExpiredType)
	fmt.Printf(" 対当注文ID       : %d\n", o.PreventedMatchID)
	fmt.Printf(" 失効量           : %s\n", o.ExpiredAmount)
	fmt.Printf(" 失効成行買金額   : %s\n", getOrNA(o.ExpiredMarketBuyAmount))
	fmt.Printf(" 注文有効期間     : %s\n", o.TimeInForce)
	fmt.Printf(" 作成日時         : %s\n", toJST(o.CreatedAt))
	fmt.Println("========================")
}

// toJST はUTCの時間をJSTに変換
func toJST(utcTime string) string {
	t, err := time.Parse(time.RFC3339, utcTime)
	if err != nil {
		return "Invalid Time"
	}
	jst := t.In(time.FixedZone("JST", 9*60*60)) // UTC+9
	return jst.Format("2006-01-02 15:04:05")
}
