package models

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	IsActive bool   `json:"is_active"`
}

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required,min=2,max=100"`
	Email    string `json:"email" binding:"required,email"`
	Age      int    `json:"age" binding:"required,gt=0,lte=150"`
	IsActive *bool  `json:"is_active" binding:"required"`
}

type UpdateUserRequest struct {
	Name     string `json:"name" binding:"required,min=2,max=100"`
	Email    string `json:"email" binding:"required,email"`
	Age      int    `json:"age" binding:"required,gt=0,lte=150"`
	IsActive *bool  `json:"is_active" binding:"required"`
}

type PatchUserRequest struct {
	Name     *string `json:"name" binding:"omitempty,min=2,max=100"`
	Email    *string `json:"email" binding:"omitempty,email"`
	Age      *int    `json:"age" binding:"omitempty,gt=0,lte=150"`
	IsActive *bool   `json:"is_active"`
}
