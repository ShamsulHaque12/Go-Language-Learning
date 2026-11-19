package service_test

import (
	"context"
	"errors"

	"day_25/internal/domain"
	"day_25/internal/repository"
)

type mockUserRepository struct {
	users       map[uint]*domain.User
	usersByMail map[string]*domain.User
	nextID      uint
}

func newMockUserRepository() repository.UserRepository {
	return &mockUserRepository{
		users:       make(map[uint]*domain.User),
		usersByMail: make(map[string]*domain.User),
		nextID:      1,
	}
}

func (m *mockUserRepository) Migrate() error {
	return nil
}

func (m *mockUserRepository) Create(ctx context.Context, user *domain.User) error {
	if _, exists := m.usersByMail[user.Email]; exists {
		return errors.New("unique violation: email already exists")
	}
	user.ID = m.nextID
	m.nextID++
	m.users[user.ID] = user
	m.usersByMail[user.Email] = user
	return nil
}

func (m *mockUserRepository) GetByID(ctx context.Context, id uint) (*domain.User, error) {
	u, exists := m.users[id]
	if !exists {
		return nil, domain.ErrUserNotFound
	}
	return u, nil
}

func (m *mockUserRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	u, exists := m.usersByMail[email]
	if !exists {
		return nil, domain.ErrUserNotFound
	}
	return u, nil
}

func (m *mockUserRepository) GetAll(ctx context.Context, offset, limit int) ([]domain.User, int64, error) {
	list := make([]domain.User, 0, len(m.users))
	for _, u := range m.users {
		list = append(list, *u)
	}
	return list, int64(len(list)), nil
}

func (m *mockUserRepository) Update(ctx context.Context, user *domain.User) error {
	if _, exists := m.users[user.ID]; !exists {
		return domain.ErrUserNotFound
	}
	m.users[user.ID] = user
	m.usersByMail[user.Email] = user
	return nil
}

func (m *mockUserRepository) Delete(ctx context.Context, id uint) error {
	u, exists := m.users[id]
	if !exists {
		return domain.ErrUserNotFound
	}
	delete(m.users, id)
	delete(m.usersByMail, u.Email)
	return nil
}

func (m *mockUserRepository) TransferBalance(ctx context.Context, fromID, toID uint, amount int64) error {
	fromUser, exists := m.users[fromID]
	if !exists {
		return domain.ErrUserNotFound
	}
	toUser, exists := m.users[toID]
	if !exists {
		return domain.ErrUserNotFound
	}
	if fromUser.Balance < amount {
		return domain.ErrInsufficientBalance
	}
	fromUser.Balance -= amount
	toUser.Balance += amount
	return nil
}
