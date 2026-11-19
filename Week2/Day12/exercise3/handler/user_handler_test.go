package handler

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"

	"go-language-learning/Week2/Day12/exercise3/models"
)

type mockUserService struct {
	getAllUsersFunc func() ([]models.User, error)
	getUserByIDFunc func(id int) (*models.User, error)
	createUserFunc  func(req *models.CreateUserRequest) (*models.User, error)
	updateUserFunc  func(id int, req *models.UpdateUserRequest) (*models.User, error)
	patchUserFunc   func(id int, req *models.PatchUserRequest) (*models.User, error)
	deleteUserFunc  func(id int) error
}

func (m *mockUserService) GetAllUsers() ([]models.User, error) {
	if m.getAllUsersFunc != nil {
		return m.getAllUsersFunc()
	}
	return []models.User{}, nil
}

func (m *mockUserService) GetUserByID(id int) (*models.User, error) {
	if m.getUserByIDFunc != nil {
		return m.getUserByIDFunc(id)
	}
	return nil, sql.ErrNoRows
}

func (m *mockUserService) CreateUser(req *models.CreateUserRequest) (*models.User, error) {
	if m.createUserFunc != nil {
		return m.createUserFunc(req)
	}
	return &models.User{ID: 1, Name: req.Name, Email: req.Email, Age: req.Age, IsActive: *req.IsActive}, nil
}

func (m *mockUserService) UpdateUser(id int, req *models.UpdateUserRequest) (*models.User, error) {
	if m.updateUserFunc != nil {
		return m.updateUserFunc(id, req)
	}
	return &models.User{ID: id, Name: req.Name, Email: req.Email, Age: req.Age, IsActive: *req.IsActive}, nil
}

func (m *mockUserService) PatchUser(id int, req *models.PatchUserRequest) (*models.User, error) {
	if m.patchUserFunc != nil {
		return m.patchUserFunc(id, req)
	}
	return &models.User{ID: id, Name: "Patched", Email: "patched@example.com", Age: 30, IsActive: true}, nil
}

func (m *mockUserService) DeleteUser(id int) error {
	if m.deleteUserFunc != nil {
		return m.deleteUserFunc(id)
	}
	return nil
}

func setupTestRouter(svc *mockUserService) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	h := NewUserHandler(svc)
	h.RegisterRoutes(r)
	return r
}

func TestGetAllUsers_Success(t *testing.T) {
	// 1. Mock service তৈরি করো
	//    getAllUsersFunc থেকে একটি user এবং nil error return করবে।
	mockSvc := &mockUserService{
		getAllUsersFunc: func() ([]models.User, error) {
			return []models.User{
				{
					ID:       1,
					Name:     "Alice",
					Email:    "alice@example.com",
					Age:      25,
					IsActive: true,
				},
			}, nil
		},
	}

	// 2. setupTestRouter(mockSvc) দিয়ে router তৈরি করো
	r := setupTestRouter(mockSvc)

	// 3. GET /users request তৈরি করো
	//    httptest.NewRequest() ব্যবহার করবে
	req := httptest.NewRequest(http.MethodGet, "/users", nil)

	// 4. httptest.NewRecorder() দিয়ে response ধরো
	w := httptest.NewRecorder()

	// 5. r.ServeHTTP(w, req) দিয়ে request চালাও
	r.ServeHTTP(w, req)

	// 6. যাচাই করো:
	//    - HTTP status 200
	//    - JSON response-এর status "success"
	//    - data-তে একটি user আছে
	//    - user-এর name "Alice"
	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	var resp struct {
		Status string        `json:"status"`
		Data   []models.User `json:"data"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp.Status != "success" {
		t.Errorf("expected status 'success', got '%s'", resp.Status)
	}

	if len(resp.Data) != 1 {
		t.Fatalf("expected 1 user in data, got %d", len(resp.Data))
	}

	if resp.Data[0].Name != "Alice" {
		t.Errorf("expected user name 'Alice', got '%s'", resp.Data[0].Name)
	}
}

func TestCreateUser_ValidationDetails(t *testing.T) {
	mockSvc := &mockUserService{}
	r := setupTestRouter(mockSvc)

	payload := []byte(`{"name": "Alice", "email": "not-an-email", "age": 25, "is_active": true}`)
	req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", w.Code)
	}

	var resp struct {
		Status  string            `json:"status"`
		Error   string            `json:"error"`
		Details map[string]string `json:"details"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to parse response JSON: %v", err)
	}

	if resp.Status != "error" {
		t.Errorf("expected status 'error', got %s", resp.Status)
	}
	if resp.Error != "Validation failed" {
		t.Errorf("expected error 'Validation failed', got %s", resp.Error)
	}
	if resp.Details["email"] != "Please provide a valid email address" {
		t.Errorf("expected email detail 'Please provide a valid email address', got '%s'", resp.Details["email"])
	}
}

