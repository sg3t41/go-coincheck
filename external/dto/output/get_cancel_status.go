package output

import (
	"fmt"
	"strings"
)

// CancelOrder は注文キャンセルのレスポンス
type CancelStatus struct {
	Success   bool   `json:"success"`    // 成功フラグ
	ID        int    `json:"id"`         // 注文のID
	Cancel    bool   `json:"cancel"`     // キャンセル済みか（true or false）
	CreatedAt string `json:"created_at"` // 注文の作成日時
}

// Print は注文キャンセルのステータスを一覧表示
func (c *CancelStatus) Print() {
	fmt.Println("\n=== 注文キャンセルのステータス ===")
	fmt.Println(strings.Repeat("=", 40))
	fmt.Printf("%-10s %-12s %-20s\n",
		"ID", "Cancel", "Created At (JST)")
	fmt.Println(strings.Repeat("=", 40))

	// JST に変換
	createdAtJST := toJST(c.CreatedAt)

	fmt.Printf("%-10d %-12v %-20s\n",
		c.ID, c.Cancel, createdAtJST)
	fmt.Println(strings.Repeat("=", 40))
}
