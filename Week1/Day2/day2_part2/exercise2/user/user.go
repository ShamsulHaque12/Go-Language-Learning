package user

type User struct {
	ID    int
	Name  string
	Email string
}

func NewUser(id int, name string, email string) User {
	if name == "" || email == "" {
		panic("Name and Email are required")
	}
	return User{ID: id, Name: name, Email: email}
}

func (u User) GetID() int {
	return u.ID
}

func (u User) GetName() string {
	return u.Name
}

func (u User) GetEmail() string {
	return u.Email
}