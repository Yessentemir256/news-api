// database/database.go
package database

import (
	"database/sql"
	"log"

	"github.com/go-reform/reform"
	"github.com/go-reform/reform/dialects/postgresql"
	_ "github.com/lib/pq"
)

func InitDatabase() *reform.DB {
	dsn := "user=username password=password dbname=mydb sslmode=disable"
	sqlDB, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	db := reform.NewDB(sqlDB, postgresql.Dialect, reform.NewPrintfLogger(log.Printf))
	return db
}
