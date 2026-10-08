package services
/*
import (
	"time"

	"go-language-learning/Week2/Day11/exercise1/dto"
	"go-language-learning/Week2/Day11/exercise1/models"
	"go-language-learning/Week2/Day11/exercise1/repositories"
)

// UserService interface defines business logic operations
type UserService interface {
	GetAllUsers() []dto.UserResponse
	GetUserByID(id int) (dto.UserResponse, error)
	CreateUser(req dto.CreateUserRequest) (dto.UserResponse, error)
}

// userService struct holds dependency on UserRepository
type userService struct {
	repo repositories.UserRepository
}

// NewUserService constructs a new UserService instance
func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetAllUsers() []dto.UserResponse {
	users := s.repo.GetAll()
	var responses []dto.UserResponse
	for _, u := range users {
		responses = append(responses, toUserResponse(u))
	}
	return responses
}

func (s *userService) GetUserByID(id int) (dto.UserResponse, error) {
	u, err := s.repo.GetByID(id)
	if err != nil {
		return dto.UserResponse{}, err
	}
	return toUserResponse(u), nil
}

func (s *userService) CreateUser(req dto.CreateUserRequest) (dto.UserResponse, error) {
	// Business Logic: Password hashing, mapping DTO to Domain Model
	newUser := models.User{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: "hashed_" + req.Password, // Mock Password Hashing logic
		Role:         req.Role,
		CreatedAt:    time.Now(),
	}

	createdUser, err := s.repo.Create(newUser)
	if err != nil {
		return dto.UserResponse{}, err
	}

	return toUserResponse(createdUser), nil
}

// Helper mapping function: Domain Model -> Response DTO
func toUserResponse(u models.User) dto.UserResponse {
	return dto.UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Role:      u.Role,
		CreatedAt: u.CreatedAt.Format(time.RFC3339),
	}
}
*/