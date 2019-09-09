package main

import (
	"fmt"
	"time"
)

func main() {
	bobCh := make(chan string)
	aliceCh := make(chan string)
	go hiAlice(aliceCh)
	go hiBob(bobCh)
	ch := fanIn(aliceCh, bobCh)

	for {
		fmt.Println("hello", <-ch)
		time.Sleep(time.Second)
	}
}

func hiBob(ch chan<- string) {
	for {
		ch <- "from Bob"
	}
}

func hiAlice(ch chan<- string) {
	for {
		ch <- "from Alice"
	}
}

func fanIn(channels ...<-chan string) <-chan string {
	in := make(chan string)

	for _, ch := range channels {
		go func(ch <-chan string) {
			for {
				in <- <-ch
			}
		}(ch)
	}

	return in
}
