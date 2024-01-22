package main

import "fmt"

type IPublisher interface {
	Register(subscriber *Subscriber)
	NotifyAll(msg string)
}

type ISubscriber interface {
	ReactToPublisherMsg(msg string)
}

type Publisher struct {
	subscribers []*Subscriber
}

func GetPublisher() *Publisher {
	return &Publisher{
		subscribers: make([]*Subscriber, 0),
	}
}

func (p *Publisher) Register(subscriber *Subscriber) {
	p.subscribers = append(p.subscribers, subscriber)
}

func (p *Publisher) NotifyAll(msg string) {
	for _, subscriber := range p.subscribers {
		fmt.Println("Publisher notifying Subscriber", subscriber.name)
		subscriber.ReactToPublisherMsg(msg)
	}
}

type Subscriber struct {
	name string
}

func GetNewSubscriber(name string) *Subscriber {
	return &Subscriber{name: name}
}

func (s *Subscriber) ReactToPublisherMsg(msg string) {
	fmt.Println(s.name, "received message", msg)
}

func main() {
	publisher := GetPublisher()

	subscriberOne := GetNewSubscriber("Subscriber 1")
	subscriberTwo := GetNewSubscriber("Subscriber 2")

	publisher.Register(subscriberOne)
	publisher.Register(subscriberTwo)

	publisher.NotifyAll("Wassup, Subscribers?")
}
