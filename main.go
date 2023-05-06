package main

import (
	"fmt"

	"github.com/nidhey27/design-patterns/memento"
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
}
