package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.TODO())
	go withContext(ctx)

	time.Sleep(5 * time.Second)
	cancel()
}

func withContext(ctx context.Context) {
	for {
		fmt.Println("hi there")

		select {
		case <-ctx.Done():
			break

		case <-time.After(time.Second):
			continue
		}
	}
}
