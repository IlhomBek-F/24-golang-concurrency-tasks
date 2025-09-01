# 🚀 24 Go Concurrency Tasks

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
14. **Pipeline Pattern** – Implement a pipeline where data flows through multiple stages, each handled by a separate goroutine.
15. **Fan-in Pattern** – Merge multiple channels into a single channel (fan-in).
16. **Worker Pool with Priorities** – Implement a worker pool that processes tasks, but each task has a priority.
17. **Rate-Limited Worker Pool with Context Cancellation** – Combine worker pools, rate limiting, and early cancellation using context.Context.
18. **Bounded Worker Pool with Backpressure** – Backpressure happens when the producer (sender) generates values faster than the consumer (worker/receiver) can handle.
19. **Log Aggregator** - Imagine you’re building a central log aggregator service that collects logs from multiple microservices
20. **Concurrent Fetch with Fallback** - Query 3 sources concurrently, take first response, cancel the others
21. **Task Multi-Service Shutdown** - You are building a system with 3 independent services running concurrently. Each service processes tasks in a loop
22. **Timeout + Heartbeat Monitor** - Timeout + Heartbeat Monitor
23. **Build a Worker Pool System** - Implement a concurrent web scraper using goroutines and channels
24. **Build a Log Aggregator** - Build a system that collects log messages from multiple sources, processes them, and writes them to an output file
---
