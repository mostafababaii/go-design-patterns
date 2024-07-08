package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mostafababaii/go-design-patterns/microservices/circuitbreaker"
)

var cb = circuitbreaker.NewCircuitBreaker(3, 5*time.Second, 2)

func main() {
	router := gin.Default()
	router.GET("/inventory", GetInventory)
	router.Run(":8082")
}

func GetInventory(c *gin.Context) {
	err := cb.Execute(func() error {
		// Simulate a failure by returning an error.
		if rand.Intn(10) == 0 {
			return fmt.Errorf("failed to get inventory")
		}

		// Return the inventory data.
		c.JSON(http.StatusOK, gin.H{"message": "Inventory data"})
		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
