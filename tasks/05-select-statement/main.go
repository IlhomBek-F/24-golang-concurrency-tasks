package main

import (
	"fmt"
	"sync"
	"time"
)

// Task 5: Select Statement
// Learn how to use select to handle multiple channels and timeouts.

func sendMessage(msg string, msgChan chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	msgChan <- msg
}

func main() {
	chan1 := make(chan string)
	chan2 := make(chan string)

	var wg sync.WaitGroup

	wg.Add(2)
	go sendMessage("chan 1", chan1, &wg)
	go sendMessage("chan 2", chan2, &wg)

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case msg1 := <-chan1:
		fmt.Println("Received from " + fmt.Sprint(msg1))
	case msg2 := <-chan2:
		fmt.Println("Received from " + fmt.Sprint(msg2))
	case <-done:
		fmt.Println("All message sent")
	default:
		<-time.After(time.Second * 3)
		fmt.Println("time out")
	}
}
