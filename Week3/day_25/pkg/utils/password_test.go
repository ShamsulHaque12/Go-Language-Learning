package utils

import (
	"testing"
)

func TestHashPasswordAndCheck(t *testing.T) {
	password := "SecretP@ssw0rd!"

	hashed, err := HashPassword(password)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(hashed) == 0 {
		t.Fatalf("expected hashed password to be non-empty")
	}

	if !CheckPassword(hashed, password) {
		t.Errorf("expected CheckPassword to return true for matching password")
	}

	if CheckPassword(hashed, "WrongPassword") {
		t.Errorf("expected CheckPassword to return false for incorrect password")
	}
}

func TestHashEmptyPassword(t *testing.T) {
	_, err := HashPassword("")
	if err == nil {
		t.Errorf("expected error for empty password, got nil")
	}
}
