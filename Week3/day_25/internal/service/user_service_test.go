package service_test

import (
	"context"
	"errors"
	"testing"

	"day_25/internal/domain"
	"day_25/internal/service"
)

func TestUserService_Operations(t *testing.T) {
	repo := newMockUserRepository()
	userService := service.NewUserService(repo)
	ctx := context.Background()

	// Seed users
	u1 := &domain.User{Name: "Bob", Email: "bob@example.com", Password: "hash", Role: "user", Balance: 1000}
	u2 := &domain.User{Name: "Charlie", Email: "charlie@example.com", Password: "hash", Role: "user", Balance: 500}
	_ = repo.Create(ctx, u1)
	_ = repo.Create(ctx, u2)

	// 1. GetProfile
	profile, err := userService.GetProfile(ctx, u1.ID)
	if err != nil {
		t.Fatalf("failed to get profile: %v", err)
	}
	if profile.Name != "Bob" || profile.Balance != 1000 {
		t.Errorf("unexpected profile data: %+v", profile)
	}

	// 2. UpdateProfile
	updated, err := userService.UpdateProfile(ctx, u1.ID, &domain.UpdateProfileRequest{Name: "Bobby"})
	if err != nil {
		t.Fatalf("failed to update profile: %v", err)
	}
	if updated.Name != "Bobby" {
		t.Errorf("expected updated name Bobby, got %s", updated.Name)
	}

	// 3. Transfer Balance
	err = userService.TransferBalance(ctx, u1.ID, u2.ID, 400)
	if err != nil {
		t.Fatalf("transfer failed: %v", err)
	}

	p1, _ := userService.GetProfile(ctx, u1.ID)
	p2, _ := userService.GetProfile(ctx, u2.ID)

	if p1.Balance != 600 {
		t.Errorf("expected sender balance 600, got %d", p1.Balance)
	}
	if p2.Balance != 900 {
		t.Errorf("expected receiver balance 900, got %d", p2.Balance)
	}

	// 4. Overdraft transfer should fail
	err = userService.TransferBalance(ctx, u1.ID, u2.ID, 10000)
	if !errors.Is(err, domain.ErrInsufficientBalance) {
		t.Errorf("expected ErrInsufficientBalance, got %v", err)
	}
}
