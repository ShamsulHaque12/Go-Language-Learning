package service

import (
	"errors"
	"testing"

	"go-language-learning/Week2/Day12/exercise3/models"
	"go-language-learning/Week2/Day12/exercise3/repository"
)

type mockUserRepository struct {
	getAllFunc  func() ([]models.User, error)
	getByIDFunc func(id int) (*models.User, error)
	createFunc  func(user *models.User) (*models.User, error)
	updateFunc  func(user *models.User) error
	deleteFunc  func(id int) error
}

func (m *mockUserRepository) GetAll() ([]models.User, error) {
	return m.getAllFunc()
}

func (m *mockUserRepository) GetByID(id int) (*models.User, error) {
	return m.getByIDFunc(id)
}

func (m *mockUserRepository) Create(user *models.User) (*models.User, error) {
	return m.createFunc(user)
}

func (m *mockUserRepository) Update(user *models.User) error {
	return m.updateFunc(user)
}

func (m *mockUserRepository) Delete(id int) error {
	return m.deleteFunc(id)
}

var _ repository.UserRepository = (*mockUserRepository)(nil)

func TestGetAllUsers(t *testing.T) {
	expected := []models.User{
		{ID: 1, Name: "Sujon", Email: "sujon@example.com", Age: 27},
	}

	repo := &mockUserRepository{
		getAllFunc: func() ([]models.User, error) {
			return expected, nil
		},
	}

	svc := NewUserService(repo)

	got, err := svc.GetAllUsers()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(got) != len(expected) {
		t.Fatalf("expected %d users, got %d", len(expected), len(got))
	}

	if got[0].ID != expected[0].ID {
		t.Errorf("expected ID %d, got %d", expected[0].ID, got[0].ID)
	}
}

