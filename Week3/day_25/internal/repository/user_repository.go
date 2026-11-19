package repository

import (
	"context"
	"errors"
	"fmt"

	"day_25/internal/domain"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	GetByID(ctx context.Context, id uint) (*domain.User, error)
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
	GetAll(ctx context.Context, offset, limit int) ([]domain.User, int64, error)
	Update(ctx context.Context, user *domain.User) error
	Delete(ctx context.Context, id uint) error
	TransferBalance(ctx context.Context, fromID, toID uint, amount int64) error
	Migrate() error
}

type postgresUserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &postgresUserRepository{db: db}
}

// Migrate performs schema auto-migration for User model
func (r *postgresUserRepository) Migrate() error {
	return r.db.AutoMigrate(&domain.User{})
}

// Create stores a new user record
func (r *postgresUserRepository) Create(ctx context.Context, user *domain.User) error {
	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		return err
	}
	return nil
}

// GetByID retrieves a user by primary key
func (r *postgresUserRepository) GetByID(ctx context.Context, id uint) (*domain.User, error) {
	var user domain.User
	if err := r.db.WithContext(ctx).First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

// GetByEmail retrieves a user by unique email
func (r *postgresUserRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

// GetAll returns a paginated slice of users and the total count
func (r *postgresUserRepository) GetAll(ctx context.Context, offset, limit int) ([]domain.User, int64, error) {
	var users []domain.User
	var total int64

	if err := r.db.WithContext(ctx).Model(&domain.User{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.WithContext(ctx).Offset(offset).Limit(limit).Order("id ASC").Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

// Update saves modified user fields
func (r *postgresUserRepository) Update(ctx context.Context, user *domain.User) error {
	result := r.db.WithContext(ctx).Save(user)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domain.ErrUserNotFound
	}
	return nil
}

// Delete removes a user record
func (r *postgresUserRepository) Delete(ctx context.Context, id uint) error {
	result := r.db.WithContext(ctx).Delete(&domain.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return domain.ErrUserNotFound
	}
	return nil
}

// TransferBalance executes an atomic fund transfer between two users with row locking
func (r *postgresUserRepository) TransferBalance(ctx context.Context, fromID, toID uint, amount int64) error {
	if amount <= 0 {
		return domain.ErrInvalidAmount
	}
	if fromID == toID {
		return domain.ErrSameAccountTransfer
	}

	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Enforce consistent locking order to prevent database deadlocks
		firstID, secondID := fromID, toID
		if firstID > secondID {
			firstID, secondID = toID, fromID
		}

		// Lock first account
		var firstUser domain.User
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&firstUser, firstID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return fmt.Errorf("account %d not found: %w", firstID, domain.ErrUserNotFound)
			}
			return err
		}

		// Lock second account
		var secondUser domain.User
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&secondUser, secondID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return fmt.Errorf("account %d not found: %w", secondID, domain.ErrUserNotFound)
			}
			return err
		}

		// Identify sender and receiver from locked rows
		var sender, receiver *domain.User
		if firstUser.ID == fromID {
			sender = &firstUser
			receiver = &secondUser
		} else {
			sender = &secondUser
			receiver = &firstUser
		}

		// Validate sufficient funds
		if sender.Balance < amount {
			return domain.ErrInsufficientBalance
		}

		// Debit sender
		if err := tx.Model(sender).Update("balance", gorm.Expr("balance - ?", amount)).Error; err != nil {
			return fmt.Errorf("failed to debit sender: %w", err)
		}

		// Credit receiver
		if err := tx.Model(receiver).Update("balance", gorm.Expr("balance + ?", amount)).Error; err != nil {
			return fmt.Errorf("failed to credit receiver: %w", err)
		}

		return nil
	})
}
