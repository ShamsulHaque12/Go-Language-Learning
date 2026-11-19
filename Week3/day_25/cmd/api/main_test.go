package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{ID: 1, Name: "Sujon"},
		{ID: 2, Name: "Rahim"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

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
		t.Fatalf("Failed to decode JSON response: %v", err)
	}

	if len(users) != 2 {
		t.Errorf("Expected 2 users, got %d", len(users))
	}
}

func TestGetUsersResponseData(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	recorder := httptest.NewRecorder()

	getUsers(recorder, req)

	var users []User

	err := json.NewDecoder(recorder.Body).Decode(&users)
	if err != nil {
		t.Fatalf("Failed to decode JSON: %v", err)
	}

	if len(users) != 2 {
		t.Fatalf("Expected 2 users, got %d", len(users))
	}

	if users[0].Name != "Sujon" {
		t.Errorf("Expected Sujon, got %s", users[0].Name)
	}

	if users[1].Name != "Rahim" {
		t.Errorf("Expected Rahim, got %s", users[1].Name)
	}
}
