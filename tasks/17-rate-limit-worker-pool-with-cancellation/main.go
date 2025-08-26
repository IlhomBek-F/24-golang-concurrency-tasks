package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 🔹 Task 18: Rate-Limited Worker Pool with Context Cancellation
// Combine worker pools, rate limiting, and early cancellation using context.Context.

func worker(ctx context.Context, tasks <-chan int, result chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Timeout reached. Stopping task feeder.")
			return
		case task := <-tasks:
			result <- task
		}
	}
}

func main() {
	tasks := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	taskChan := make(chan int)
	resultChan := make(chan int)

	var wg sync.WaitGroup
	wg.Add(3)
	for i := 1; i <= 3; i++ {
		go worker(ctx, taskChan, resultChan, &wg)
	}

	go func() {
		for _, task := range tasks {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				taskChan <- task
			}
		}
		close(taskChan)
	}()

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for result := range resultChan {
		fmt.Println("Processed task:", result)
	}
	fmt.Println("All done!")
}
