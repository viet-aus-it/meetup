package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    threads := 5
    locks := make([]*sync.Mutex, threads)

    for i := 0; i < threads; i++ {
        locks[i] = &sync.Mutex{}
    }

    for i := 0; i < 100; i++ {
        i := i
        lock := locks[i%5]
        lock.Lock()

        fmt.Println("lock", i, " => ", i%5, lock)

        go func() {
            fmt.Println("processing: ", i)
            time.Sleep(5 * time.Second)

            lock.Unlock()
        }()
    }
}
