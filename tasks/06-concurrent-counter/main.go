package main

import (
	"fmt"
	"sync"
)

// Task 6: Concurrent Counter with sync.Mutex
// Create a counter that multiple goroutines increment concurrently.
// Use sync.Mutex to prevent race conditions.

func increment(count *int, mt *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	mt.Lock()
	*count++
	mt.Unlock()
}

func main() {
	var count int
	var mt sync.Mutex
	var wg sync.WaitGroup

	for range 1000 {
		wg.Add(5)
		go increment(&count, &mt, &wg)
		go increment(&count, &mt, &wg)
		go increment(&count, &mt, &wg)
		go increment(&count, &mt, &wg)
		go increment(&count, &mt, &wg)
	}

	wg.Wait()

	fmt.Println("Final counter ", +count) // Final counter 5000
}
