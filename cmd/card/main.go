package main

import (
	"fmt"
	"go_hw4/pkg/card"
	"go_hw4/pkg/transfer"
)

func main() {
	masterCard := &card.Card{
		Id:       0,
		Issuer:   "MasterCard",
		Balance:  45_000_00,
		Currency: "RUB",
		Number:   "0808",
		Icon:     "https://dlpng.com/png/6794578",
	}
	visa := &card.Card{
		Id:       1,
		Issuer:   "Visa",
		Balance:  3400_00,
		Currency: "RUB",
		Number:   "0002",
		Icon:     "https://...",
	}

	bank := card.NewService("Tinkoff")
	bankTransfers := transfer.NewService(bank, 0.5, 10_00)

	bank.AddCards([]*card.Card{masterCard, visa})

	total1, ok1 := bankTransfers.Card2Card("0002", "0001", 3500_00)
	total2, ok2 := bankTransfers.Card2Card("0808", "0002", 3500_00)

	fmt.Println(bank)
	fmt.Println(total1, ok1)
	fmt.Println(total2, ok2)
}
