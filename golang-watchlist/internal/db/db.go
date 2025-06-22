package db

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB() (*sql.DB, error) {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=UTF8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"))

	const maxRetries = 10
	const retryInterval = 2 * time.Second

	var db *sql.DB
	var err error

	for i := 0; i < maxRetries; i++ {
		db, err = sql.Open("mysql", connStr)
		if err == nil {
			break
		}
		fmt.Printf("Failed to connect to database (attempt %d/%d): %v\n", i+1, maxRetries, err)
		time.Sleep(retryInterval)
	}

	if err != nil {
		return nil, fmt.Errorf("Failed to connect to database after %d attempts: %v\n", maxRetries, err)
	}

	schema := `
	CREATE TABLE IF NOT EXISTS watch_list (
		id SERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		total_episodes INTEGER,
		watched_episodes INTEGER,
		type TEXT,
		status TEXT NOT NULL
	)
	`

	_, err = db.Exec(schema)
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("Failed to create watch_list table: %v", err)
	}

	return db, nil
}
