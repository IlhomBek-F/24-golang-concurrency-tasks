package main

import (
	"fmt"
	"sync"
	"time"
)

// Task 7: Fan-Out Pattern
// Distribute work across multiple goroutines to process tasks concurrently.

func worker(taskChan <-chan int, resultChan chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range taskChan {
		fmt.Print("Proccesing task " + fmt.Sprint(task) + "\n")
		resultChan <- "Completed task " + fmt.Sprint(task)
		time.Sleep(time.Second * 2)
	}
}

func main() {
	taskChan := make(chan int)
	resultChan := make(chan string)

	var wg sync.WaitGroup

	for range 4 {
		wg.Add(1)
		go worker(taskChan, resultChan, &wg)
	}

	go func() {
		for i := range 10 {
			taskChan <- i
		}
		close(taskChan)
	}()

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for result := range resultChan {
		fmt.Println(result)
	}
}
