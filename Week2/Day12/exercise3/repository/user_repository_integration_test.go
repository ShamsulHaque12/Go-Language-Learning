package repository

import (
	"database/sql"
	"errors"
	"os"
	"testing"

	"go-language-learning/Week2/Day12/exercise3/database"
	"go-language-learning/Week2/Day12/exercise3/models"
)

func TestUserRepository_CreateAndGetByID_Integration(t *testing.T) {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		t.Skip("Skipping integration test: DATABASE_URL is not set")
	}

	db, err := database.ConnectDB(databaseURL)
	if err != nil {
		t.Fatalf("failed to connect to test database: %v", err)
	}
	defer db.Close()

	if err := database.CreateUserTable(db); err != nil {
		t.Fatalf("failed to create users table: %v", err)
	}

	repo := NewUserRepository(db)

	user := &models.User{
		Name:     "Integration Test User",
		Email:    "integration_test_day16@example.com",
		Age:      27,
		IsActive: true,
	}

	// Clean up this test's email before and after the test.
	_, err = db.Exec(
		"DELETE FROM users WHERE email = $1",
		user.Email,
	)
	if err != nil {
		t.Fatalf("failed to clean up previous test data: %v", err)
	}

	defer func() {
		_, err := db.Exec(
			"DELETE FROM users WHERE email = $1",
			user.Email,
		)
		if err != nil {
			t.Errorf("failed to clean up test user: %v", err)
		}
	}()

	// Create a user in the real PostgreSQL database.
	createdUser, err := repo.Create(user)
	if err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	if createdUser.ID <= 0 {
		t.Fatalf("expected generated user ID > 0, got %d", createdUser.ID)
	}

	// Retrieve the created user from PostgreSQL.
	foundUser, err := repo.GetByID(createdUser.ID)
	if err != nil {
		t.Fatalf("failed to get user by ID: %v", err)
	}

	if foundUser.Name != user.Name {
		t.Errorf("expected name %q, got %q", user.Name, foundUser.Name)
	}

	if foundUser.Email != user.Email {
		t.Errorf("expected email %q, got %q", user.Email, foundUser.Email)
	}

	if foundUser.Age != user.Age {
		t.Errorf("expected age %d, got %d", user.Age, foundUser.Age)
	}

	if foundUser.IsActive != user.IsActive {
		t.Errorf("expected IsActive %t, got %t", user.IsActive, foundUser.IsActive)
	}

	// Sanity check: a nonexistent ID should return an error.
	_, err = repo.GetByID(-1)
	if err == nil {
		t.Error("expected an error for nonexistent user ID, got nil")
	} else if err == sql.ErrNoRows {
		// Expected database/sql error.
	} else {
		t.Logf("GetByID returned an error for nonexistent ID: %v", err)
	}
}
func TestUserRepository_GetAll_Integration(t *testing.T) {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		t.Skip("Skipping integration test: DATABASE_URL is not set")
	}

	db, err := database.ConnectDB(databaseURL)
	if err != nil {
		t.Fatalf("failed to connect to test database: %v", err)
	}
	defer db.Close()

	if err := database.CreateUserTable(db); err != nil {
		t.Fatalf("failed to create users table: %v", err)
	}

	repo := NewUserRepository(db)

	user := &models.User{
		Name:     "GetAll Integration User",
		Email:    "getall_integration_day16@example.com",
		Age:      27,
		IsActive: true,
	}

	_, err = db.Exec("DELETE FROM users WHERE email = $1", user.Email)
	if err != nil {
		t.Fatalf("failed to clean up previous test data: %v", err)
	}
	defer db.Exec("DELETE FROM users WHERE email = $1", user.Email)

	_, err = repo.Create(user)
	if err != nil {
		t.Fatalf("failed to create test user: %v", err)
	}

	users, err := repo.GetAll()
	if err != nil {
		t.Fatalf("GetAll failed: %v", err)
	}

	found := false
	for _, u := range users {
		if u.Email == user.Email {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("expected GetAll to include user %q", user.Email)
	}
}

func TestUserRepository_Delete_Integration(t *testing.T) {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		t.Skip("Skipping integration test: DATABASE_URL is not set")
	}

	db, err := database.ConnectDB(databaseURL)
	if err != nil {
		t.Fatalf("failed to connect to test database: %v", err)
	}
	defer db.Close()

	if err := database.CreateUserTable(db); err != nil {
		t.Fatalf("failed to create users table: %v", err)
	}

	repo := NewUserRepository(db)

	user := &models.User{
		Name:     "Delete Integration User",
		Email:    "delete_integration_day16@example.com",
		Age:      27,
		IsActive: true,
	}

	_, err = db.Exec("DELETE FROM users WHERE email = $1", user.Email)
	if err != nil {
		t.Fatalf("failed to clean up previous test data: %v", err)
	}

	createdUser, err := repo.Create(user)
	if err != nil {
		t.Fatalf("failed to create test user: %v", err)
	}

	if err := repo.Delete(createdUser.ID); err != nil {
		t.Fatalf("Delete failed: %v", err)
	}

	_, err = repo.GetByID(createdUser.ID)
	if !errors.Is(err, sql.ErrNoRows) {
		t.Errorf("expected sql.ErrNoRows after deletion, got %v", err)
	}
}
func TestUserRepository_Update_Integration(t *testing.T) {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		t.Skip("Skipping integration test: DATABASE_URL is not set")
	}

	db, err := database.ConnectDB(databaseURL)
	if err != nil {
		t.Fatalf("failed to connect to test database: %v", err)
	}
	defer db.Close()

	if err := database.CreateUserTable(db); err != nil {
		t.Fatalf("failed to create users table: %v", err)
	}

	repo := NewUserRepository(db)

	user := &models.User{
		Name:     "Before Update",
		Email:    "update_integration_day16@example.com",
		Age:      25,
		IsActive: true,
	}

	_, err = db.Exec("DELETE FROM users WHERE email = $1", user.Email)
	if err != nil {
		t.Fatalf("failed to clean up previous test data: %v", err)
	}
	defer db.Exec("DELETE FROM users WHERE email = $1", user.Email)

	createdUser, err := repo.Create(user)
	if err != nil {
		t.Fatalf("failed to create test user: %v", err)
	}

	createdUser.Name = "After Update"
	createdUser.Age = 28
	createdUser.IsActive = false

	if err := repo.Update(createdUser); err != nil {
		t.Fatalf("Update failed: %v", err)
	}

	updatedUser, err := repo.GetByID(createdUser.ID)
	if err != nil {
		t.Fatalf("failed to retrieve updated user: %v", err)
	}

	if updatedUser.Name != "After Update" {
		t.Errorf("expected name %q, got %q",
			"After Update", updatedUser.Name)
	}

	if updatedUser.Age != 28 {
		t.Errorf("expected age 28, got %d", updatedUser.Age)
	}

	if updatedUser.IsActive {
		t.Error("expected IsActive to be false after update")
	}
}
