package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

// TASK 24 Build a Log Aggregator
// Build a system that collects log messages from multiple sources, processes them, and writes them to an output file.

type LogMessage struct {
	Source  string
	Level   string // INFO, WARN, ERROR
	Message string
}

type ProcessedLog struct {
	Timestamp time.Time
	Source    string
	Level     string
	Message   string
}

func webSource(messageCh chan<- LogMessage, wg *sync.WaitGroup) {
	defer wg.Done()
	for range 5 {
		messageCh <- LogMessage{Source: "Web", Level: "INFO", Message: "info log"}
		time.Sleep(time.Second * 5)
	}
}

func dbSource(messages chan<- LogMessage, wg *sync.WaitGroup) {
	defer wg.Done()
	for range 3 {
		messages <- LogMessage{Source: "db", Level: "ERROR", Message: "error log"}
		time.Sleep(time.Second * 3)
	}
}

func authSource(messages chan<- LogMessage, wg *sync.WaitGroup) {
	defer wg.Done()
	for range 4 {
		messages <- LogMessage{Source: "db", Level: "WARN", Message: "warn log"}
		time.Sleep(time.Second * 2)
	}
}

func addLog(messages <-chan LogMessage, writer *bufio.Writer, done chan<- struct{}) {
	defer close(done)
	for message := range messages {
		formatted := fmt.Sprintf("[%s] [%s] [%s]: [%s]\n", time.Now().Format("2006-01-02 15:04:05"), message.Level, message.Source, message.Message)
		_, err := writer.WriteString(formatted)
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println(formatted)
	}
}

func main() {
	messages := make(chan LogMessage)
	done := make(chan struct{})
	file, err := os.Create("output.log")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	var wg sync.WaitGroup

	go addLog(messages, writer, done)

	wg.Add(3)
	go webSource(messages, &wg)
	go dbSource(messages, &wg)
	go authSource(messages, &wg)

	wg.Wait()
	close(messages)

	<-done
	fmt.Println("All logs processed successfully!")

}
