package state

import "fmt"

type BrushTool struct{}

func (s *BrushTool) MouseDown() {
	fmt.Println("Brush Icon")
}

func (s *BrushTool) MouseUp() {
	fmt.Println("Draw a line")
}
