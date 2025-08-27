package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 21 Task Multi-Service Shutdown
// You are building a system with 3 independent services running concurrently. Each service processes tasks in a loop.

func worker(ctx context.Context, task int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stopped")
			return
		default:
			fmt.Println("processing task", task)
			time.Sleep(time.Second * 2)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	wg.Add(1)
	go worker(ctx, 1, &wg)

	time.Sleep(time.Second * 5)

	cancel()
	wg.Wait()
}
