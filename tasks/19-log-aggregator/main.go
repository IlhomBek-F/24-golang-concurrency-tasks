package main

import (
	"fmt"
	"sync"
)

// System Task 19 – Log Aggregator
// Imagine you’re building a central log aggregator service that collects logs from multiple microservices.

func service(srv string, logChan chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range 5 {
		logChan <- i
	}
}

func collector(logChan <-chan int, done chan<- bool) {
	for log := range logChan {
		fmt.Println(log)
	}

	done <- true
}

func main() {
	logChan := make(chan int)
	done := make(chan bool)

	go collector(logChan, done)

	var wg sync.WaitGroup

	services := []string{"serviceA", "serviceB", "serviceC"}

	for _, srv := range services {
		wg.Add(1)
		go service(srv, logChan, &wg)
	}

	go func() {
		wg.Wait()
		close(logChan)
	}()

	<-done

	println("All log printed")
}
