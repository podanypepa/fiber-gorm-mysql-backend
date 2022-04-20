package main

import (
	"log"

	"github.com/podanypepa/fiber-gorm-mysql-backend/pkg/database"
	"github.com/podanypepa/fiber-gorm-mysql-backend/pkg/model/todo"
	"github.com/podanypepa/fiber-gorm-mysql-backend/pkg/rest"
)

func main() {
	database.Connect()
	if err := database.DB.AutoMigrate(&todo.TODO{}); err != nil {
		log.Fatal(err)
	}

	if err := rest.Create().Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}
