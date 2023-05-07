package main

import (
	"fmt"

	"github.com/nidhey27/design-patterns/memento"
	"github.com/nidhey27/design-patterns/observer"
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

	ws := &observer.WeatherStation{}

	pd := &observer.PhoneDisplay{}
	pd2 := &observer.PhoneDisplay{}
	ld := &observer.LEDDisplay{}

	ws.Subsribe(pd)
	ws.Subsribe(ld)
	ws.Subsribe(pd2)
	ws.Unsubscribe(ld)

	ws.SetTemprature(62.6)
	ws.Notify()

}
