package db

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB is a struct that represents the database
type DB struct{}

// NewDB is a function that returns a new DB
func NewDB() *DB {
	return &DB{}
}

// ConnectDB is a method that connects to the database
func (db *DB) ConnectDB() *gorm.DB {
	fmt.Println("ConnectDB is being called")

	if os.Getenv("GO_ENV") == "dev" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalln("Error loading .env file:", err)
		}
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=%s",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
		url.QueryEscape("Local"))

	// DB Retry Logic
	var conn *gorm.DB
	var err error
	retries := 5
	for i := 0; i < retries; i++ {
		conn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			fmt.Println("Connection was successful!")
			return conn
		}
		fmt.Printf("Retrying to connect to database (%d/%d)...\n", i+1, retries)
		time.Sleep(2 * time.Second)
	}

	log.Fatalln("Failed to connect to database after retries:", err)
	return nil
}
