package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)

type AccountModel struct {
	AccountId string          `json:"accountId"`
	Balance   decimal.Decimal `json:"balance"`
}

func main() {
	balance := decimal.NewFromFloat(9.9)
	fmt.Println(balance)
	money := balance.Add(decimal.NewFromFloat(9.9))
	fmt.Println("balance", balance)
	fmt.Println("money", money)

	balance, _ = decimal.NewFromString("19.9")
	fmt.Println("balance", balance)

	balance = balance.Add(decimal.NewFromFloat(199.9))
	fmt.Println("balance", balance)
}
