package main

import (
	"fmt"
	"sync"
)

// 🔹 Task 17: Worker Pool with Priorities
// Implement a worker pool that processes tasks, but each task has a priority.

type Task struct {
	ID       int
	Priority int
}

func process(highChan, lowChan <-chan Task, workerInd int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case high, ok := <-highChan:
			if !ok {
				return
			}
			fmt.Printf("Worker %d processing task %d priority %d\n", workerInd, high.ID, high.Priority)
		case lowChan, ok := <-lowChan:
			if !ok {
				return
			}
			fmt.Printf("Worker %d processing task %d priority %d\n", workerInd, lowChan.ID, lowChan.Priority)
		}
	}
}

func main() {
	tasks := []Task{{ID: 1, Priority: 10}, {ID: 2, Priority: 3}, {ID: 3, Priority: 5}}
	highChan := make(chan Task)
	lowChan := make(chan Task)

	var wg sync.WaitGroup

	wg.Add(2)
	go process(highChan, lowChan, 1, &wg)
	go process(highChan, lowChan, 2, &wg)

	for _, task := range tasks {
		if task.Priority >= 5 {
			highChan <- task
		} else {
			lowChan <- task
		}
	}
	close(highChan)
	close(lowChan)

	wg.Wait()
}
