package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUsers(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	recorder := httptest.NewRecorder()

	getUsers(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d",
			http.StatusOK, recorder.Code)
	}

	contentType := recorder.Header().Get("Content-Type")

	if contentType != "application/json" {
		t.Errorf("Expected application/json, got %s", contentType)
	}

	var users []User
	err := json.Unmarshal(recorder.Body.Bytes(), &users)
	if err != nil {
		t.Fatalf("Failed to decode response JSON: %v", err)
	}

	if len(users) != 2 {
		t.Errorf("Expected 2 users, got %d", len(users))
	}
}
