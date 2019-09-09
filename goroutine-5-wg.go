package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go withWg(wg)
	wg.Wait()
}

func withWg(wg *sync.WaitGroup) {
	fmt.Println("hi there")
	wg.Done()
}