func TestPatchUser_AtLeastOneFieldRequired(t *testing.T) {
	mockSvc := &mockUserService{}
	r := setupTestRouter(mockSvc)

	payload := []byte(`{}`)
	req, _ := http.NewRequest(http.MethodPatch, "/users/1", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", w.Code)
	}

	var resp map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to parse response JSON: %v", err)
	}
	if resp["error"] != "At least one field is required to update" {
		t.Errorf("expected 'At least one field is required to update', got '%v'", resp["error"])
	}
}

func TestGetUserByID_NotFound(t *testing.T) {
	mockSvc := &mockUserService{}
	r := setupTestRouter(mockSvc)

	req := httptest.NewRequest(http.MethodGet, "/users/999", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("expected status 404, got %d", w.Code)
	}

	var resp struct {
		Status string `json:"status"`
		Error  string `json:"error"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to parse response JSON: %v", err)
	}

	if resp.Status != "error" {
		t.Errorf("expected status 'error', got '%s'", resp.Status)
	}

	if resp.Error != "User not found" {
		t.Errorf("expected error 'User not found', got '%s'", resp.Error)
	}
}

func TestCreateUser_Success(t *testing.T) {
	mockSvc := &mockUserService{
		createUserFunc: func(req *models.CreateUserRequest) (*models.User, error) {
			return &models.User{
				ID:       1,
				Name:     req.Name,
				Email:    req.Email,
				Age:      req.Age,
				IsActive: *req.IsActive,
			}, nil
		},
	}
	r := setupTestRouter(mockSvc)

	payload := []byte(`{
		"name": "Alice",
		"email": "alice@example.com",
		"age": 25,
		"is_active": true
	}`)
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", w.Code)
	}

	var resp struct {
		Status string      `json:"status"`
		Data   models.User `json:"data"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to parse response JSON: %v", err)
	}

	if resp.Status != "success" {
		t.Errorf("expected status 'success', got '%s'", resp.Status)
	}

	if resp.Data.Name != "Alice" {
		t.Errorf("expected data.name 'Alice', got '%s'", resp.Data.Name)
	}

	if resp.Data.Email != "alice@example.com" {
		t.Errorf("expected data.email 'alice@example.com', got '%s'", resp.Data.Email)
	}
}

func TestDeleteUser_Success(t *testing.T) {
	mockSvc := &mockUserService{}
	r := setupTestRouter(mockSvc)

	req := httptest.NewRequest(http.MethodDelete, "/users/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	var resp struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to parse response JSON: %v", err)
	}

	if resp.Status != "success" {
		t.Errorf("expected status 'success', got '%s'", resp.Status)
	}

	if resp.Message != "User deleted successfully" {
		t.Errorf("expected message 'User deleted successfully', got '%s'", resp.Message)
	}
}

