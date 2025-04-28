package output

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

// GetTrades は Coincheck API の取引履歴を表す
type GetTrades struct {
	Success    bool       `json:"success"`
	Pagination Pagination `json:"pagination"`
	Data       []Trade    `json:"data"`
}

// Pagination はページネーション情報
type Pagination struct {
	Limit         int    `json:"limit"`
	Order         string `json:"order"`
	StartingAfter *int   `json:"starting_after"`
	EndingBefore  *int   `json:"ending_before"`
}

// Trade は取引データ
type Trade struct {
	ID        int    `json:"id"`
	Amount    string `json:"amount"`
	Rate      string `json:"rate"`
	Pair      string `json:"pair"`
	OrderType string `json:"order_type"`
	CreatedAt string `json:"created_at"`
}

// SortedAndGroupedTrades はソート済みの取引データを格納する構造体
type SortedAndGroupedTrades struct {
	BuyTrades  []Trade
	SellTrades []Trade
}

// SortAndGroupTrades は取引を新しい順にソートし、買い・売りで分ける（JSTに変換）
func (gt *GetTrades) SortAndGroupTrades() SortedAndGroupedTrades {
	// JST タイムゾーンを取得
	loc, _ := time.LoadLocation("Asia/Tokyo")

	// 取引を新しい順にソート（降順）
	sort.Slice(gt.Data, func(i, j int) bool {
		// 日付をパースして JST に変換
		ti, _ := time.Parse(time.RFC3339, gt.Data[i].CreatedAt)
		tj, _ := time.Parse(time.RFC3339, gt.Data[j].CreatedAt)
		ti = ti.In(loc) // JST に変換
		tj = tj.In(loc)
		gt.Data[i].CreatedAt = ti.Format(time.RFC3339) // JST で更新
		gt.Data[j].CreatedAt = tj.Format(time.RFC3339)

		// 降順ソート（新しいものが先頭）
		return ti.After(tj)
	})

	// 買い（buy）と売り（sell）で分類
	var buyTrades []Trade
	var sellTrades []Trade

	for _, trade := range gt.Data {
		if trade.OrderType == "buy" {
			buyTrades = append(buyTrades, trade)
		} else if trade.OrderType == "sell" {
			sellTrades = append(sellTrades, trade)
		}
	}

	// 結果を返す
	return SortedAndGroupedTrades{
		BuyTrades:  buyTrades,
		SellTrades: sellTrades,
	}
}

func (gt *GetTrades) Print() {
	ts := gt.SortAndGroupTrades()
	if len(ts.BuyTrades)+len(ts.SellTrades) == 0 {
		fmt.Println("取引なし")
		return
	}

	fmt.Println("売り注文")
	fmt.Printf("%-10s %-10s %-10s %-25s\n", "ID", "Amount", "Rate", "CreatedAt (JST)")
	fmt.Println(strings.Repeat("-", 60))

	for _, t := range ts.BuyTrades {
		fmt.Printf("%-10d %-10s %-10s %-25s\n", t.ID, t.Amount, t.Rate, t.CreatedAt)
	}

	fmt.Println("買い注文")
	fmt.Printf("%-10s %-10s %-10s %-25s\n", "ID", "Amount", "Rate", "CreatedAt (JST)")
	fmt.Println(strings.Repeat("-", 60))

	for _, t := range ts.SellTrades {
		fmt.Printf("%-10d %-10s %-10s %-25s\n", t.ID, t.Amount, t.Rate, t.CreatedAt)
	}

}
