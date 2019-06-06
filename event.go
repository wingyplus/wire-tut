package main

import "fmt"

type Message string

func NewMessage() Message {
	return Message("Hi, there!")
}

type Greeter struct {
	Message Message
}

func (g *Greeter) Greet() string {
	return string(g.Message) 
}

func NewGreeter(msg Message) Greeter {
	return Greeter{Message: msg}
}

type Event struct {
	Greeter Greeter
}

func (evt *Event) Start() {
	fmt.Println(evt.Greeter.Greet())
}

func NewEvent(greeter Greeter) Event {
	return Event{Greeter: greeter}
}
