package main

func main() {
	// this is original way to do dependency injection.
	// evt := NewEvent(NewGreeter(NewMessage()))
	// evt.Start()

	// using wire.
	evt := InitializeEvent()
	evt.Start()
}
