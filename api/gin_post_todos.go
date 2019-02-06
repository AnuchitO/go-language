package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

// START OMIT
func createTodosHandler(c *gin.Context) { // HL
	t := Todo{}
	if err := c.ShouldBindJSON(&t); err != nil { // HL
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} // HL

	i := len(todos)
	i++
	id := strconv.Itoa(i)
	t.ID = id
	todos[id] = &t

	c.JSON(http.StatusCreated, "created todo.") // HL
} // HL

func main() {
	r := gin.Default()
	r.GET("/todos", getTodosHandler)
	r.POST("/todos", createTodosHandler) // HL

	r.Run(":1234") // HL
}

// END OMIT
