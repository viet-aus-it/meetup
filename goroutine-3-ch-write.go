package main

import (
	"fmt"
	"time"
)

func main() {
	stop := make(chan bool)
	go withWritingChannel(stop)

	time.Sleep(5 * time.Second)
	stop <- true
}

func withWritingChannel(stop <-chan bool) {
	for {
		fmt.Println("hi there")

		select {
		case <-stop:
			break

		case <-time.After(time.Second):
			continue
		}
	}
}
