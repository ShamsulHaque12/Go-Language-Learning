package database

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestCreateUserTable_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}
	defer db.Close()

	mock.ExpectExec("CREATE TABLE IF NOT EXISTS users").
		WillReturnResult(sqlmock.NewResult(0, 0))

	err = CreateUserTable(db)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet SQL expectations: %v", err)
	}
}

func TestCreateUserTable_DatabaseError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}
	defer db.Close()

	mock.ExpectExec("CREATE TABLE IF NOT EXISTS users").
		WillReturnError(assertionError{})

	err = CreateUserTable(db)
	if err == nil {
		t.Fatal("expected database error, got nil")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet SQL expectations: %v", err)
	}
}

type assertionError struct{}

func (assertionError) Error() string {
	return "simulated database error"
}

func TestConnectDB_InvalidURL(t *testing.T) {
	db, err := ConnectDB("://invalid-url")

	if err == nil {
		if db != nil {
			db.Close()
		}
		t.Fatal("expected connection error, got nil")
	}

	if db != nil {
		db.Close()
		t.Error("expected nil database connection on error")
	}
}