func TestGetUserByID_NotFound(t *testing.T) {
	repo := &mockUserRepository{
		getByIDFunc: func(id int) (*models.User, error) {
			return nil, errors.New("user not found")
		},
	}

	svc := NewUserService(repo)

	got, err := svc.GetUserByID(999)
	if err == nil {
		t.Fatal("expected an error, got nil")
	}

	if got != nil {
		t.Errorf("expected nil user, got %+v", got)
	}
}
func TestCreateUser(t *testing.T) {
	isActive := true

	req := &models.CreateUserRequest{
		Name:     "Sujon",
		Email:    "sujon@example.com",
		Age:      27,
		IsActive: &isActive,
	}

	repo := &mockUserRepository{
		createFunc: func(user *models.User) (*models.User, error) {
			user.ID = 1
			return user, nil
		},
	}

	svc := NewUserService(repo)

	got, err := svc.CreateUser(req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if got.ID != 1 {
		t.Errorf("expected ID 1, got %d", got.ID)
	}

	if got.Name != req.Name {
		t.Errorf("expected name %q, got %q", req.Name, got.Name)
	}

	if got.Email != req.Email {
		t.Errorf("expected email %q, got %q", req.Email, got.Email)
	}

	if got.Age != req.Age {
		t.Errorf("expected age %d, got %d", req.Age, got.Age)
	}

	if !got.IsActive {
		t.Error("expected IsActive to be true")
	}
}
func TestUpdateUser_Success(t *testing.T) {
	isActive := true

	req := &models.UpdateUserRequest{
		Name:     "Updated Sujon",
		Email:    "updated@example.com",
		Age:      28,
		IsActive: &isActive,
	}

	repo := &mockUserRepository{
		updateFunc: func(user *models.User) error {
			if user.ID != 10 {
				t.Errorf("expected ID 10, got %d", user.ID)
			}
			if user.Name != req.Name {
				t.Errorf("expected name %q, got %q", req.Name, user.Name)
			}
			return nil
		},
	}

	svc := NewUserService(repo)

	got, err := svc.UpdateUser(10, req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if got.ID != 10 {
		t.Errorf("expected ID 10, got %d", got.ID)
	}
	if got.Name != req.Name {
		t.Errorf("expected name %q, got %q", req.Name, got.Name)
	}
}

func TestUpdateUser_RepositoryError(t *testing.T) {
	isActive := true
	expectedErr := errors.New("database update failed")

	req := &models.UpdateUserRequest{
		Name:     "Sujon",
		Email:    "sujon@example.com",
		Age:      27,
		IsActive: &isActive,
	}

	repo := &mockUserRepository{
		updateFunc: func(user *models.User) error {
			return expectedErr
		},
	}

	svc := NewUserService(repo)

	got, err := svc.UpdateUser(10, req)

	if !errors.Is(err, expectedErr) {
		t.Errorf("expected error %v, got %v", expectedErr, err)
	}

	if got != nil {
		t.Errorf("expected nil user, got %+v", got)
	}
}

func TestPatchUser_Success(t *testing.T) {
	existing := &models.User{
		ID:       5,
		Name:     "Old Name",
		Email:    "old@example.com",
		Age:      25,
		IsActive: true,
	}

	newName := "Updated Name"

	repo := &mockUserRepository{
		getByIDFunc: func(id int) (*models.User, error) {
			if id != 5 {
				t.Errorf("expected ID 5, got %d", id)
			}
			return existing, nil
		},
		updateFunc: func(user *models.User) error {
			return nil
		},
	}

	svc := NewUserService(repo)

	req := &models.PatchUserRequest{
		Name: &newName,
	}

	got, err := svc.PatchUser(5, req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if got.Name != newName {
		t.Errorf("expected name %q, got %q", newName, got.Name)
	}

	if got.Email != "old@example.com" {
		t.Errorf("email should remain unchanged, got %q", got.Email)
	}

	if got.Age != 25 {
		t.Errorf("age should remain unchanged, got %d", got.Age)
	}
}

func TestPatchUser_GetByIDError(t *testing.T) {
	expectedErr := errors.New("user not found")

	repo := &mockUserRepository{
		getByIDFunc: func(id int) (*models.User, error) {
			return nil, expectedErr
		},
	}

	svc := NewUserService(repo)

	newName := "Updated Name"
	req := &models.PatchUserRequest{
		Name: &newName,
	}

	got, err := svc.PatchUser(5, req)

	if !errors.Is(err, expectedErr) {
		t.Errorf("expected error %v, got %v", expectedErr, err)
	}

	if got != nil {
		t.Errorf("expected nil user, got %+v", got)
	}
}

func TestPatchUser_UpdateError(t *testing.T) {
	expectedErr := errors.New("database update failed")

	existing := &models.User{
		ID:       5,
		Name:     "Old Name",
		Email:    "old@example.com",
		Age:      25,
		IsActive: true,
	}

	newName := "Updated Name"

	repo := &mockUserRepository{
		getByIDFunc: func(id int) (*models.User, error) {
			return existing, nil
		},
		updateFunc: func(user *models.User) error {
			return expectedErr
		},
	}

	svc := NewUserService(repo)

	req := &models.PatchUserRequest{
		Name: &newName,
	}

	got, err := svc.PatchUser(5, req)

	if !errors.Is(err, expectedErr) {
		t.Errorf("expected error %v, got %v", expectedErr, err)
	}

	if got != nil {
		t.Errorf("expected nil user, got %+v", got)
	}
}

func TestDeleteUser_Success(t *testing.T) {
	deletedID := 0

	repo := &mockUserRepository{
		deleteFunc: func(id int) error {
			deletedID = id
			return nil
		},
	}

	svc := NewUserService(repo)

	err := svc.DeleteUser(10)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if deletedID != 10 {
		t.Errorf("expected ID 10, got %d", deletedID)
	}
}

func TestDeleteUser_RepositoryError(t *testing.T) {
	expectedErr := errors.New("database delete failed")

	repo := &mockUserRepository{
		deleteFunc: func(id int) error {
			return expectedErr
		},
	}

	svc := NewUserService(repo)

	err := svc.DeleteUser(10)

	if !errors.Is(err, expectedErr) {
		t.Errorf("expected error %v, got %v", expectedErr, err)
	}
}
