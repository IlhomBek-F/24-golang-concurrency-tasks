package main

import (
	"context"
	"fmt"
	"time"
)

// 🔹 Task 12: Goroutines with Timeout
// Start multiple workers that do some work, but automatically stop all workers after a timeout using Go’s context.WithTimeout.

func worker(ctx context.Context, task int) {
	fmt.Printf("Worker %d started...\n", task)
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d stopped due to time out\n", task)
			return
		default:
			fmt.Printf("Worker %d working...\n", task)
			time.Sleep(time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	for i := 0; i < 5; i++ {
		go worker(ctx, i+1)
	}

	<-ctx.Done()

	fmt.Println("All worker stopped")
}
