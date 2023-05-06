package main

import (
	"fmt"

	"github.com/nidhey27/design-patterns/memento"
	"github.com/nidhey27/design-patterns/strategy"
)

func main() {
	e := &memento.Editor{}
	h := &memento.History{}

	e.SetContent("A")
	h.Push(*e.CreateState())

	e.SetContent("B")
	h.Push(*e.CreateState())

	e.SetContent("C")
	e.Restore(h.Pop())
	e.Restore(h.Pop())

	fmt.Println(e.GetContent())

	cart := &strategy.Cart{}

	cart.SetCart(115)

	paymentStrategy1 := &strategy.CardStrategy{}
	paymentStrategy1.SetCardStrategy("Nidhey", "23456789", "000", "2023-09-11")
	cart.Pay(paymentStrategy1)

	paymentStrategy2 := &strategy.UPIStrategy{}
	paymentStrategy2.SetUPIStrategy("234567812@test")
	cart.Pay(paymentStrategy2)
}
