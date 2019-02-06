package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

// TODOS OMIT
var todos = map[string]*Todo{
	"1": &Todo{ID: "1", Title: "pay phone bills", Status: "active"}, // HL
	"2": &Todo{ID: "2", Title: "pay credit", Status: "completed"},
	"3": &Todo{ID: "3", Title: "homework go language", Status: "active"}, // HL
	"4": &Todo{ID: "4", Title: "buy new shoes", Status: "completed"},
}

// ENDTODOS OMIT

// START OMIT
func getTodosHandler(c *gin.Context) { // HL
	status := c.Query("status") // HL
	items := []*Todo{}

	for _, item := range todos {
		if status != "" {
			if item.Status == status { // HL
				items = append(items, item)
			} // HL
		} else {
			items = append(items, item)
		}
	}

	c.JSON(http.StatusOK, items)
} // HL

func main() {
	r := gin.Default()
	r.GET("/todos", getTodosHandler)
	r.Run(":1234")
}

// END OMIT
