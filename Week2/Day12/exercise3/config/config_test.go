package config

import (
	"os"
	"testing"
)

func TestLoadConfig_WithEnvironmentVariables(t *testing.T) {
	t.Setenv("DATABASE_URL", "postgres://test:test@localhost:5432/testdb?sslmode=disable")
	t.Setenv("PORT", "9090")

	cfg := LoadConfig()

	if cfg.DatabaseURL != "postgres://test:test@localhost:5432/testdb?sslmode=disable" {
		t.Errorf("unexpected DATABASE_URL: %s", cfg.DatabaseURL)
	}

	if cfg.Port != "9090" {
		t.Errorf("expected port 9090, got %s", cfg.Port)
	}
}

func TestLoadConfig_DefaultPort(t *testing.T) {
	t.Setenv("DATABASE_URL", "postgres://test:test@localhost:5432/testdb?sslmode=disable")
	t.Setenv("PORT", "")

	cfg := LoadConfig()

	if cfg.Port != "8080" {
		t.Errorf("expected default port 8080, got %s", cfg.Port)
	}
}

func TestLoadConfig_DatabaseURLFromEnvironment(t *testing.T) {
	expectedURL := "postgres://test:test@localhost:5432/testdb?sslmode=disable"
	t.Setenv("DATABASE_URL", expectedURL)
	t.Setenv("PORT", "8080")

	cfg := LoadConfig()

	if cfg.DatabaseURL != expectedURL {
		t.Errorf("expected DATABASE_URL %s, got %s", expectedURL, cfg.DatabaseURL)
	}

	if os.Getenv("DATABASE_URL") != expectedURL {
		t.Error("DATABASE_URL environment variable was not set as expected")
	}
}
