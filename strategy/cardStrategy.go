package strategy

import "fmt"

type CardStrategy struct {
	name       string
	cardNumber string
	cvv        string
	dateOfExp  string
}

func (cs *CardStrategy) SetCardStrategy(name, cardNumber, cvv, dateOfExp string) {
	cs.name = name
	cs.cardNumber = cardNumber
	cs.cvv = cvv
	cs.dateOfExp = dateOfExp
}

func (cs *CardStrategy) pay(amount int) {
	fmt.Println("Card Details:")
	fmt.Println("Name", cs.name)
	fmt.Println("Card Number", cs.cardNumber)
	fmt.Println("CVV", cs.cvv)
	fmt.Println("Date Of Expiry", cs.dateOfExp)
	fmt.Printf("Paying $%v by Credit/Debit Card\n", amount)
}
