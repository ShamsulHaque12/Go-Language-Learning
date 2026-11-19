package database

import (
	"os"
	"testing"
)

func TestConnectDB_Integration(t *testing.T) {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		t.Skip("Skipping integration test: DATABASE_URL is not set")
	}

	db, err := ConnectDB(databaseURL)
	if err != nil {
		t.Fatalf("failed to connect to test database: %v", err)
	}
	defer db.Close()

	var dbName, userName string
	err = db.QueryRow(
		"SELECT current_database(), current_user",
	).Scan(&dbName, &userName)
	if err != nil {
		t.Fatalf("failed to query database identity: %v", err)
	}

	if dbName != "go_learning_test" {
		t.Errorf("expected database go_learning_test, got %s", dbName)
	}

	if userName != "go_test_user" {
		t.Errorf("expected user go_test_user, got %s", userName)
	}
}

func TestCreateUserTable_Integration(t *testing.T) {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		t.Skip("Skipping integration test: DATABASE_URL is not set")
	}

	db, err := ConnectDB(databaseURL)
	if err != nil {
		t.Fatalf("failed to connect to test database: %v", err)
	}
	defer db.Close()

	err = CreateUserTable(db)
	if err != nil {
		t.Fatalf("failed to create users table: %v", err)
	}

	var tableExists bool
	err = db.QueryRow(`
        SELECT EXISTS (
            SELECT 1
            FROM information_schema.tables
            WHERE table_schema = 'public'
              AND table_name = 'users'
        )
    `).Scan(&tableExists)

	if err != nil {
		t.Fatalf("failed to verify users table: %v", err)
	}

	if !tableExists {
		t.Fatal("expected users table to exist")
	}
}
