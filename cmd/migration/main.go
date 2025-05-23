package main

import (
	"fmt"
	"os"

	infrastructure "github.com/victor-ferrer/gprc-sample/infrastructure/db"

	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {

	connStr, err := getDBConfig()
	if err != nil {
		log.Fatalf("Failed to retrieve database configuration: %v", err)
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}

	log.Println("Successfully connected to the database")

	err = infrastructure.RunDatabaseMigrations(db)
	if err != nil {
		log.Fatalf("Failed to run database migrations: %v", err)
	}

	log.Println("Database migrations completed successfully")

}

func getDBConfig() (string, error) {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	if user == "" || password == "" || dbname == "" || host == "" || port == "" {
		return "", fmt.Errorf("missing required environment variables for database configuration")
	}

	return fmt.Sprintf("postgresql://%s:%s/%s?user=%s&password=%s&sslmode=disable", host, port, dbname, user, password), nil
}
