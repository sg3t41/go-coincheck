package output

import (
	"fmt"
	"time"
)

// GetTicker は GetTicker のレスポンス
type GetTicker struct {
	// Last は最新の約定価格
	Last float64 `json:"last"`
	// Bid は現在の最高買い注文（最も高い買い希望価格）
	Bid float64 `json:"bid"`
	// Ask は現在の最安売り注文（最も安い売り希望価格）
	Ask float64 `json:"ask"`
	// High は過去24時間の最高価格
	High float64 `json:"high"`
	// Low は過去24時間の最安価格
	Low float64 `json:"low"`
	// Volume は過去24時間の取引量
	Volume float64 `json:"volume"`
	// Timestamp はデータ取得時のUNIXタイムスタンプ
	Timestamp float64 `json:"timestamp"`
}

// Print は GetTickerResponse を見やすくフォーマットして出力する
func (r GetTicker) Print() {
	t := time.Unix(int64(r.Timestamp), 0).Format("2006-01-02 15:04:05")

	fmt.Println("=== Ticker Info ===")
	fmt.Printf(" Latest Price  : %10.2f JPY\n", r.Last)
	fmt.Printf(" Highest Bid   : %10.2f JPY\n", r.Bid)
	fmt.Printf(" Lowest Ask    : %10.2f JPY\n", r.Ask)
	fmt.Printf(" 24h High      : %10.2f JPY\n", r.High)
	fmt.Printf(" 24h Low       : %10.2f JPY\n", r.Low)
	fmt.Printf(" 24h Volume    : %10.2f\n", r.Volume)
	fmt.Printf(" Timestamp     : %s\n", t)
	fmt.Println("===================")
}
