package models
/*
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

// In-memory user database
var users = []User{
	{ID: 1, Name: "Sujon", Email: "sujon@gmail.com", Age: 27},
	{ID: 2, Name: "Rahim", Email: "rahim@gmail.com", Age: 25},
	{ID: 3, Name: "Karim", Email: "karim@gmail.com", Age: 30},
}

// GetAllUsers returns all users
func GetAllUsers() []User {
	return users
}

// GetUserByID searches a user by ID
func GetUserByID(id int) (User, bool) {
	for _, user := range users {
		if user.ID == id {
			return user, true
		}
	}
	return User{}, false
}

// CreateUser generates auto ID and adds user
func CreateUser(user *User) {
	maxID := 0
	for _, u := range users {
		if u.ID > maxID {
			maxID = u.ID
		}
	}
	user.ID = maxID + 1
	users = append(users, *user)
}

// UpdateUser updates user by ID
func UpdateUser(id int, updatedUser User) (User, bool) {
	for i, user := range users {
		if user.ID == id {
			updatedUser.ID = id
			users[i] = updatedUser
			return users[i], true
		}
	}
	return User{}, false
}

// DeleteUser deletes user by ID
func DeleteUser(id int) bool {
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			return true
		}
	}
	return false
}
*/