package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 🔹 Task 14: Worker Pool + Result Collection + Timeout + Rate Limiting
// Combine multiple concurrency concepts into one task:

// Worker Pool – fixed number of workers processing tasks.
// Result Collection – workers send results back to main via a channel.
// Timeout – stop all processing automatically after a specified time.
// Rate Limiting – ensure tasks are picked up at a controlled rate (e.g., 1 per second).

func worker(tasks <-chan int, ctx context.Context, result chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case task := <-tasks:
			result <- fmt.Sprintf("Worker processed task %d -> %d", task, task*task)
		case <-ctx.Done():
			return
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	taskChan := make(chan int)
	resulChan := make(chan string)

	var wg sync.WaitGroup

	defer cancel()

	for range 5 {
		wg.Add(1)
		go worker(taskChan, ctx, resulChan, &wg)
	}

	go func() {
		ticker := time.NewTicker(time.Second * 1)
		defer ticker.Stop()

		for i := 0; i < 5; i++ {
			select {
			case <-ctx.Done():
				close(taskChan)
				return
			case <-ticker.C:
				taskChan <- i + 1
			}
		}
	}()

	go func() {
		wg.Wait()
		close(resulChan)
	}()

	for result := range resulChan {
		fmt.Println(result)
	}
	fmt.Println("Timeout reached! Stopping all workers.")
}
