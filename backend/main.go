package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Initialize database
	initDB()

	// Routes
	r.GET("/todos", getTodos)
	r.POST("/todos", addTodo)

	// Run the app
	if err := r.Run(":8080"); err != nil {
		fmt.Println(err)
	}

}

// handler to get all todos
func getTodos(c *gin.Context) {
	var todos []Todo
	if err := db.Find(&todos).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, todos)
	}
}

// handler to add a new todo
func addTodo(c *gin.Context) {
	var todo Todo
	c.BindJSON(&todo)

	if err := db.Create(&todo).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}
