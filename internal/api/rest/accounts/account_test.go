package accounts

import (
	"fmt"
	"log"
	"testing"

	"github.com/sg3t41/go-coincheck/internal/client"
)

var c client.Client

func init() {
	c, err := client.New("", "")
	if err != nil {
		log.Fatalln(err)
	}
	c = c
}

func TestAccounts(t *testing.T) {
	target := New(c)
	fmt.Println(target)
}
