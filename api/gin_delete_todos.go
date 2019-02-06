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
	todos[id] = &t // HL

	c.JSON(http.StatusCreated, "created todo.") // HL
} // HL

func getTodoByIdHandler(c *gin.Context) {
	id := c.Param("id") // HL

	t, ok := todos[id] // HL
	if !ok {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	c.JSON(http.StatusOK, t) // HL
}
func updateTodosHandler(c *gin.Context) {
	id := c.Param("id")
	t := todos[id]                              // HL
	if err := c.ShouldBindJSON(t); err != nil { // HL
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} // HL

	c.JSON(http.StatusOK, t)
}

// START OMIT
func deleteTodosHandler(c *gin.Context) {
	id := c.Param("id")
	delete(todos, id)
	c.JSON(http.StatusOK, "deleted todo.")
}

func main() {
	r := gin.Default()
	r.GET("/todos", getTodosHandler)
	r.GET("/todos/:id", getTodoByIdHandler)
	r.POST("/todos", createTodosHandler)
	r.PUT("todos/:id", updateTodosHandler)
	r.DELETE("todos/:id", deleteTodosHandler) //HL

	r.Run(":1234")
}

// END OMIT
