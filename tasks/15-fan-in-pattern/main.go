package main

import (
	"fmt"
	"sync"
)

// Task 15: Fan-in Pattern
// Merge multiple channels into a single channel (fan-in).
// This is useful when you have multiple producers and you want to read all results through one single channel.

func fanIn(inputChan <-chan int, mergeChan chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for input := range inputChan {
		mergeChan <- input
	}
}

func main() {
	evenChan := make(chan int)
	oddChan := make(chan int)
	mergedChan := make(chan int)

	var wg sync.WaitGroup

	wg.Add(2)
	go fanIn(evenChan, mergedChan, &wg)
	go fanIn(oddChan, mergedChan, &wg)

	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 0 {
				evenChan <- i
			} else {
				oddChan <- i
			}
		}

		close(evenChan)
		close(oddChan)
	}()

	go func() {
		wg.Wait()
		close(mergedChan)
	}()

	for result := range mergedChan {
		fmt.Println(result)
	}
}
