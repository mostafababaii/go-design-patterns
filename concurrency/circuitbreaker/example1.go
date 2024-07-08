package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type State int

const (
	Closed State = iota
	Open
	HalfOpen
	maxFailures     = 3
	resetTimeout    = 5 * time.Second
	requestTimeout  = 1 * time.Second
	halfOpenTimeout = 2 * time.Second
)

type CircuitBreaker struct {
	state                State
	failures             uint32
	lastFailureTimestamp time.Time
	lastRestTimestamp    time.Time
	requestCh            chan string
	responseCh           chan bool
}

func (cb *CircuitBreaker) monitor() {
	for {
		switch cb.state {
		case Closed:
			go cb.makeRequest()
		case Open:
			select {
			case <-time.After(resetTimeout):
				cb.attemptReset()
			case <-cb.requestCh:
				cb.responseCh <- false
			}
		case HalfOpen:
			cb.attemptReset()
			go cb.makeRequest()
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func (cb *CircuitBreaker) makeRequest() {
	cb.recordRequest()

	select {
	case <-time.After(requestTimeout):
		cb.recordFailure()
		cb.responseCh <- false
	case <-cb.requestCh:
		cb.recordSuccess()
		cb.responseCh <- true
	}
}

func (cb *CircuitBreaker) recordRequest() {
	fmt.Println("Recording request...")
}

func (cb *CircuitBreaker) recordSuccess() {
	fmt.Println("Recording success...")
	if cb.state == HalfOpen {
		cb.state = Closed
		cb.lastRestTimestamp = time.Now()
	}
}

func (cb *CircuitBreaker) recordFailure() {
	fmt.Println("Recording failure...")
	atomic.AddUint32(&cb.failures, 1)

	if cb.failures >= maxFailures && time.Since(cb.lastRestTimestamp) >= resetTimeout {
		cb.state = Open
		cb.lastFailureTimestamp = time.Now()
	}
}

func (cb *CircuitBreaker) attemptReset() {
	if time.Since(cb.lastFailureTimestamp) >= halfOpenTimeout {
		cb.state = HalfOpen
	}
}

func (cb *CircuitBreaker) Execute() bool {
	cb.requestCh <- "GET /example"

	return <-cb.responseCh
}

func NewCircuitBreaker() *CircuitBreaker {
	cb := &CircuitBreaker{
		state:      Closed,
		failures:   0,
		requestCh:  make(chan string),
		responseCh: make(chan bool),
	}

	go cb.monitor()

	return cb
}

func main() {
	cb := NewCircuitBreaker()

	for i := 0; i < 5; i++ {
		result := cb.Execute()
		fmt.Printf("Request result: %v\n", result)
	}

	time.Sleep(5 * time.Second)
}
