package input

type OrdersRate struct {
	OrderType string  `json:"order_type"`
	Amount    float64 `json:"amount"`
	Pair      string  `json:"pair"`
	Price     float64 `json:"price"`
}
