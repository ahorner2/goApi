package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func GetDBConnection() *sql.DB {
    connStr := os.Getenv("DATABASE_URL") 
    if connStr  == "" {
        log.Fatal("DATABASE_URL not set in env")
    }
    
    db, err := sql.Open("postgres", connStr) 
    if err != nil {
        log.Fatalf("Failed to open the database: %v", err)
    }

    return db
}
