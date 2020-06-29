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
		Number:   "5106217896540004",
		Icon:     "https://dlpng.com/png/6794578",
	}
	visa := &card.Card{
		Id:       1,
		Issuer:   "Visa",
		Balance:  3400_00,
		Currency: "RUB",
		Number:   "5106218888990009",
		Icon:     "https://...",
	}

	bank := card.NewService("Tinkoff", "510621")
	bankTransfers := transfer.NewService(bank, 0.5, 10_00)

	bank.AddCards(masterCard, visa)

	// В этом кейсе недостаточно средств для перевода
	total, err := bankTransfers.Card2Card("5106218888990009", "5106217896540004", 3500_00)

	if err != nil {
		switch err {
		case transfer.ErrSourceCardBalanceNotEnough:
			fmt.Println("Sorry, insufficient funds")
		case transfer.ErrTargetCardNotFound:
			fmt.Println("Please check target card number")
		case transfer.ErrSourceCardBalanceNotEnough:
			fmt.Println("Please check source card number")
		case transfer.ErrInvalidSourceCardNumber:
			fmt.Println("Source card number is invalid")
		case transfer.ErrInvalidTargetCardNumber:
			fmt.Println("Target card number is invalid")
		default:
			fmt.Println("Something went wrong. Try again later")
		}
	} else {
		fmt.Println("Your transfer amount was successful! Amount:", total)
	}
}
