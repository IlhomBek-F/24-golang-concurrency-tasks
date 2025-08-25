package main

import (
	"fmt"
)

// Task 4: Buffered Channels
// Understand how buffered channels work vs unbuffered channels.

func main() {
	bufferedChan := make(chan string, 2) // buffer size 2

	go func() {
		bufferedChan <- "message 1"
		fmt.Println("Sent message 1")
		bufferedChan <- "message 2"
		fmt.Println("Sent message 2")
		bufferedChan <- "message 3"   // will block here until main receives
		fmt.Println("Sent message 3") // wil not printed until message 3 read
	}()

	fmt.Println("Receiving first message")
	fmt.Println(<-bufferedChan)
	fmt.Println("Receiving second message")
	fmt.Println(<-bufferedChan)
	fmt.Println("Receiving third message")
	fmt.Println(<-bufferedChan)
}
