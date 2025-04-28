package output

// CreateOrder は CreateOrder のレスポンス
type CancelOrder struct {
	Success bool `json:"success"`
	ID      int  `json:"id"`
}