func TestPatchUser_Success(t *testing.T) {
	mockSvc := &mockUserService{}
	r := setupTestRouter(mockSvc)

	payload := []byte(`{"name": "Alice"}`)
	req := httptest.NewRequest(http.MethodPatch, "/users/1", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	var resp struct {
		Status string      `json:"status"`
		Data   models.User `json:"data"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to parse response JSON: %v", err)
	}

	if resp.Status != "success" {
		t.Errorf("expected status 'success', got '%s'", resp.Status)
	}

	if resp.Data.Name != "Patched" {
		t.Errorf("expected data.name 'Patched', got '%s'", resp.Data.Name)
	}
}

func TestUpdateUser_Success(t *testing.T) {
	mockSvc := &mockUserService{}
	r := setupTestRouter(mockSvc)

	payload := []byte(`{
		"name": "Bob Updated",
		"email": "bob.updated@example.com",
		"age": 30,
		"is_active": true
	}`)
	req := httptest.NewRequest(http.MethodPut, "/users/1", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	var resp struct {
		Status string      `json:"status"`
		Data   models.User `json:"data"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to parse response JSON: %v", err)
	}

	if resp.Status != "success" {
		t.Errorf("expected status 'success', got '%s'", resp.Status)
	}

	if resp.Data.Name != "Bob Updated" {
		t.Errorf("expected data.name 'Bob Updated', got '%s'", resp.Data.Name)
	}

	if resp.Data.Email != "bob.updated@example.com" {
		t.Errorf("expected data.email 'bob.updated@example.com', got '%s'", resp.Data.Email)
	}
}

func TestUpdateUser_NotFound(t *testing.T) {
	mockSvc := &mockUserService{
		updateUserFunc: func(id int, req *models.UpdateUserRequest) (*models.User, error) {
			return nil, sql.ErrNoRows
		},
	}
	r := setupTestRouter(mockSvc)

	payload := []byte(`{
		"name": "Bob Updated",
		"email": "bob.updated@example.com",
		"age": 30,
		"is_active": true
	}`)
	req := httptest.NewRequest(http.MethodPut, "/users/999", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("expected status 404, got %d", w.Code)
	}

	var resp struct {
		Status string `json:"status"`
		Error  string `json:"error"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to parse response JSON: %v", err)
	}

	if resp.Status != "error" {
		t.Errorf("expected status 'error', got '%s'", resp.Status)
	}

	if resp.Error != "User not found" {
		t.Errorf("expected error 'User not found', got '%s'", resp.Error)
	}
}

func TestUpdateUser_ValidationFailure(t *testing.T) {
	mockSvc := &mockUserService{}
	r := setupTestRouter(mockSvc)

	payload := []byte(`{"name": ""}`)
	req := httptest.NewRequest(http.MethodPut, "/users/1", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", w.Code)
	}

	var resp struct {
		Status string `json:"status"`
		Error  string `json:"error"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to parse response JSON: %v", err)
	}

	if resp.Status != "error" {
		t.Errorf("expected status 'error', got '%s'", resp.Status)
	}

	if resp.Error != "Validation failed" {
		t.Errorf("expected error 'Validation failed', got '%s'", resp.Error)
	}
}

func TestHandleError_NotFound(t *testing.T) {
	mockSvc := &mockUserService{
		getUserByIDFunc: func(id int) (*models.User, error) {
			return nil, sql.ErrNoRows
		},
	}
	r := setupTestRouter(mockSvc)

	req := httptest.NewRequest(http.MethodGet, "/users/999", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("expected status 404, got %d", w.Code)
	}

	var resp struct {
		Status string `json:"status"`
		Error  string `json:"error"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to parse response JSON: %v", err)
	}

	if resp.Status != "error" {
		t.Errorf("expected status 'error', got '%s'", resp.Status)
	}
	if resp.Error != "User not found" {
		t.Errorf("expected error 'User not found', got '%s'", resp.Error)
	}
}

func TestHandleError_DuplicateEmail(t *testing.T) {
	mockSvc := &mockUserService{
		createUserFunc: func(req *models.CreateUserRequest) (*models.User, error) {
			return nil, &pq.Error{Code: "23505"}
		},
	}
	r := setupTestRouter(mockSvc)

	payload := []byte(`{
		"name": "Alice",
		"email": "duplicate@example.com",
		"age": 25,
		"is_active": true
	}`)
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusConflict {
		t.Fatalf("expected status 409, got %d", w.Code)
	}

	var resp struct {
		Status string `json:"status"`
		Error  string `json:"error"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to parse response JSON: %v", err)
	}

	if resp.Status != "error" {
		t.Errorf("expected status 'error', got '%s'", resp.Status)
	}
	if resp.Error != "Email already exists" {
		t.Errorf("expected error 'Email already exists', got '%s'", resp.Error)
	}
}

func TestHandleError_InternalServerError(t *testing.T) {
	mockSvc := &mockUserService{
		getUserByIDFunc: func(id int) (*models.User, error) {
			return nil, errors.New("database connection failed")
		},
	}
	r := setupTestRouter(mockSvc)

	req := httptest.NewRequest(http.MethodGet, "/users/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("expected status 500, got %d", w.Code)
	}

	var resp struct {
		Status string `json:"status"`
		Error  string `json:"error"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to parse response JSON: %v", err)
	}

	if resp.Status != "error" {
		t.Errorf("expected status 'error', got '%s'", resp.Status)
	}
	if resp.Error != "Internal server error" {
		t.Errorf("expected error 'Internal server error', got '%s'", resp.Error)
	}
}

func TestHandleError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		err            error
		expectedStatus int
		expectedError  string
	}{
		{
			name:           "User not found",
			err:            sql.ErrNoRows,
			expectedStatus: http.StatusNotFound,
			expectedError:  "User not found",
		},
		{
			name:           "Duplicate email",
			err:            &pq.Error{Code: "23505"},
			expectedStatus: http.StatusConflict,
			expectedError:  "Email already exists",
		},
		{
			name:           "Internal server error",
			err:            errors.New("database connection failed"),
			expectedStatus: http.StatusInternalServerError,
			expectedError:  "Internal server error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(recorder)

			svc := &mockUserService{}
			h := NewUserHandler(svc)

			h.handleError(c, tt.err)

			if recorder.Code != tt.expectedStatus {
				t.Errorf(
					"expected status %d, got %d",
					tt.expectedStatus,
					recorder.Code,
				)
			}

			var response map[string]interface{}
			if err := json.Unmarshal(
				recorder.Body.Bytes(),
				&response,
			); err != nil {
				t.Fatalf("failed to parse response: %v", err)
			}

			if response["error"] != tt.expectedError {
				t.Errorf(
					"expected error %q, got %v",
					tt.expectedError,
					response["error"],
				)
			}
		})
	}
}

func TestDeleteUser_NotFound(t *testing.T) {
	svc := &mockUserService{
		deleteUserFunc: func(id int) error {
			return sql.ErrNoRows
		},
	}

	router := setupTestRouter(svc)
	req := httptest.NewRequest(
		http.MethodDelete,
		"/users/999",
		nil,
	)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusNotFound {
		t.Errorf(
			"expected status %d, got %d",
			http.StatusNotFound,
			recorder.Code,
		)
	}
}

func TestDeleteUser_InvalidID(t *testing.T) {
	svc := &mockUserService{}
	router := setupTestRouter(svc)

	req := httptest.NewRequest(
		http.MethodDelete,
		"/users/abc",
		nil,
	)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusBadRequest {
		t.Errorf(
			"expected status %d, got %d",
			http.StatusBadRequest,
			recorder.Code,
		)
	}
}

func TestGetAllUsers_ServiceError(t *testing.T) {
	svc := &mockUserService{
		getAllUsersFunc: func() ([]models.User, error) {
			return nil, errors.New("database connection failed")
		},
	}

	router := setupTestRouter(svc)

	req := httptest.NewRequest(
		http.MethodGet,
		"/users",
		nil,
	)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusInternalServerError {
		t.Errorf(
			"expected status %d, got %d",
			http.StatusInternalServerError,
			recorder.Code,
		)
	}
}
