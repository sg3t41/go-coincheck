package output

import (
	"fmt"
)

// PaginationOrder はページネーション情報を表す構造体
type PaginationOrder struct {
	Limit         int    `json:"limit"`          // 取得する取引の最大数
	Order         string `json:"order"`          // 取得する順序 ("asc" または "desc")
	StartingAfter *int   `json:"starting_after"` // 指定されたIDの取引の次の取引から取得
	EndingBefore  *int   `json:"ending_before"`  // 指定されたIDの取引の前の取引まで取得
}

// TransactionsPagination はページネーション付き取引履歴のレスポンスを表す構造体
type TransactionsPagination struct {
	Success    bool            `json:"success"`    // リクエストの成功を示す
	Pagination PaginationOrder `json:"pagination"` // ページネーション情報
	Data       []Transaction   `json:"data"`       // 取引履歴
}

// Print は GetTransactionsPagination を見やすくフォーマットして出力する
func (t TransactionsPagination) Print() {
	fmt.Println("=== ページネーション付き取引履歴 ===")
	fmt.Printf(" リクエストの成功: %t\n", t.Success)
	fmt.Printf(" ページネーション情報:\n")
	fmt.Printf("  取得する取引の最大数: %d\n", t.Pagination.Limit)
	fmt.Printf("  取得する順序       : %s\n", t.Pagination.Order)
	if t.Pagination.StartingAfter != nil {
		fmt.Printf("  次の取引IDから取得 : %d\n", *t.Pagination.StartingAfter)
	}
	if t.Pagination.EndingBefore != nil {
		fmt.Printf("  前の取引IDまで取得 : %d\n", *t.Pagination.EndingBefore)
	}
	fmt.Println("---")
	for _, transaction := range t.Data {
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
		fmt.Println("---")
	}
	fmt.Println("==================")
}
