package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// START OMIT
func helloHandler(c *gin.Context) { // HL
	log.Println("in helloHandler")
	c.JSON(http.StatusOK, gin.H{ // HL
		"message": "hello",
	}) // HL
} // HL

func main() {
	r := gin.Default()

	r.Use(func(c *gin.Context) { // HL
		log.Println("start middleware")
		c.Next() // HL
		log.Println("end middleware")
	}) // HL

	r.GET("/hello", helloHandler)
	r.Run(":1234")
}

// END OMIT
