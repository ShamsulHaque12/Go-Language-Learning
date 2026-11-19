package repository

import (
	"database/sql"
	"errors"
	"testing"

	"go-language-learning/Week2/Day12/exercise3/models"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetAll_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows(
		[]string{"id", "name", "email", "age", "is_active"},
	).
		AddRow(1, "Sujon", "sujon@example.com", 27, true).
		AddRow(2, "Rahim", "rahim@example.com", 25, false)

	mock.ExpectQuery(
		`SELECT id, name, email, age, is_active FROM users ORDER BY id ASC;`,
	).WillReturnRows(rows)

	repo := NewUserRepository(db)

	users, err := repo.GetAll()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(users) != 2 {
		t.Fatalf("expected 2 users, got %d", len(users))
	}

	expected := []models.User{
		{ID: 1, Name: "Sujon", Email: "sujon@example.com", Age: 27, IsActive: true},
		{ID: 2, Name: "Rahim", Email: "rahim@example.com", Age: 25, IsActive: false},
	}

	for i, user := range users {
		if user != expected[i] {
			t.Errorf("user[%d]: expected %+v, got %+v", i, expected[i], user)
		}
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet SQL expectations: %v", err)
	}
}

func TestGetAll_QueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}
	defer db.Close()

	expectedErr := errors.New("database query failed")

	mock.ExpectQuery(
		`SELECT id, name, email, age, is_active FROM users ORDER BY id ASC;`,
	).WillReturnError(expectedErr)

	repo := NewUserRepository(db)

	users, err := repo.GetAll()

	if err == nil {
		t.Fatal("expected an error, got nil")
	}

	if users != nil {
		t.Errorf("expected nil users, got %+v", users)
	}

	if !errors.Is(err, expectedErr) {
		t.Errorf("expected wrapped error %v, got %v", expectedErr, err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet SQL expectations: %v", err)
	}
}

func TestGetByID_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows(
		[]string{"id", "name", "email", "age", "is_active"},
	).AddRow(1, "Sujon", "sujon@example.com", 27, true)

	mock.ExpectQuery(
		`SELECT id, name, email, age, is_active FROM users WHERE id = \$1;`,
	).WithArgs(1).WillReturnRows(rows)

	repo := NewUserRepository(db)

	user, err := repo.GetByID(1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	expected := models.User{
		ID: 1, Name: "Sujon", Email: "sujon@example.com",
		Age: 27, IsActive: true,
	}

	if *user != expected {
		t.Errorf("expected %+v, got %+v", expected, *user)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet SQL expectations: %v", err)
	}
}

func TestGetByID_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}
	defer db.Close()

	mock.ExpectQuery(
		`SELECT id, name, email, age, is_active FROM users WHERE id = \$1;`,
	).WithArgs(99).WillReturnRows(
		sqlmock.NewRows(
			[]string{"id", "name", "email", "age", "is_active"},
		),
	)

	repo := NewUserRepository(db)

	user, err := repo.GetByID(99)

	if !errors.Is(err, sql.ErrNoRows) {
		t.Errorf("expected sql.ErrNoRows, got %v", err)
	}

	if user != nil {
		t.Errorf("expected nil user, got %+v", user)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet SQL expectations: %v", err)
	}
}

func TestCreate_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}
	defer db.Close()

	user := &models.User{
		Name:     "Sujon",
		Email:    "sujon@example.com",
		Age:      27,
		IsActive: true,
	}

	mock.ExpectQuery(
		`INSERT INTO users \(name, email, age, is_active\)
        VALUES \(\$1, \$2, \$3, \$4\)
        RETURNING id;`,
	).
		WithArgs(user.Name, user.Email, user.Age, user.IsActive).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}).AddRow(10),
		)

	repo := NewUserRepository(db)

	createdUser, err := repo.Create(user)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if createdUser.ID != 10 {
		t.Errorf("expected ID 10, got %d", createdUser.ID)
	}

	if createdUser.Name != "Sujon" {
		t.Errorf("expected name Sujon, got %s", createdUser.Name)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet SQL expectations: %v", err)
	}
}

func TestCreate_DatabaseError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}
	defer db.Close()

	user := &models.User{
		Name:     "Sujon",
		Email:    "sujon@example.com",
		Age:      27,
		IsActive: true,
	}

	expectedErr := errors.New("database insert failed")

	mock.ExpectQuery(
		`INSERT INTO users \(name, email, age, is_active\)
        VALUES \(\$1, \$2, \$3, \$4\)
        RETURNING id;`,
	).
		WithArgs(user.Name, user.Email, user.Age, user.IsActive).
		WillReturnError(expectedErr)

	repo := NewUserRepository(db)

	createdUser, err := repo.Create(user)

	if !errors.Is(err, expectedErr) {
		t.Errorf("expected error %v, got %v", expectedErr, err)
	}

	if createdUser != nil {
		t.Errorf("expected nil user, got %+v", createdUser)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet SQL expectations: %v", err)
	}
}

func TestUpdate_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}
	defer db.Close()

	user := &models.User{
		ID:       5,
		Name:     "Sujon",
		Email:    "[EMAIL_ADDRESS]",
		Age:      27,
		IsActive: true,
	}

	mock.ExpectExec(
		`UPDATE users SET name = \$1, email = \$2, age = \$3, is_active = \$4 WHERE id = \$5;`,
	).
		WithArgs(user.Name, user.Email, user.Age, user.IsActive, user.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	repo := NewUserRepository(db)

	err = repo.Update(user)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet SQL expectations: %v", err)
	}
}

func TestUpdate_DatabaseError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}
	defer db.Close()

	user := &models.User{
		ID:       5,
		Name:     "Sujon",
		Email:    "[EMAIL_ADDRESS]",
		Age:      27,
		IsActive: true,
	}

	expectedErr := errors.New("database update failed")

	mock.ExpectExec(
		`UPDATE users SET name = \$1, email = \$2, age = \$3, is_active = \$4 WHERE id = \$5;`,
	).
		WithArgs(user.Name, user.Email, user.Age, user.IsActive, user.ID).
		WillReturnError(expectedErr)

	repo := NewUserRepository(db)

	err = repo.Update(user)

	if !errors.Is(err, expectedErr) {
		t.Errorf("expected error %v, got %v", expectedErr, err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet SQL expectations: %v", err)
	}
}

func TestDelete_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}
	defer db.Close()

	mock.ExpectExec(
		`DELETE FROM users WHERE id = \$1;`,
	).
		WithArgs(5).
		WillReturnResult(sqlmock.NewResult(0, 1))

	repo := NewUserRepository(db)

	err = repo.Delete(5)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet SQL expectations: %v", err)
	}
}

func TestDelete_DatabaseError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}
	defer db.Close()

	expectedErr := errors.New("database delete failed")

	mock.ExpectExec(
		`DELETE FROM users WHERE id = \$1;`,
	).
		WithArgs(5).
		WillReturnError(expectedErr)

	repo := NewUserRepository(db)

	err = repo.Delete(5)

	if !errors.Is(err, expectedErr) {
		t.Errorf("expected error %v, got %v", expectedErr, err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unmet SQL expectations: %v", err)
	}
}
