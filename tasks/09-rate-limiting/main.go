package main

import (
	"fmt"
	"sync"
	"time"
)

// 🔹 Task 10: Rate Limiting with Goroutines
// Implement a simple rate limiter so that goroutines don’t process tasks faster than a given rate (e.g., 1 task per second).

func worker(task int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Processing task %d at %s\n", task, time.Now().Format("15:04:05"))
}

func main() {
	tasks := []int{1, 2, 3, 4, 5}

	var wg sync.WaitGroup
	ticker := time.NewTicker(time.Second * 3)
	defer ticker.Stop()

	for task := range tasks {
		<-ticker.C
		wg.Add(1)
		go worker(task, &wg)
	}

	wg.Wait()
}
