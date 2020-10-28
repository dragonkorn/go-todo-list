package controllers

import (
	"github.com/dragonkorn/go-todo-list/src/db"
	"gorm.io/gorm"
)

// TodoQuery - query all todo
func TodoQuery() []db.TodoItemModel {
	var todo []db.TodoItemModel
	db.DB.Find(&todo)
	return todo
}

// TodoCreate return pointer of todo
func TodoCreate(item *db.TodoItemModel) *gorm.DB {
	// todo := db.TodoItemModel{
	// 	Title:       item.title,
	// 	Description: item.description,
	// 	CreatedAt:   time.Now(),
	// 	DueDate:     time.Now(),
	// 	Stared:      false,
	// }
	result := db.DB.Create(&item)
	return result
}
