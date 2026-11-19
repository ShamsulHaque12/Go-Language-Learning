package domain

import (
	"time"
)

// User represents the domain model and database entity for users
type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	Email     string    `gorm:"type:varchar(150);uniqueIndex;not null" json:"email"`
	Password  string    `gorm:"type:varchar(255);not null" json:"-"`
	Role      string    `gorm:"type:varchar(50);default:'user'" json:"role"`
	Balance   int64     `gorm:"default:0" json:"balance"` // balance in cents or smallest currency unit
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// RegisterRequest represents the payload for registering a new user
type RegisterRequest struct {
	Name            string `json:"name" binding:"required,min=2,max=100"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required,min=6"`
	InitialBalance  int64  `json:"initial_balance" binding:"gte=0"`
}

// LoginRequest represents the payload for user authentication
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// UpdateProfileRequest represents the payload for updating user details
type UpdateProfileRequest struct {
	Name string `json:"name" binding:"required,min=2,max=100"`
}

// TransferBalanceRequest represents the payload for balance transfer
type TransferBalanceRequest struct {
	ToUserID uint  `json:"to_user_id" binding:"required,gt=0"`
	Amount   int64 `json:"amount" binding:"required,gt=0"`
}

// UserResponse represents safe user data to be returned to clients
type UserResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// AuthResponse represents the response containing the generated token and user profile
type AuthResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}

// ToResponse converts a domain User entity into a UserResponse DTO
func (u *User) ToResponse() UserResponse {
	return UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Role:      u.Role,
		Balance:   u.Balance,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
