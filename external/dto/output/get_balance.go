package output

import (
	"fmt"
)

// Balance はアカウントの残高情報を持つ構造体
type Balance struct {
	Success      bool   `json:"success"`
	JPY          string `json:"jpy"`
	BTC          string `json:"btc"`
	JPYReserved  string `json:"jpy_reserved"`
	BTCReserved  string `json:"btc_reserved"`
	JPYLending   string `json:"jpy_lending"`
	BTCLending   string `json:"btc_lending"`
	JPYLendInUse string `json:"jpy_lend_in_use"`
	BTCLendInUse string `json:"btc_lend_in_use"`
	JPYLent      string `json:"jpy_lent"`
	BTCLent      string `json:"btc_lent"`
	JPYDebt      string `json:"jpy_debt"`
	BTCDebt      string `json:"btc_debt"`
	JPYTsumitate string `json:"jpy_tsumitate"`
	BTCTsumitate string `json:"btc_tsumitate"`
}

// Print は AccountBalance を見やすくフォーマットして出力する
func (a Balance) Print() {
	fmt.Println("=== アカウント残高情報 ===")
	fmt.Printf(" 成功             : %t\n", a.Success)
	fmt.Printf(" 日本円残高       : %s 円\n", a.JPY)
	fmt.Printf(" ビットコイン残高 : %s BTC\n", a.BTC)
	fmt.Printf(" 未決済の日本円   : %s 円\n", a.JPYReserved)
	fmt.Printf(" 未決済のビットコイン : %s BTC\n", a.BTCReserved)
	fmt.Printf(" 貸出前の日本円   : %s 円\n", a.JPYLending)
	fmt.Printf(" 貸出前のビットコイン : %s BTC\n", a.BTCLending)
	fmt.Printf(" 貸出申請中の日本円 : %s 円\n", a.JPYLendInUse)
	fmt.Printf(" 貸出申請中のビットコイン : %s BTC\n", a.BTCLendInUse)
	fmt.Printf(" 貸出中の日本円   : %s 円\n", a.JPYLent)
	fmt.Printf(" 貸出中のビットコイン : %s BTC\n", a.BTCLent)
	fmt.Printf(" 借りている日本円 : %s 円\n", a.JPYDebt)
	fmt.Printf(" 借りているビットコイン : %s BTC\n", a.BTCDebt)
	fmt.Printf(" つみたて中の日本円 : %s 円\n", a.JPYTsumitate)
	fmt.Printf(" つみたて中のビットコイン : %s BTC\n", a.BTCTsumitate)
	fmt.Println("==========================")
}
