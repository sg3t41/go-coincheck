package output

import (
	"fmt"
	"sort"
)

// Accounts はアカウント情報を持つ構造体
type Accounts struct {
	Success        bool                       `json:"success"`
	ID             int                        `json:"id"`
	Email          string                     `json:"email"`
	IdentityStatus string                     `json:"identity_status"`
	BitcoinAddress string                     `json:"bitcoin_address"`
	TakerFee       string                     `json:"taker_fee"`
	MakerFee       string                     `json:"maker_fee"`
	ExchangeFees   map[string]ExchangeFeeRate `json:"exchange_fees"`
}

// ExchangeFeeRate は板ごとの手数料情報を持つ構造体
type ExchangeFeeRate struct {
	MakerFeeRate string `json:"maker_fee_rate"`
	TakerFeeRate string `json:"taker_fee_rate"`
}

// Print はアカウント情報を見やすくフォーマットして出力する
func (a Accounts) Print() {
	fmt.Println("=== アカウント情報 ===")
	fmt.Printf(" 成功             : %t\n", a.Success)
	fmt.Printf(" アカウントID     : %d\n", a.ID)
	fmt.Printf(" メールアドレス   : %s\n", a.Email)
	fmt.Printf(" 本人確認状況     : %s\n", a.IdentityStatus)
	fmt.Printf(" ビットコインアドレス : %s\n", a.BitcoinAddress)
	fmt.Printf(" Taker手数料      : %s%%\n", a.TakerFee)
	fmt.Printf(" Maker手数料      : %s%%\n", a.MakerFee)
	fmt.Println("=== 板ごとの手数料 ===")
	fmt.Printf("%-10s %-12s %-12s\n", "ペア", "Maker手数料", "Taker手数料")

	// マップのキーをソート
	keys := make([]string, 0, len(a.ExchangeFees))
	for k := range a.ExchangeFees {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// ソートされたキーで出力
	for _, pair := range keys {
		fee := a.ExchangeFees[pair]
		fmt.Printf("%-10s %-12s %-12s\n", pair, fee.MakerFeeRate, fee.TakerFeeRate)
	}

	fmt.Println("======================")
}
