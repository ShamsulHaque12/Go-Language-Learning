package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

func ConnectDB(url string) (*sql.DB, error) {
	return connectDB(url, sql.Open)
}

func connectDB(
	url string,
	openDB func(string, string) (*sql.DB, error),
) (*sql.DB, error) {
	db, err := openDB("postgres", url)
	if err != nil {
		return nil, fmt.Errorf("database connection failed: %w", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	fmt.Println("PostgreSQL connected successfully with connection pool!")
	return db, nil
}

// Connect is an alias for ConnectDB.
func Connect(url string) (*sql.DB, error) {
	return ConnectDB(url)
}

func CreateUserTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(100) NOT NULL UNIQUE,
		age INT,
		is_active BOOLEAN DEFAULT TRUE
	);`

	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create users table: %w", err)
	}

	fmt.Println("Users table created successfully!")
	return nil
}

// CreateUsersTable is an alias for CreateUserTable.
func CreateUsersTable(db *sql.DB) error {
	return CreateUserTable(db)
}
