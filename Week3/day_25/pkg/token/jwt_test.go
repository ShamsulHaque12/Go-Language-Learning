package token

import (
	"testing"
	"time"
)

func TestGenerateAndValidateToken(t *testing.T) {
	secret := "test-secret-key-12345"
	userID := uint(42)
	email := "user@example.com"
	role := "admin"

	tokenStr, err := GenerateToken(userID, email, role, secret, time.Hour)
	if err != nil {
		t.Fatalf("failed to generate token: %v", err)
	}

	claims, err := ValidateToken(tokenStr, secret)
	if err != nil {
		t.Fatalf("failed to validate token: %v", err)
	}

	if claims.UserID != userID {
		t.Errorf("expected userID %d, got %d", userID, claims.UserID)
	}
	if claims.Email != email {
		t.Errorf("expected email %s, got %s", email, claims.Email)
	}
	if claims.Role != role {
		t.Errorf("expected role %s, got %s", role, claims.Role)
	}
}

func TestExpiredToken(t *testing.T) {
	secret := "test-secret-key-12345"
	tokenStr, err := GenerateToken(1, "test@test.com", "user", secret, -time.Minute)
	if err != nil {
		t.Fatalf("failed to generate expired token: %v", err)
	}

	_, err = ValidateToken(tokenStr, secret)
	if err == nil {
		t.Errorf("expected validation to fail for expired token")
	}
}

func TestInvalidSecret(t *testing.T) {
	tokenStr, err := GenerateToken(1, "test@test.com", "user", "secret-1", time.Hour)
	if err != nil {
		t.Fatalf("failed to generate token: %v", err)
	}

	_, err = ValidateToken(tokenStr, "wrong-secret")
	if err == nil {
		t.Errorf("expected validation to fail with wrong secret")
	}
}
