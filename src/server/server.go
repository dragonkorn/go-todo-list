package server

import (
	// "github.com/dragonkorn/go-todo-list/src/db"
	"github.com/labstack/echo"
)

// Init initial server
func Init() {
	e := echo.New()

	// db.Connect()

	e = getRouter(e)

	e.Logger.Fatal(e.Start(":8080"))
}
