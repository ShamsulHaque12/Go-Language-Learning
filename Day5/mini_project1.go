package main

/*
import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// ==========================================
// 1. STRUCT DEFINITIONS
// ==========================================

// Config holds app settings loaded from .env
type Config struct {
	AppName  string
	AppEnv   string
	DataFile string
}

// User represents a user entity in our system
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UserManager manages user data operations
type UserManager struct {
	filePath string
	users    []User
}

// ==========================================
// 2. CONFIG & ENV FUNCTIONS
// ==========================================

func loadConfig() Config {
	// 9. .env Loading
	err := godotenv.Load()
	if err != nil {
		log.Println("Note: .env file not found, using default environment variables.")
	}

	// 8. Environment Variables (os.Getenv)
	appName := os.Getenv("APP_NAME")
	if appName == "" {
		appName = "User Management Service"
	}

	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "development"
	}

	dataFile := os.Getenv("DATA_FILE")
	if dataFile == "" {
		dataFile = "users_db.json" // default JSON file
	}

	return Config{
		AppName:  appName,
		AppEnv:   appEnv,
		DataFile: dataFile,
	}
}

// ==========================================
// 3. FILE HANDLING & JSON (LOAD / SAVE)
// ==========================================

func NewUserManager(filePath string) (*UserManager, error) {
	manager := &UserManager{filePath: filePath}
	err := manager.loadFromFile()
	if err != nil {
		return nil, err
	}
	return manager, nil
}

// 4 & 5. File Handling & JSON Unmarshal
func (m *UserManager) loadFromFile() error {
	// Check if file exists
	if _, err := os.Stat(m.filePath); os.IsNotExist(err) {
		m.users = []User{}
		return m.saveToFile() // Create empty file if not existing
	}

	data, err := os.ReadFile(m.filePath)
	if err != nil {
		// 3. Error Handling
		return fmt.Errorf("failed to read file %s: %w", m.filePath, err)
	}

	if len(data) == 0 {
		m.users = []User{}
		return nil
	}

	err = json.Unmarshal(data, &m.users)
	if err != nil {
		return fmt.Errorf("failed to parse JSON data: %w", err)
	}

	return nil
}

// 4 & 5. File Handling & JSON Marshal
func (m *UserManager) saveToFile() error {
	data, err := json.MarshalIndent(m.users, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal users to JSON: %w", err)
	}

	err = os.WriteFile(m.filePath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write to file %s: %w", m.filePath, err)
	}

	return nil
}

// ==========================================
// 6. CRUD OPERATIONS
// ==========================================

// CREATE: Add new user
func (m *UserManager) Create(name, email, role string) (*User, error) {
	if name == "" || email == "" {
		return nil, errors.New("name and email cannot be empty")
	}

	// Check for duplicate email
	for _, u := range m.users {
		if u.Email == email {
			return nil, fmt.Errorf("user with email '%s' already exists", email)
		}
	}

	nextID := 1
	if len(m.users) > 0 {
		nextID = m.users[len(m.users)-1].ID + 1
	}

	// 7. time package (time.Now)
	now := time.Now()

	newUser := User{
		ID:        nextID,
		Name:      name,
		Email:     email,
		Role:      role,
		CreatedAt: now,
		UpdatedAt: now,
	}

	m.users = append(m.users, newUser)

	err := m.saveToFile()
	if err != nil {
		return nil, err
	}

	return &newUser, nil
}

// READ ALL: Get all users
func (m *UserManager) GetAll() []User {
	return m.users
}

// READ ONE: Find user by ID
func (m *UserManager) GetByID(id int) (*User, error) {
	for _, u := range m.users {
		if u.ID == id {
			return &u, nil
		}
	}
	return nil, fmt.Errorf("user with ID %d not found", id)
}

// UPDATE: Update user info by ID
func (m *UserManager) Update(id int, name, role string) (*User, error) {
	for i, u := range m.users {
		if u.ID == id {
			if name != "" {
				m.users[i].Name = name
			}
			if role != "" {
				m.users[i].Role = role
			}
			// 7. time package (updating timestamp)
			m.users[i].UpdatedAt = time.Now()

			err := m.saveToFile()
			if err != nil {
				return nil, err
			}
			return &m.users[i], nil
		}
	}
	return nil, fmt.Errorf("cannot update: user with ID %d not found", id)
}

// DELETE: Remove user by ID
func (m *UserManager) Delete(id int) error {
	foundIndex := -1
	for i, u := range m.users {
		if u.ID == id {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		return fmt.Errorf("cannot delete: user with ID %d not found", id)
	}

	// Remove element from slice
	m.users = append(m.users[:foundIndex], m.users[foundIndex+1:]...)

	return m.saveToFile()
}

// ==========================================
// MAIN FUNCTION (DEMO FLOW)
// ==========================================

func main() {
	fmt.Println("===========================================")
	fmt.Println("  GO BACKEND MINI PROJECT: USER MANAGEMENT ")
	fmt.Println("===========================================")

	// 1. Load Configurations (.env & os.Getenv)
	cfg := loadConfig()
	fmt.Printf(" App Name : %s\n", cfg.AppName)
	fmt.Printf(" App Env  : %s\n", cfg.AppEnv)
	fmt.Printf(" Data File: %s\n", cfg.DataFile)
	fmt.Println("-------------------------------------------")

	// 2. Initialize Manager (File Handling & JSON)
	manager, err := NewUserManager(cfg.DataFile)
	if err != nil {
		log.Fatalf("Fatal Error initializing manager: %v\n", err)
	}

	// 3. CREATE Operations
	fmt.Println("\n [1] --- Creating Users (CREATE) ---")
	u1, err := manager.Create("Sujon", "sujon@gmail.com", "Go Developer")
	if err != nil {
		fmt.Println("Error creating user:", err)
	} else {
		fmt.Printf("Created User: ID=%d, Name=%s, Time=%s\n", u1.ID, u1.Name, u1.CreatedAt.Format("15:04:05"))
	}

	u2, err := manager.Create("Shamsul", "shamsul@example.com", "Fullstack Engineer")
	if err != nil {
		fmt.Println("Error creating user:", err)
	} else {
		fmt.Printf("Created User: ID=%d, Name=%s, Time=%s\n", u2.ID, u2.Name, u2.CreatedAt.Format("15:04:05"))
	}

	// 4. READ Operations
	fmt.Println("\n [2] --- Listing All Users (READ ALL) ---")
	allUsers := manager.GetAll()
	for _, u := range allUsers {
		fmt.Printf("- ID: %d | Name: %-10s | Role: %-18s | Created: %s\n",
			u.ID, u.Name, u.Role, u.CreatedAt.Format("2006-01-02 15:04:05"))
	}

	// 5. UPDATE Operations
	fmt.Println("\n [3] --- Updating User ID 1 (UPDATE) ---")
	updatedUser, err := manager.Update(1, "Sujon Haque", "Senior Go Developer")
	if err != nil {
		fmt.Println("Update Error:", err)
	} else {
		fmt.Printf("Updated User: ID=%d, Name=%s, Role=%s, UpdatedAt=%s\n",
			updatedUser.ID, updatedUser.Name, updatedUser.Role, updatedUser.UpdatedAt.Format("15:04:05"))
	}

	// 6. READ ONE Operation
	fmt.Println("\n [4] --- Fetching User ID 1 (READ ONE) ---")
	foundUser, err := manager.GetByID(1)
	if err != nil {
		fmt.Println("Get User Error:", err)
	} else {
		fmt.Printf("Found User: %+v\n", *foundUser)
	}

	// 7. DELETE Operation (Optional Demo)
	fmt.Println("\n [5] --- Deleting User ID 2 (DELETE) ---")
	err = manager.Delete(2)
	if err != nil {
		fmt.Println("Delete Error:", err)
	} else {
		fmt.Println("User ID 2 successfully deleted.")
	}

	// 8. FINAL READ ALL
	fmt.Println("\n [6] --- Final User List ---")
	for _, u := range manager.GetAll() {
		fmt.Printf("- ID: %d | Name: %-12s | Role: %-20s\n", u.ID, u.Name, u.Role)
	}

	fmt.Println("\n===========================================")
	fmt.Printf(" Data successfully persisted in '%s'\n", cfg.DataFile)
	fmt.Println("===========================================")
}

*/
