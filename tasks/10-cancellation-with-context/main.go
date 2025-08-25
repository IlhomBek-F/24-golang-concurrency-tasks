package main

import (
	"context"
	"fmt"
	"time"
)

// 🔹 Task 11: Cancellation with context.Context
// Learn how to cancel goroutines gracefully using Go’s context package.

func worker(ctx context.Context, id int) {
	fmt.Printf("Worker %d started\n", id)
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("worker %d stopped\n", id)
			return
		default:
			fmt.Printf("worker %d working...\n", id)
			time.Sleep(time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	for i := range 5 {
		go worker(ctx, i+1)
	}

	time.Sleep(time.Second * 3)

	cancel()

	time.Sleep(time.Second * 1)
	fmt.Println("All workers stopped")
}
