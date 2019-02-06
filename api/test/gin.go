package main

import "github.com/gin-gonic/gin"

func setupRouter() *gin.Engine { // HL
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "hi.")
	})
	return r
} // HL

func main() {
	r := setupRouter() // HL
	r.Run(":1234")
}
