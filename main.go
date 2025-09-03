package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthResponse represents the JSON structure for /health response
type HealthResponse struct {
	Status string `json:"status"`
}

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Define the /health route
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, HealthResponse{Status: "ok"})
	})

	// Run the server on port 8080
	router.Run(":8080")
}
