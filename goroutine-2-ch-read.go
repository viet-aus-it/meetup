package main

import "fmt"

func main() {
	done := make(chan bool)
	go withReadingChannel(done)

	<-done
}

func withReadingChannel(done chan<- bool) {
	fmt.Println("hi there")
	done <- true
}
