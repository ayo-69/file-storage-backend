package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initalize Gin router
	router := gin.Default()

	//Defile a route
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello from Gin!")
	})

	//Get PORT
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	//Run the server
	router.Run(":" + port)
}
