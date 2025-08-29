package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	ID  int
	URL string
}

type Result struct {
	JobID    int
	URL      string
	Success  bool
	DataSize int
	Duration time.Duration
	Error    string
}

// Task 23: Build a Worker Pool System
// Your task is to implement a concurrent web scraper using goroutines and channels:

func main() {
	numWorkers := 5
	maxJobs := 20
	rateLimitPerSecond := 3

	fmt.Println("=== Web Scraper Worker Pool Challenge ===")
	fmt.Printf("Workers: %d, Max Jobs: %d, Rate Limit: %d/sec\n\n",
		numWorkers, maxJobs, rateLimitPerSecond)

	var wg sync.WaitGroup

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	jobQueue := make(chan Job, 100)
	resultQueue := make(chan Result, 100)

	startWorkerPool(ctx, numWorkers, jobQueue, resultQueue, rateLimitPerSecond, &wg)

	go jobGenerator(ctx, jobQueue, maxJobs)

	go func() {
		collectResults(ctx, resultQueue, maxJobs)
		cancel()
	}()

	wg.Wait()
}

func startWorkerPool(ctx context.Context, numWorkers int, jobQueue <-chan Job,
	resultQueue chan<- Result, rateLimit int, wg *sync.WaitGroup) {

	rateLimiter := time.NewTicker(time.Second / time.Duration(rateLimit))

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(ctx, i, jobQueue, resultQueue, rateLimiter.C, wg)
	}

	fmt.Println("TODO: Implement startWorkerPool function")
}

func worker(ctx context.Context, id int, jobQueue <-chan Job,
	resultQueue chan<- Result, rateLimiter <-chan time.Time, wg *sync.WaitGroup) {

	defer wg.Done()
	fmt.Printf("TODO: Implement worker %d\n", id)
	// Your code here...

	for {
		select {
		case job, ok := <-jobQueue:
			if !ok {
				fmt.Printf("Worker %d: Job queue closed, shutting down\n", id)
				return
			}

			// TODO: Wait for rate limiter
			result := processJobWithRetry(ctx, id, job, rateLimiter, 3) // max 3 retries
			// TODO: Process the job with retry logic
			// TODO: Send result
			select {
			case resultQueue <- result:
				// Successfully sent result
			case <-ctx.Done():
				fmt.Printf("Worker %d: Context cancelled while sending result\n", id)
				return
			}
		case <-ctx.Done():
			fmt.Printf("Worker %d: Context cancelled, shutting down\n", id)
			return
		}
	}
}

func processJobWithRetry(ctx context.Context, workerID int, job Job,
	rateLimiter <-chan time.Time, maxRetries int) Result {

	var lastError string

	for attempt := 1; attempt <= maxRetries; attempt++ {
		// Wait for rate limiter before making request
		select {
		case <-rateLimiter:
			// Proceed with request
		case <-ctx.Done():
			return Result{
				JobID:   job.ID,
				URL:     job.URL,
				Success: false,
				Error:   "context cancelled",
			}
		}

		fmt.Printf("Worker %d: Processing job %d (attempt %d/%d): %s\n",
			workerID, job.ID, attempt, maxRetries, job.URL)

		// Simulate web scraping
		success, dataSize, duration, errorMsg := simulateWebScrape(job.URL)

		if success {
			return Result{
				JobID:    job.ID,
				URL:      job.URL,
				Success:  true,
				Error:    errorMsg,
				DataSize: dataSize,
				Duration: duration,
			}
		}

		lastError = errorMsg
		fmt.Printf("Worker %d: Job %d failed (attempt %d): %s\n",
			workerID, job.ID, attempt, errorMsg)

		// Wait before retry (exponential backoff)
		if attempt < maxRetries {
			backoffTime := time.Duration(attempt) * 500 * time.Millisecond
			select {
			case <-time.After(backoffTime):
				// Continue to next retry
			case <-ctx.Done():
				return Result{
					JobID:   job.ID,
					URL:     job.URL,
					Success: false,
					Error:   "context cancelled during retry",
				}
			}
		}
	}

	// All retries failed
	return Result{
		JobID:   job.ID,
		URL:     job.URL,
		Success: false,
		Error:   fmt.Sprintf("failed after %d attempts: %s", maxRetries, lastError),
	}
}

func jobGenerator(ctx context.Context, jobQueue chan<- Job, maxJobs int) {
	defer close(jobQueue)

	urls := []string{
		"https://example.com/page1",
		"https://api.service.com/data",
		"https://news.site.com/articles",
		"https://store.shop.com/products",
		"https://blog.tech.com/posts",
	}

	for i := 0; i < maxJobs; i++ {
		job := Job{
			ID:  i + 1,
			URL: urls[rand.Intn(len(urls))] + fmt.Sprintf("/%d", i),
		}

		select {
		case jobQueue <- job:
			fmt.Printf("Generated job %d: %s\n", job.ID, job.URL)
		case <-ctx.Done():
			fmt.Println("Job generator: Context cancelled")
			return
		}

		time.Sleep(5 * time.Second) // Simulate job creation delay
	}
}

func collectResults(ctx context.Context, resultQueue <-chan Result, expectedResults int) {
	var results []Result
	var successCount int
	var totalDuration time.Duration

	fmt.Println("\n=== Collecting Results ===")

	for i := 0; i < expectedResults; i++ {
		select {
		case result := <-resultQueue:
			results = append(results, result)
			if result.Success {
				successCount++
			}
			totalDuration += result.Duration

			status := "✓"
			if !result.Success {
				status = "✗"
			}

			fmt.Printf("%s Job %d: %s (%.2fs)\n",
				status, result.JobID, result.URL, result.Duration.Seconds())

		case <-ctx.Done():
			fmt.Println("Result collector: Context cancelled")
			return
		}
	}

	successRate := float64(successCount) / float64(len(results)) * 100
	avgDuration := totalDuration / time.Duration(len(results))

	fmt.Printf("\n=== Final Statistics ===\n")
	fmt.Printf("Total Jobs: %d\n", len(results))
	fmt.Printf("Success Rate: %.1f%%\n", successRate)
	fmt.Printf("Average Duration: %.2fs\n", avgDuration.Seconds())
}

func simulateWebScrape(url string) (bool, int, time.Duration, string) {
	start := time.Now()

	delay := time.Duration(rand.Intn(2000)+500) * time.Millisecond
	time.Sleep(delay)

	// Random success/failure (80% success rate)
	success := rand.Float64() < 0.8
	dataSize := 0
	errorMsg := ""

	if success {
		dataSize = rand.Intn(5000) + 1000 // 1KB to 6KB
	} else {
		errors := []string{
			"connection timeout",
			"404 not found",
			"rate limited",
			"server error",
		}
		errorMsg = errors[rand.Intn(len(errors))]
	}

	return success, dataSize, time.Since(start), errorMsg
}
