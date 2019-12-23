package repository

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"tutorial/todo-list/database"
	"tutorial/todo-list/models"
)

func CreateTodo(todo models.Todo) *gorm.DB {
	db := database.GetDB()
	return db.Create(&todo)
}

func UpdateTodo(todo models.Todo) *gorm.DB {
	db := database.GetDB()
	return db.Save(&todo)
}

func GetTodoById(id uuid.UUID) *gorm.DB {
	var todo = models.Todo{}
	db := database.GetDB()
	return db.First(&todo, id)
}

func GetTodos() *[]models.Todo {
	var todo = []models.Todo{}
	db := database.GetDB()
	db.Find(&todo)
	return &todo
}

func RemoveTodosById(id uuid.UUID) *gorm.DB {
	var todo = models.Todo{}
	db := database.GetDB()
	return db.Delete(&todo, id)
}
