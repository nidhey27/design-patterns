package observer

type Publisher interface {
	subsribe(subscriber Subscriber)
	unsubscribe(subscriber Subscriber)
	notify()
	getTemprature()
}
