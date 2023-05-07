package observer

type Subscriber interface {
	update(context string)
	display(message string)
}
