package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func OpenConn() *sql.DB {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5432"
	}

	user := os.Getenv("DB_USER")
	if user == "" {
		user = "postgres"
	}

	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		password = "password"
	}

	dbname := os.Getenv("DB_NAME")
	if dbname == "" {
		dbname = "youfood"
	}

	connInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	driver := os.Getenv("DB_DRIVER")
	if driver == "" {
		driver = "postgres"
	}

	fmt.Printf("Connecting to %s database at %s:%s with user=%s dbname=%s\n", driver, host, port, user, dbname)

	db, err := sql.Open(driver, connInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to db!")

	return db
}

func CloseConn(db *sql.DB) {
	defer db.Close()
}
