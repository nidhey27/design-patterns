package main

import "fmt"

type Notifier struct {
	wrappee BaseDecorator
}

func (n *Notifier) send(message string) {
	n.wrappee.send(message)
}

type BaseDecorator interface {
	send(message string)
}

type SMSDecorator struct {
}

func (s *SMSDecorator) send(message string) {
	fmt.Println("Sending SMS", message, "notification.")
}

type WhatsappDecorator struct {
}

func (w *WhatsappDecorator) send(message string) {
	fmt.Println("Sending Whatsapp", message, "notification.")
}

type SlackDecorator struct {
}

func (w *SlackDecorator) send(message string) {
	fmt.Println("Sending Slack", message, "notification.")
}

func main() {
	fmt.Println("Decorator Pattern")

	notifier := &Notifier{
		wrappee: &SMSDecorator{},
	}
	notifier.send("[ALERT] You have been hacked.")

	notifier.wrappee = &SlackDecorator{}
	notifier.send("[ALERT] You have been hacked.")

	notifier.wrappee = &WhatsappDecorator{}
	notifier.send("[ALERT] You have been hacked.")
}
