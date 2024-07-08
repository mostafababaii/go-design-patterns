package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mostafababaii/go-design-patterns/microservices/circuitbreaker"
)

var cb = circuitbreaker.NewCircuitBreaker(3, 5*time.Second, 2)

func main() {
	router := gin.Default()
	router.GET("/users/:id", GetUser)
	router.Run(":8080")
}

func GetUser(c *gin.Context) {
	id := c.Param("id")

	err := cb.Execute(func() error {
		// Make a request to the OrderService to get the user's orders.
		resp, err := http.Get(fmt.Sprintf("http://localhost:8081/orders?user_id=%s", id))
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		// Check the response status code.
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("failed to get orders: %s", resp.Status)
		}

		// Process the response body.
		// ...

		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the user data.
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("User data for ID %s", id)})
}
