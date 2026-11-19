package models

/*
import "time"

// User representaion of Database / Domain Entity
type User struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"` // PasswordHash DB-তে থাকবে কিন্তু JSON API-তে রিটার্ন হবে না
	Role         string    `json:"role"`
	CreatedAt    time.Time `json:"created_at"`
}

// UserDB is a mock database slice
var UserDB = []User{
	{
		ID:           1,
		Name:         "Sujon",
		Email:        "sujon@gmail.com",
		PasswordHash: "$2a$12$eImiTXuWVxfM37uY4JANjO8...", // Hashed password
		Role:         "admin",
		CreatedAt:    time.Now(),
	},
}
*/