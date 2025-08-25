package main

import (
	"fmt"
)

// Task 3: Channel Basics
// Learn how to send and receive values using a channel.

func main() {
	msgChan := make(chan string)

	go func() {
		msgChan <- "Sending message to channel..."
		msgChan <- "Received: Hello from goroutine!..."
		close(msgChan)
	}()

	for msg := range msgChan {
		fmt.Println(msg)
	}
}
