package repository

import (
	"database/sql"
	"fmt"

	"go-language-learning/Week2/Day12/exercise2/models"
)

type UserRepository interface {
	GetAll() ([]models.User, error)
	GetByID(id int) (*models.User, error)
	Create(user *models.User) (*models.User, error)
	Update(user *models.User) error
	Delete(id int) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetAll() ([]models.User, error) {
	query := `SELECT id, name, email, age, is_active FROM users ORDER BY id ASC;`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query users: %w", err)
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Age, &u.IsActive); err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating users: %w", err)
	}

	return users, nil
}

func (r *userRepository) GetByID(id int) (*models.User, error) {
	query := `SELECT id, name, email, age, is_active FROM users WHERE id = $1;`
	var u models.User
	err := r.db.QueryRow(query, id).Scan(&u.ID, &u.Name, &u.Email, &u.Age, &u.IsActive)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *userRepository) Create(user *models.User) (*models.User, error) {
	query := `
	INSERT INTO users (name, email, age, is_active)
	VALUES ($1, $2, $3, $4)
	RETURNING id;`

	err := r.db.QueryRow(query, user.Name, user.Email, user.Age, user.IsActive).Scan(&user.ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) Update(user *models.User) error {
	query := `
	UPDATE users
	SET name = $1, email = $2, age = $3, is_active = $4
	WHERE id = $5;`

	res, err := r.db.Exec(query, user.Name, user.Email, user.Age, user.IsActive, user.ID)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *userRepository) Delete(id int) error {
	query := `DELETE FROM users WHERE id = $1;`

	res, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}
