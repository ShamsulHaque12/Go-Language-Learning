package repositories
/*
import (
	"errors"
	"time"

	"go-language-learning/Week2/Day11/exercise1/models"
)

// UserRepository interface decouples storage logic
type UserRepository interface {
	GetAll() []models.User
	GetByID(id int) (models.User, error)
	Create(user models.User) (models.User, error)
}

// userRepository implements UserRepository interface
type userRepository struct {
	users []models.User
}

// NewUserRepository constructs a new UserRepository instance
func NewUserRepository() UserRepository {
	return &userRepository{
		users: models.UserDB,
	}
}

func (r *userRepository) GetAll() []models.User {
	return r.users
}

func (r *userRepository) GetByID(id int) (models.User, error) {
	for _, u := range r.users {
		if u.ID == id {
			return u, nil
		}
	}
	return models.User{}, errors.New("user not found")
}

func (r *userRepository) Create(user models.User) (models.User, error) {
	// Duplicate email validation check
	for _, u := range r.users {
		if u.Email == user.Email {
			return models.User{}, errors.New("user with this email already exists")
		}
	}

	user.ID = len(r.users) + 1
	if user.CreatedAt.IsZero() {
		user.CreatedAt = time.Now()
	}
	r.users = append(r.users, user)
	return user, nil
}
*/