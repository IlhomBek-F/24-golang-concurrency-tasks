package main

import (
	"fmt"
	"time"
)

// Task 1: Hello Goroutine

// Create a program where:

// The main goroutine prints "Hello from main".

// A separate goroutine prints "Hello from goroutine".

// Ensure the program waits long enough so the goroutine actually executes before the program exits.

func sayHello() {
	fmt.Println("Hello from go routine")
}

func main() {

	go sayHello()

	time.Sleep(time.Second * 1)
	fmt.Println("hello from main")
}
