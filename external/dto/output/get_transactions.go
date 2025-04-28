package output

import (
	"fmt"
	"time"
)

// Transaction は取引情報を表す構造体
type Transaction struct {
	ID          int       `json:"id"`           // 取引ID
	OrderID     int       `json:"order_id"`     // 注文ID
	CreatedAt   time.Time `json:"created_at"`   // 取引日時
	Funds       Funds     `json:"funds"`        // 各残高の増減分
	Pair        string    `json:"pair"`         // 取引ペア
	Rate        string    `json:"rate"`         // 約定価格
	FeeCurrency string    `json:"fee_currency"` // 手数料の通貨
	Fee         string    `json:"fee"`          // 発生した手数料
	Liquidity   string    `json:"liquidity"`    // "T" (Taker) or "M" (Maker) or "itayose" (Itayose)
	Side        string    `json:"side"`         // "sell" or "buy"
}

// Funds は各残高の増減分を表す構造体
type Funds struct {
	BTC string `json:"btc"` // ビットコインの増減分
	JPY string `json:"jpy"` // 日本円の増減分
}

// GetTransactions は取引履歴のレスポンスを表す構造体
type GetTransactions struct {
	Success      bool          `json:"success"`      // リクエストの成功を示す
	Transactions []Transaction `json:"transactions"` // 取引履歴
}

// Print は Transactions を見やすくフォーマットして出力する
func (t GetTransactions) Print() {
	fmt.Println("=== 取引履歴 ===")
	for _, transaction := range t.Transactions {
		fmt.Println("---")
		fmt.Printf(" 取引ID            : %d\n", transaction.ID)
		fmt.Printf(" 注文ID            : %d\n", transaction.OrderID)
		fmt.Printf(" 取引日時          : %s\n", transaction.CreatedAt.Format("2006-01-02 15:04:05"))
		fmt.Printf(" ビットコインの増減: %s BTC\n", transaction.Funds.BTC)
		fmt.Printf(" 日本円の増減      : %s 円\n", transaction.Funds.JPY)
		fmt.Printf(" 取引ペア          : %s\n", transaction.Pair)
		fmt.Printf(" 約定価格          : %s 円\n", transaction.Rate)
		fmt.Printf(" 手数料の通貨      : %s\n", transaction.FeeCurrency)
		fmt.Printf(" 発生した手数料    : %s\n", transaction.Fee)
		fmt.Printf(" 流動性            : %s\n", transaction.Liquidity)
		fmt.Printf(" 売買区分          : %s\n", transaction.Side)
	}
	fmt.Println("==================")
}
