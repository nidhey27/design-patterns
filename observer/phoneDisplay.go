package observer

import "fmt"

type PhoneDisplay struct{}

func (pd *PhoneDisplay) update(context string) {
	fmt.Println("Notified to Phone Display")
}

func (pd *PhoneDisplay) display(context string) {
	fmt.Printf("Phone Display => %v\n", context)
}
