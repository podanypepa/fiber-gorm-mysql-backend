package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/podanypepa/fiber-gorm-mysql-backend/pkg/database"
	"github.com/podanypepa/fiber-gorm-mysql-backend/pkg/rest"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	database.Connect()

	port := os.Getenv("PORT")
	rest.Create().Listen(fmt.Sprintf(":%s", port))
}
