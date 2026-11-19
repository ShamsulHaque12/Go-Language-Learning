package service

import (
	"context"

	"day_25/internal/domain"
	"day_25/internal/repository"
)

type UserService interface {
	GetProfile(ctx context.Context, userID uint) (*domain.UserResponse, error)
	GetAllUsers(ctx context.Context, page, pageSize int) ([]domain.UserResponse, int64, error)
	UpdateProfile(ctx context.Context, userID uint, req *domain.UpdateProfileRequest) (*domain.UserResponse, error)
	TransferBalance(ctx context.Context, fromID, toID uint, amount int64) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) GetProfile(ctx context.Context, userID uint) (*domain.UserResponse, error) {
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	res := user.ToResponse()
	return &res, nil
}

func (s *userService) GetAllUsers(ctx context.Context, page, pageSize int) ([]domain.UserResponse, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize
	users, total, err := s.userRepo.GetAll(ctx, offset, pageSize)
	if err != nil {
		return nil, 0, err
	}

	responses := make([]domain.UserResponse, len(users))
	for i, u := range users {
		responses[i] = u.ToResponse()
	}

	return responses, total, nil
}

func (s *userService) UpdateProfile(ctx context.Context, userID uint, req *domain.UpdateProfileRequest) (*domain.UserResponse, error) {
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	user.Name = req.Name
	if err := s.userRepo.Update(ctx, user); err != nil {
		return nil, err
	}

	res := user.ToResponse()
	return &res, nil
}

func (s *userService) TransferBalance(ctx context.Context, fromID, toID uint, amount int64) error {
	return s.userRepo.TransferBalance(ctx, fromID, toID, amount)
}
