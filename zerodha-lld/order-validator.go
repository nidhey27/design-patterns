package main

import (
	"fmt"
)

type OrderValidator struct {
	System *System
}

func (ov *OrderValidator) ValidateOrder(userID string, order *Order) bool {
	user := ov.System.GetUser(userID)
	if order.GetTxnType() == BUY {
		fmt.Println("Checking if ", user.UserID, " had FUNDS to BUY.")
	} else {
		fmt.Println("Checking if ", user.UserID, " has STOCKS to SELL.")
	}
	fmt.Println("Assuming user has made a valid request.")
	return true
}
