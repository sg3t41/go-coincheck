package output

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// GetOrderBooks はオーダーブックの情報を保持する
type GetOrderBooks struct {
	Asks [][]string `json:"asks"`
	Bids [][]string `json:"bids"`
}

// Print はオーダーブックを板情報のようにフォーマットして表示
func (o *GetOrderBooks) Print() {
	fmt.Println("\n=== Order Book ===")
	fmt.Println(strings.Repeat("-", 30))
	fmt.Printf("%10s | %10s\n", "Price", "Amount")
	fmt.Println(strings.Repeat("-", 30))

	// 売り注文を価格の高い順で表示
	asks := o.Asks
	sort.Slice(asks, func(i, j int) bool {
		a, _ := strconv.ParseFloat(asks[i][0], 64)
		b, _ := strconv.ParseFloat(asks[j][0], 64)
		return a > b // 価格が高い順
	})
	for _, ask := range asks {
		fmt.Printf("%10s | %10s  (Ask)\n", ask[0], ask[1])
	}

	// 中央の区切り線
	fmt.Println(strings.Repeat("-", 30))

	// 買い注文を価格の高い順で表示
	bids := o.Bids
	sort.Slice(bids, func(i, j int) bool {
		a, _ := strconv.ParseFloat(bids[i][0], 64)
		b, _ := strconv.ParseFloat(bids[j][0], 64)
		return a > b // 価格が高い順
	})
	for _, bid := range bids {
		fmt.Printf("%10s | %10s  (Bid)\n", bid[0], bid[1])
	}
	fmt.Println(strings.Repeat("-", 30))
}
