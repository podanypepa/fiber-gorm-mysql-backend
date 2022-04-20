package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dbhost     = ""
	dbport     = ""
	dbuser     = ""
	dbpassword = ""
	dbname     = ""
)

// DB is connected MySQL DB
var DB *gorm.DB

func init() {
	dbhost = os.Getenv("DBHOST")
	dbport = os.Getenv("DBPORT")
	dbuser = os.Getenv("DBUSER")
	dbpassword = os.Getenv("DBPASSWORD")
	dbname = os.Getenv("DBNAME")
}

// Connect to MySQL server
func Connect() {
	fmt.Println(dbhost)
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbuser,
		dbpassword,
		dbhost,
		dbport,
		dbname,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	DB = db
}

// GetConfig for debuging
func GetConfig() (string, string, string, string, string) {
	return dbhost, dbport, dbuser, dbpassword, dbname
}
