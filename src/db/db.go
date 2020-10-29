package db

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localPostgres"
	port     = 5432
	user     = "postgres"
	password = "5100"
	dbname   = "go-todo-list"
)

// TodoItemModel - model for todo item
type TodoItemModel struct {
	ID          uint
	Title       string
	Description string
	CreatedAt   time.Time
	DueDate     time.Time
	Done        bool
	Stared      bool
}

// DB - Database
var DB *gorm.DB

// Connect = connect to database
func Connect() *gorm.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})

	// dsn := "host=localPostgres user=postgres password=5100 dbname=go-todo-list port=5432 sslmode=disable TimeZone=Asia/Bangkok"
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db
	autoMigrate()
	// defer db.Close()

	fmt.Println("Successfully connected!")
	return db
}

// auto Migrate
func autoMigrate() {
	DB.AutoMigrate(&TodoItemModel{})
}

// CheckDb for check database type
func CheckDb() {
	fmt.Printf("%T\n", DB)
}
