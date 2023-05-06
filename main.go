package main

import (
	"fmt"

	"github.com/nidhey27/design-patterns/memento"
	"github.com/nidhey27/design-patterns/state"
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


	canvas := &state.Canvas{}

	canvas.SetTool(&state.SelectionTool{})
	canvas.MouseDown()
	canvas.MouseUp()

	canvas.SetTool(&state.BrushTool{})
	canvas.MouseDown()
	canvas.MouseUp()


}
