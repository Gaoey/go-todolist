package handler

import (
	"net/http"
	"tutorial/todo-list/models"
	"tutorial/todo-list/repository"

	uuid "github.com/satori/go.uuid"

	"github.com/labstack/echo"
)

func CreateTodo(c echo.Context) error {
	todo := models.Todo{}
	if err := c.Bind(&todo); err != nil {
		return err
	}
	db := repository.CreateTodo(todo)
	if db.Error != nil {
		return c.JSON(http.StatusInternalServerError, "cannot create todo")
	}
	return c.JSON(http.StatusOK, db.Value)
}

func GetTodos(c echo.Context) error {
	todo := repository.GetTodos()
	return c.JSON(http.StatusOK, &todo)
}

func GetTodoById(c echo.Context) error {
	id := c.Param("id")

	u, err := uuid.FromString(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "id is not uuid")
	}

	db := repository.GetTodoById(u)
	if db.Error != nil {
		return c.JSON(http.StatusInternalServerError, "cannot get todo by id")
	}
	return c.JSON(http.StatusOK, db.Value)
}

// func UpdateTodo(c echo.Context) error {
// 	// id, _ := strconv.Atoi(c.Param("id"))
// 	return c.JSON(http.StatusOK, nil)
// }

// func RemoveTodo(c echo.Context) error {
// 	// id, _ := strconv.Atoi(c.Param("id"))
// 	return c.JSON(http.StatusOK, nil)
// }
