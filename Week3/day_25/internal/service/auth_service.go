package service

import (
	"context"
	"errors"
	"time"

	"day_25/internal/domain"
	"day_25/internal/repository"
	"day_25/pkg/token"
	"day_25/pkg/utils"
)

type AuthService interface {
	Register(ctx context.Context, req *domain.RegisterRequest) (*domain.AuthResponse, error)
	Login(ctx context.Context, req *domain.LoginRequest) (*domain.AuthResponse, error)
}

type authService struct {
	userRepo     repository.UserRepository
	jwtSecret    string
	jwtExpiresIn time.Duration
}

func NewAuthService(userRepo repository.UserRepository, jwtSecret string, jwtExpiresIn time.Duration) AuthService {
	return &authService{
		userRepo:     userRepo,
		jwtSecret:    jwtSecret,
		jwtExpiresIn: jwtExpiresIn,
	}
}

func (s *authService) Register(ctx context.Context, req *domain.RegisterRequest) (*domain.AuthResponse, error) {
	// Check if user already exists
	existingUser, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err == nil && existingUser != nil {
		return nil, domain.ErrEmailAlreadyExists
	}
	if err != nil && !errors.Is(err, domain.ErrUserNotFound) {
		return nil, err
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	newUser := &domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
		Role:     "user",
		Balance:  req.InitialBalance,
	}

	if err := s.userRepo.Create(ctx, newUser); err != nil {
		return nil, err
	}

	// Generate JWT
	jwtToken, err := token.GenerateToken(newUser.ID, newUser.Email, newUser.Role, s.jwtSecret, s.jwtExpiresIn)
	if err != nil {
		return nil, err
	}

	return &domain.AuthResponse{
		Token: jwtToken,
		User:  newUser.ToResponse(),
	}, nil
}

func (s *authService) Login(ctx context.Context, req *domain.LoginRequest) (*domain.AuthResponse, error) {
	user, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return nil, domain.ErrInvalidCredentials
		}
		return nil, err
	}

	// Verify password
	if !utils.CheckPassword(user.Password, req.Password) {
		return nil, domain.ErrInvalidCredentials
	}

	// Generate JWT
	jwtToken, err := token.GenerateToken(user.ID, user.Email, user.Role, s.jwtSecret, s.jwtExpiresIn)
	if err != nil {
		return nil, err
	}

	return &domain.AuthResponse{
		Token: jwtToken,
		User:  user.ToResponse(),
	}, nil
}
