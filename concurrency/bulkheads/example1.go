package main

import (
	"fmt"
	"sync"
	"time"
)

// simulateRequest simulates a request to a service.
func simulateRequest(name string, delay time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(delay)
	fmt.Printf("Service %s responded in %v\n", name, delay)
}

func main() {
	var wg sync.WaitGroup
	// Define the number of concurrent requests each service can handle.
	serviceALimit := 2
	serviceBLimit := 2

	// Create buffered channels to act as Bulkheads.
	serviceARequests := make(chan struct{}, serviceALimit)
	serviceBRequests := make(chan struct{}, serviceBLimit)

	// Simulate requests to Service A.
	for i := 0; i < serviceALimit; i++ {
		wg.Add(1)
		serviceARequests <- struct{}{} // Acquire a slot in the bulkhead.

		go func(i int) {
			simulateRequest(fmt.Sprintf("A%d", i), 2*time.Second, &wg)
			<-serviceARequests // Release the slot.
		}(i)
	}

	// Simulate requests to Service B.
	for i := 0; i < serviceBLimit; i++ {
		wg.Add(1)
		serviceBRequests <- struct{}{} // Acquire a slot in the bulkhead.

		go func(i int) {
			simulateRequest(fmt.Sprintf("B%d", i), 1*time.Second, &wg)
			<-serviceBRequests // Release the slot.
		}(i)
	}

	// Wait for all goroutines to finish.
	wg.Wait()
	fmt.Println("All services have responded.")
}
