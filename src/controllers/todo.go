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

// // TodoGetByStar - query all todo with started equal to true
// func TodoGetByStar() []db.TodoItemModel {
// 	var todo []db.TodoItemModel
// 	db.DB.where("Stared = true").First(&todo)
// 	return todo
// }

// // TodoGetByID - query todo by Id
// func TodoGetByID(Id) db.TodoItemModel {
// 	var todo db.TodoItemModel
// 	db.DB.First(&todo, Id)
// 	return todo
// }

// TodoCreate return pointer of todo
func TodoCreate(item *db.TodoItemModel) *gorm.DB {
	result := db.DB.Create(&item)
	return result
}

// TodoEdit - edit todo by item
func TodoEdit(item *db.TodoItemModel) db.TodoItemModel {
	var todo db.TodoItemModel
	db.DB.First(&todo, item.ID)

	todo.Title = item.Title
	todo.Description = item.Description
	todo.Done = item.Done
	todo.Stared = item.Stared
	db.DB.Save(&todo)

	return todo
}

// TodoRemove - remove todo item
func TodoRemove(ID string) string {
	db.DB.Delete(&db.TodoItemModel{}, ID)
	return "SUCCESS"
}
