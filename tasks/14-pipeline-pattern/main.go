package main

import (
	"fmt"
	"sync"
)

// 🔹 Task 15: Pipeline Pattern
// Implement a pipeline where data flows through multiple stages, each handled by a separate goroutine.

func square(input <-chan int, squareChan chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range input {
		squareChan <- n * n
	}
	close(squareChan)
}

func double(squareChan <-chan int, resultChan chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for sq := range squareChan {
		resultChan <- sq * 2
	}
	close(resultChan)
}

func main() {
	inputChan := make(chan int)
	squareChan := make(chan int)
	resultChan := make(chan int)

	var wg sync.WaitGroup

	wg.Add(2)
	go square(inputChan, squareChan, &wg)
	go double(squareChan, resultChan, &wg)

	go func() {
		for i := 1; i <= 5; i++ {
			inputChan <- i + 1
		}
		close(inputChan)
	}()

	for result := range resultChan {
		fmt.Println(result)
	}
}
