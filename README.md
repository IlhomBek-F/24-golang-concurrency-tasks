# 🚀 50 Go Concurrency Tasks

This repository contains 50 hands-on tasks to learn and practice **Golang concurrency** with goroutines, channels, and sync utilities.

---
## 📋 Task List
  
### Basics
1. **Hello Goroutine** – Run a simple goroutine that prints "Hello from goroutine".
2. **WaitGroup Example** – Launch multiple goroutines and wait for them to finish.
3. **Channel Basics** – Send and receive data using channels.
4. **Buffered Channel** – Demonstrate buffered vs unbuffered channels.
5. **Select Statement** – Use `select` to listen to multiple channels.

### Practical Examples
6. **Concurrent Counter** – Safely increment a counter with `sync.Mutex`.
7. **Fan-out** – Split tasks across multiple goroutines.
8. **Worker Pool** – Limit concurrency with a fixed number of workers.
9. **Rate limiting** – Implement a simple rate limiter so that goroutines don’t process tasks faster than a given rate.
10. **Cancellation with context** – Cancellation with context.Context
11. **Cancellation with context timeout** – Cancellation with context.Context time out
12. **Rate-Limited Workers with Timeout** – Automatically stop all workers after a timeout using context.WithTimeout
13. **Worker Pool + Result Collection + Timeout + Rate Limiting** – Combine multiple concurrency concepts into one task
---
