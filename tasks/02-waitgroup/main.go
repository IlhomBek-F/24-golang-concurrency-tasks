package main

import (
	"fmt"
	"sync"
)

// Task 2: WaitGroup Example

// Run multiple goroutines that each print a message, and use a sync.WaitGroup to make sure the main goroutine waits until all of them finish.

func printMessage(msg string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Print(msg)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go printMessage("hello from go routine "+fmt.Sprint(i)+"\n", &wg)
	}

	wg.Wait()

	fmt.Println("All goroutines finished!")
}
