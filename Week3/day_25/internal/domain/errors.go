package domain

import "errors"

var (
	ErrUserNotFound        = errors.New("user not found")
	ErrEmailAlreadyExists  = errors.New("email is already registered")
	ErrInvalidCredentials  = errors.New("invalid email or password")
	ErrInsufficientBalance = errors.New("insufficient balance for transfer")
	ErrInvalidAmount       = errors.New("transfer amount must be greater than zero")
	ErrSameAccountTransfer = errors.New("source and destination accounts must be different")
	ErrUnauthorized        = errors.New("unauthorized access")
	ErrForbidden           = errors.New("permission denied")
	ErrInternalServer      = errors.New("internal server error")

	// Product errors
	ErrProductNotFound = errors.New("product not found")
	ErrDuplicateSKU    = errors.New("product with this SKU already exists")
	ErrInvalidPrice    = errors.New("price must be greater than or equal to 0")
	ErrInvalidStock    = errors.New("stock cannot be negative")
)
