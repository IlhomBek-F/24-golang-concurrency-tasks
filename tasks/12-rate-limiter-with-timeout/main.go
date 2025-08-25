package main

import (
	"context"
	"fmt"
	"time"
)

// 🔹 Task 13: Rate-Limited Workers with Timeout
// Start workers that process tasks at a limited rate (e.g., 1 task per second).
// Automatically stop all workers after a timeout using context.WithTimeout.

func worker(ctx context.Context, workerIndex, task int) {
	select {
	case <-ctx.Done():
		return
	default:
		fmt.Printf("Worker %d processing task %d\n", workerIndex, task)
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
	ticker := time.NewTicker(time.Second)

	defer ticker.Stop()
	defer cancel()

	for i := range 5 {
		<-ticker.C
		go worker(ctx, i+1, i+1)
	}

	<-ctx.Done()

	fmt.Println("Timeout reached! Stopping all workers")
}
