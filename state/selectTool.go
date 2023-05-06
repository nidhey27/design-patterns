package state

import "fmt"

type SelectionTool struct{}

func (s *SelectionTool) MouseDown() {
	fmt.Println("Selection Icon")
}

func (s *SelectionTool) MouseUp() {
	fmt.Println("Draw a dashed rectangle")
}
