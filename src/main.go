package main

import (
	"fmt"

	"github.com/dragonkorn/go-todo-list/src/server"
	"github.com/dragonkorn/go-todo-list/src/test"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Hello world from main.go")
	test.Println()
	// fmt.Println("Hello world")
	server.Init()
}
