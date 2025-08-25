package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func worker(url string, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	res, err := http.Get(url)

	if err != nil {
		results <- "Failed fetching " + url
	}

	results <- "Fetched " + url
	res.Body.Close()
}

func main() {
	urls := []string{"https://example.com", "https://golang.org", "https://github.com", "https://example.com", "https://golang.org", "https://github.com"}

	resultChan := make(chan string)

	var wg sync.WaitGroup

	for _, url := range urls {
		fmt.Println("Fetching " + url)
		time.Sleep(time.Second * 1)
		wg.Add(1)
		go worker(url, resultChan, &wg)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for res := range resultChan {
		fmt.Println(res)
	}
}
