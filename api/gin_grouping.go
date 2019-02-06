package main

import "github.com/gin-gonic/gin"

// START OMIT
func main() {
	router := gin.Default()

	// Simple group: v1
	v1 := router.Group("/v1") // HL
	v1.POST("/login", loginEndpoint)
	v1.POST("/submit", submitEndpoint)
	v1.POST("/read", readEndpoint)

	// Simple group: v2
	v2 := router.Group("/v2") // HL
	v2.POST("/login", loginEndpoint)
	v2.POST("/submit", submitEndpoint)
	v2.POST("/read", readEndpoint)

	router.Run(":1234")
}

// END OMIT
