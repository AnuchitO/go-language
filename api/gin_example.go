package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func helloHandler(c *gin.Context) { // HL
	c.JSON(http.StatusOK, gin.H{ // HL
		"message": "hello",
	}) // HL
} // HL

func main() {
	r := gin.Default()            // HL
	r.GET("/hello", helloHandler) // HL
	r.Run()                       // listen and serve on 127.0.0.0:8080  // HL
}
