package service_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"day_25/internal/domain"
	"day_25/internal/service"
)

func TestAuthService_RegisterAndLogin(t *testing.T) {
	repo := newMockUserRepository()
	jwtSecret := "test-secret"
	authService := service.NewAuthService(repo, jwtSecret, time.Hour)
	ctx := context.Background()

	// 1. Register user
	regReq := &domain.RegisterRequest{
		Name:           "Alice Doe",
		Email:          "alice@example.com",
		Password:       "Password123!",
		InitialBalance: 5000,
	}

	authRes, err := authService.Register(ctx, regReq)
	if err != nil {
		t.Fatalf("unexpected registration error: %v", err)
	}

	if authRes.Token == "" {
		t.Errorf("expected non-empty auth token")
	}
	if authRes.User.Email != "alice@example.com" {
		t.Errorf("expected email alice@example.com, got %s", authRes.User.Email)
	}
	if authRes.User.Balance != 5000 {
		t.Errorf("expected balance 5000, got %d", authRes.User.Balance)
	}

	// 2. Duplicate registration should fail
	_, err = authService.Register(ctx, regReq)
	if !errors.Is(err, domain.ErrEmailAlreadyExists) {
		t.Errorf("expected ErrEmailAlreadyExists, got %v", err)
	}

	// 3. Login with correct credentials
	loginReq := &domain.LoginRequest{
		Email:    "alice@example.com",
		Password: "Password123!",
	}
	loginRes, err := authService.Login(ctx, loginReq)
	if err != nil {
		t.Fatalf("unexpected login error: %v", err)
	}
	if loginRes.Token == "" {
		t.Errorf("expected login token")
	}

	// 4. Login with invalid password
	invalidLoginReq := &domain.LoginRequest{
		Email:    "alice@example.com",
		Password: "WrongPassword",
	}
	_, err = authService.Login(ctx, invalidLoginReq)
	if !errors.Is(err, domain.ErrInvalidCredentials) {
		t.Errorf("expected ErrInvalidCredentials, got %v", err)
	}
}
