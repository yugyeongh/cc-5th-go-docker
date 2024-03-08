package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello world")
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	router.GET("/hello", func(c *gin.Context) {
		name := c.Query("name")
		c.String(http.StatusOK, "Hello, %s!", name)
	})

	router.Run(":8080")
}
