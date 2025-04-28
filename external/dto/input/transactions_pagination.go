package input

type TransactionsPagination struct {
	Limit         int    `json:"limit"`
	Order         string `json:"order"`
	StartingAfter *int   `json:"starting_after"`
	EndingBefore  *int   `json:"ending_before"`
}
