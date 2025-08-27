package main

import (
	"context"
	"fmt"
	"time"
)

// Task 20 Concurrent Fetch with Fallback
// Query 3 sources concurrently, take first response, cancel the others.

func fetch(result chan<- string, ctx context.Context, ind string) {
	select {
	case result <- ind:
	case <-ctx.Done():
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	ctxTimeout, cancelTimeout := context.WithTimeout(context.Background(), time.Second*3)

	defer cancelTimeout()

	result := make(chan string)
	result2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 3)
		fetch(result, ctx, "second")
	}()

	go func() {
		time.Sleep(time.Second * 2)
		fetch(result2, ctx, "second")
	}()

	select {
	case first := <-result:
		fmt.Println(first)
		cancel()
	case second := <-result2:
		fmt.Println(second)
		cancel()
	case <-ctxTimeout.Done():
		fmt.Println("time out")
		cancel()
	}
}
