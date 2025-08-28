package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Task 22 Timeout + Heartbeat Monitor

func worker(ctx context.Context, tasks <-chan int, workerInd int, heartbeat chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case task := <-tasks:
			fmt.Printf("Worker %d processing task %d\n", workerInd, task)
			heartbeat <- fmt.Sprintf("Worker %d processed task %d\n", workerInd, task)
		case <-ctx.Done():
			fmt.Printf("Worker %d stopped by context\n", workerInd)
			return
		}
	}
}

func main() {
	tasks := make(chan int)
	heartbeatChan := make(chan string, 10)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var wg sync.WaitGroup

	wg.Add(1)
	go worker(ctx, tasks, 1, heartbeatChan, &wg)

	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second * 2)
			tasks <- i
		}
		close(tasks)
	}()

	go func() {
		for signal := range heartbeatChan {
			fmt.Println(signal)
		}
	}()

	wg.Wait()
	close(heartbeatChan)
	fmt.Println("wait")
}
