package main

import (
	"fmt"
	"time"
)

func main() {
	go doSomething()

	time.Sleep(1 * time.Second)
}

func doSomething() {
	fmt.Println("hi there")
}
