package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// START OMIT
func helloHandler(c *gin.Context) { // HL
	log.Println("in helloHandler")
	c.JSON(http.StatusOK, gin.H{"message": "hello"})
} // HL

func authMiddleware(c *gin.Context) { // HL
	log.Println("start middleware")
	authKey := c.GetHeader("Authorization") // HL
	if authKey != "Bearer token123" {
		c.JSON(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized)) // HL
		c.Abort()                                                                 // HL
		return
	}

	c.Next()
	log.Println("end middleware")
}

func main() {
	r := gin.Default()
	r.Use(authMiddleware) // HL
	r.GET("/hello", helloHandler)
	r.Run(":1234")
}

// END OMIT
