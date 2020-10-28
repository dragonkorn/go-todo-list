package server

import (
	"fmt"
	"net/http"

	"github.com/dragonkorn/go-todo-list/src/db"

	"github.com/dragonkorn/go-todo-list/src/controllers"
	"github.com/labstack/echo"
)

func getRouter(router *echo.Echo) *echo.Echo {
	router.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<h1>Hello world</h1>")
	})

	router.POST("/todo", func(c echo.Context) error {
		todo := new(db.TodoItemModel)
		if err := c.Bind(todo); err != nil {
			return err
		}
		fmt.Println(todo)
		controllers.TodoCreate(todo)
		return c.String(http.StatusOK, "SUCCESS")
	})

	router.GET("/todos", func(c echo.Context) error {
		result := controllers.TodoQuery()
		fmt.Println(result)
		return c.JSON(http.StatusOK, result)
	})

	router.File("/index", "public/templates/index.html")

	router.GET("/ping", controllers.Ping)

	return router
}
