package database

import (
	"database/sql"
	"log"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"

	"github.com/Yessentemir256/news-api/config"
)

func InitDatabase() (*reform.DB, *sql.DB) {
	dsn := config.GetDatabaseDSN()

	// Open *sql.DB
	sqlDB, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Configure connection pool if needed (example)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)

	// Initialize reform.DB
	db := reform.NewDB(sqlDB, postgresql.Dialect, reform.NewPrintfLogger(log.Printf))

	return db, sqlDB
}
