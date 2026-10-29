package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	jobs := make(chan int)
	results := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(4)
	completed, sum := 0, 0

	// producer
	go func() {
		defer close(jobs)
		for i := range 20 {
			select {
			case jobs <- i + 1:
			case <-ctx.Done():
				return
			}
		}
	}()

	// workers
	for i := 1; i < 5; i++ {
		go func(ctx context.Context) {
			defer wg.Done()
			for job := range jobs {
				select {
				case <-ctx.Done():
					return
				case <-time.After(300 * time.Millisecond):
					select {
					case <-ctx.Done():
						return
					case results <- job * job:
					}
				}
			}
		}(ctx)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		completed++
		sum += result
	}
	fmt.Printf("completed: %d\nsum: %d\n", completed, sum)
}
