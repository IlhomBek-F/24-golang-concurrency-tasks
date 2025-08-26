package main

import (
	"fmt"
	"sync"
)

// 🔹 Task 8: Worker Pool (Fixed Number of Workers)
// Instead of starting one goroutine per task (like before),
// create a fixed pool of N workers that continuously process tasks from a channel. This is the classic worker pool pattern in Go.

func worker(workerIndex int, tasks <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		fmt.Printf("Worker %d processing task %d\n", workerIndex, task)
	}
}

func main() {
	tasks := make(chan int)
	var wg sync.WaitGroup

	wg.Add(3)
	go worker(1, tasks, &wg)
	go worker(2, tasks, &wg)
	go worker(3, tasks, &wg)

	for i := range 10 {
		tasks <- i
	}
	close(tasks)
	wg.Wait()
}
