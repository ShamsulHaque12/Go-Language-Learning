package database
/*
import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	IsActive bool   `json:"is_active"`
}

func ConnectDB() (*sql.DB, error) {
	connStr := "postgresql://neondb_owner:npg_wZk1aJo5hQTY@ep-summer-waterfall-b307hnfl-pooler.c-4.ap-southeast-1.aws.neon.tech/neondb?sslmode=require&channel_binding=require"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	fmt.Println("PostgreSQL connected successfully!")
	return db, nil
}

func CreateUsersTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL,
		age INT NOT NULL,
		is_active BOOLEAN DEFAULT TRUE
	);`

	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create users table: %w", err)
	}

	fmt.Println("Users table checked/created successfully!")
	return nil
}

func GetUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT id, name, email, age, is_active FROM users ORDER BY id ASC")
	if err != nil {
		return nil, fmt.Errorf("failed to query users: %w", err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Age, &u.IsActive); err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during row iteration: %w", err)
	}

	return users, nil
}

func GetUserByID(db *sql.DB, id int) (*User, error) {
	query := `SELECT id, name, email, age, is_active FROM users WHERE id = $1;`
	var user User
	err := db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Age, &user.IsActive)
	if err != nil {
		return nil, fmt.Errorf("failed to query user by id: %w", err)
	}
	return &user, nil
}

func CreateUser(db *sql.DB, user *User) (*User, error) {
	query := `
	INSERT INTO users (name, email, age, is_active)
	VALUES ($1, $2, $3, $4)
	RETURNING id;`

	err := db.QueryRow(query, user.Name, user.Email, user.Age, user.IsActive).Scan(&user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to insert user: %w", err)
	}

	return user, nil
}

func UpdateUser(db *sql.DB, user *User) error {
	query := `
	UPDATE users
	SET name = $1, email = $2, age = $3, is_active = $4
	WHERE id = $5;`

	result, err := db.Exec(query, user.Name, user.Email, user.Age, user.IsActive, user.ID)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func DeleteUser(db *sql.DB, id int) error {
	query := `DELETE FROM users WHERE id = $1;`

	result, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

*/