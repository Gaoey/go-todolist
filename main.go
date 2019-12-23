package main

import (
	"tutorial/todo-list/database"
	"tutorial/todo-list/handler"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	database.Init()
	defer database.CloseDB()

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/todos", handler.CreateTodo)
	e.GET("/todos", handler.GetTodos)
	e.GET("/todos/:id", handler.GetTodoById)
	// e.PUT("/todos/:id", handler.UpdateTodo)
	// e.DELETE("/todos/:id", handler.RemoveTodo)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
