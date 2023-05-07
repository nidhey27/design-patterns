package observer

import "fmt"

type LEDDisplay struct{}

func (pd *LEDDisplay) update(context string) {
	fmt.Println("Notified to LED Display")
}

func (pd *LEDDisplay) display(context string) {
	fmt.Printf("LED Display => %v\n", context)
}
