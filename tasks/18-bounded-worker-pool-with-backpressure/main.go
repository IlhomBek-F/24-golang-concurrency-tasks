package main

import (
	"fmt"
	"sync"
	"time"
)

// 📝 Task 18: Bounded Worker Pool with Backpressure
// Backpressure happens when the producer (sender) generates values faster than the consumer (worker/receiver) can handle.

func worker(tasks <-chan int, wg *sync.WaitGroup, workerInd int) {
	defer wg.Done()
	for task := range tasks {
		fmt.Printf("Worker %d processing task %d\n", workerInd, task)
		time.Sleep(time.Second * 2)
	}
}

func main() {
	tasks := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 123}
	taskChan := make(chan int, len(tasks))

	var wg sync.WaitGroup

	wg.Add(2)
	go worker(taskChan, &wg, 1)
	go worker(taskChan, &wg, 2)

	for _, task := range tasks {
		fmt.Println("sending")
		taskChan <- task
		fmt.Println("sent")
	}

	close(taskChan)
	wg.Wait()
}
