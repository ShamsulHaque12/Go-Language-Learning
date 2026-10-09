package service

import (
	"go-language-learning/Week2/Day12/exercise2/models"
	"go-language-learning/Week2/Day12/exercise2/repository"
)

type UserService interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(id int) (*models.User, error)
	CreateUser(req *models.CreateUserRequest) (*models.User, error)
	UpdateUser(id int, req *models.UpdateUserRequest) (*models.User, error)
	PatchUser(id int, req *models.PatchUserRequest) (*models.User, error)
	DeleteUser(id int) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAll()
}

func (s *userService) GetUserByID(id int) (*models.User, error) {
	return s.repo.GetByID(id)
}

func (s *userService) CreateUser(req *models.CreateUserRequest) (*models.User, error) {
	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Age:      req.Age,
		IsActive: *req.IsActive,
	}
	return s.repo.Create(user)
}

func (s *userService) UpdateUser(id int, req *models.UpdateUserRequest) (*models.User, error) {
	user := &models.User{
		ID:       id,
		Name:     req.Name,
		Email:    req.Email,
		Age:      req.Age,
		IsActive: *req.IsActive,
	}
	if err := s.repo.Update(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) PatchUser(id int, req *models.PatchUserRequest) (*models.User, error) {
	existing, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if req.Name != nil {
		existing.Name = *req.Name
	}
	if req.Email != nil {
		existing.Email = *req.Email
	}
	if req.Age != nil {
		existing.Age = *req.Age
	}
	if req.IsActive != nil {
		existing.IsActive = *req.IsActive
	}

	if err := s.repo.Update(existing); err != nil {
		return nil, err
	}
	return existing, nil
}

func (s *userService) DeleteUser(id int) error {
	return s.repo.Delete(id)
}
