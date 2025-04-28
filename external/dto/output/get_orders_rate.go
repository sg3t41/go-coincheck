package output

import "fmt"

type OrdersRate struct {
	Success bool   `json:"success"`
	Rate    string `json:"rate"`
	Price   string `json:"price"`
	Amount  string `json:"amount"`
}

// Printメソッドを追加
func (r *OrdersRate) Print() {
	fmt.Println("＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝")
	fmt.Println("取引所注文レート:")
	fmt.Println("＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝")
	fmt.Printf("  Success: %v\n", r.Success)
	fmt.Printf("  Rate   : %s 円/DOGE\n", r.Rate)
	fmt.Printf("  Price  : %s 円\n", r.Price)
	fmt.Printf("  Amount : %s DOGE\n", r.Amount)
	fmt.Println("＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝＝")
}
