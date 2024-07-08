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
	router.GET("/orders", GetOrders)
	router.Run(":8081")
}

func GetOrders(c *gin.Context) {
	userID := c.Query("user_id")

	err := cb.Execute(func() error {
		// Make a request to the InventoryService to get the product inventory.
		resp, err := http.Get("http://localhost:8082/inventory")
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		// Check the response status code.
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("failed to get inventory: %s", resp.Status)
		}

		// Process the response body.
		// ...

		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the order data.
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Order data for user ID %s", userID)})
}
