package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/qor/validations"
	"os"
)

func Connect() (db *gorm.DB, err error) {
	server := os.Getenv("DB_SERVER")
	database := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	port := os.Getenv("DB_PORT")
	conn, err := gorm.Open(
		"postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", server, port, dbUser, database, password),
	)
	validations.RegisterCallbacks(conn)
	return conn, err
}
