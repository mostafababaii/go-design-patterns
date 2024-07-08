
# Distributed System with Microservices

## Overview
This is an example of a distributed system with microservices built using Go, the Gin framework, and the Circuit Breaker pattern. The system consists of three microservices:

•  **UserService**: Responsible for managing user data.

•  **OrderService**: Responsible for managing orders.

•  **InventoryService**: Responsible for managing product inventory.


Each microservice is implemented as a separate Go program using the Gin framework, and they communicate with each other over HTTP using REST APIs.

## Circuit Breaker
A separate `CircuitBreaker` package provides a common interface for implementing the pattern. It defines a `CircuitBreaker` struct that maintains the state of the circuit breaker, the maximum number of failures allowed before tripping, the timeout before transitioning to the half-open state, and the success threshold required to transition back to the closed state.

## Running the System
To run the system, start each microservice in a separate terminal window:

•  **UserService**: `go run ./userService`

•  **OrderService**: `go run ./orderService`

•  **InventoryService**: `go run ./inventoryService`


Each microservice listens on a different port:

•  **UserService**: `8080`

•  **OrderService**: `8081`

•  **InventoryService**: `8082`


## Testing the System
Use tools like `curl` or `httpie` to send HTTP requests to the microservices. Example requests:

•  Get user data: `curl http://localhost:8080/users/1`

•  Get orders for a user: `curl http://localhost:8081/orders?user_id=1`

•  Get product inventory: `curl http://localhost:8082/inventory`


To test the circuit breaker, simulate a failure in one of the microservices by stopping it or introducing a delay in the response.

```go
func GetInventory(c *gin.Context) {
	err := cb.Execute(func() error {
		// Simulate a delay by sleeping for 5 seconds.
		time.Sleep(5 * time.Second)

		// Return the inventory data.
		c.JSON(http.StatusOK, gin.H{"message": "Inventory data"})
		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
```

## Conclusion
This example demonstrates building a distributed system with microservices using Go, the Gin framework, and the Circuit Breaker pattern. The Circuit Breaker pattern improves the resilience of the system and prevents cascading failures when a microservice fails.
