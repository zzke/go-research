package main

import "golang.design/x/mainthread"

func main() {
	mainthread.Init(fn)
}

func fn() {
	// ... do what ever we want to do in main ...
	done := make(chan struct{})
	go gn(done)
	<-done
}

func gn(done chan<- struct{}) {
	// Wherever gn is running, the call will be executed on the main thread.
	mainthread.Call(func() {
		// ... do whatever we want to run on the main thread ...
		println("call on the main thread.")
	})
	done <- struct{}{}
}
