package main

import (
	"fmt"
	"sync"
	"time"
)

// taskProcessor process the task
func taskProcessor(task int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Second)
	fmt.Printf("Processing task %d\n", task)
}

func main() {
	var wg sync.WaitGroup
	bulkheadLimit := 3
	taskCount := 10

	// Create the buffered channel to limit the number of active bulkheads
	bulkheadCh := make(chan struct{}, bulkheadLimit)

	// Add tasks to the wait group and process them in the separate bulkheads
	for i := 0; i < taskCount; i++ {
		wg.Add(1)

		// Acquire a bulkhead slot
		bulkheadCh <- struct{}{}

		go func(task int) {
			defer func() {
				// Release the bulkhead slot
				<-bulkheadCh
			}()

			// Process the task in a separate bulkhead
			taskProcessor(task, &wg)
		}(i)
	}

	wg.Wait()
}
