package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// START OMIT
type Todo struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var todos = map[string]*Todo{
	"1": &Todo{ID: "1", Title: "pay phone bills", Status: "active"},
}

func getTodosHandler(c *gin.Context) { // HL
	items := []*Todo{}
	for _, item := range todos {
		items = append(items, item)
	}
	c.JSON(http.StatusOK, items) // HL
} // HL

func main() {
	r := gin.Default()
	r.GET("/todos", getTodosHandler)
	r.Run()
}

// END OMIT
