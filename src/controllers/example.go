package controllers

import (
	"net/http"

	"github.com/dragonkorn/go-todo-list/src/db"
	"github.com/labstack/echo"
)

// Ping - Pong
func Ping(c echo.Context) error {
	db.CheckDb()
	return c.String(http.StatusOK, "Pong")
}
