package main

import (
	"fmt"
	"sync"
)

func main() {
	log := []string{}
	wg := sync.WaitGroup{}
	mu := &sync.Mutex{}

	race := func() {
		mu.Lock()
		defer mu.Unlock()
		defer wg.Done()

		log = append(log, "hi")
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go race()
	}
	wg.Wait()

	fmt.Println(log)
}
