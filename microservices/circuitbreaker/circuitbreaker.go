package circuitbreaker

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
)

type CircuitBreaker struct {
	state State

	// MaxFailures is the maximum number of failures allowed before tripping the circuit breaker.
	MaxFailures int

	// Timeout is the time to wait before transitioning to the half-open state.
	Timeout time.Duration

	// SuccessThreshold is the number of successful requests required in the half-open state to transition back to the closed state.
	SuccessThreshold int

	// FailureCount is the number of failed requests.
	FailureCount uint32

	// LastFailureTime is the time of the last failed request.
	LastFailureTime time.Time

	// SuccessCount is the number of successful requests in the half-open state.
	SuccessCount uint32
}

func NewCircuitBreaker(maxFailures int, timeout time.Duration, successThreshold int) *CircuitBreaker {
	return &CircuitBreaker{
		state:            Closed,
		MaxFailures:      maxFailures,
		Timeout:          timeout,
		SuccessThreshold: successThreshold,
	}
}

func (cb *CircuitBreaker) Execute(operation func() error) error {
	switch cb.state {
	case Closed:
		err := operation()
		if err != nil {
			atomic.AddUint32(&cb.FailureCount, 1)
			cb.LastFailureTime = time.Now()
			if atomic.LoadUint32(&cb.FailureCount) >= uint32(cb.MaxFailures) {
				cb.state = Open
			}
			return err
		}
		atomic.StoreUint32(&cb.FailureCount, 0)
		return nil
	case Open:
		if time.Since(cb.LastFailureTime) >= cb.Timeout {
			cb.state = HalfOpen
			atomic.StoreUint32(&cb.SuccessCount, 0)
		}
		return fmt.Errorf("circuit breaker is open")
	case HalfOpen:
		err := operation()
		if err != nil {
			atomic.AddUint32(&cb.FailureCount, 1)
			cb.LastFailureTime = time.Now()
			if atomic.LoadUint32(&cb.FailureCount) >= uint32(cb.MaxFailures) {
				cb.state = Open
			}
			return err
		}
		atomic.AddUint32(&cb.SuccessCount, 1)
		if atomic.LoadUint32(&cb.SuccessCount) >= uint32(cb.SuccessThreshold) {
			cb.state = Closed
		}
		return nil
	}
	return nil
}
