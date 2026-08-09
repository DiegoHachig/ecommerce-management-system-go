package users

import "errors"

// User representa un usuario del sistema.
type User struct {
	id       int
	name     string
	email    string
	password string
}

// Constructor
func NewUser(
	id int,
	name string,
	email string,
	password string,
) (*User, error) {

	if name == "" {
		return nil, errors.New("el nombre es obligatorio")
	}

	if email == "" {
		return nil, errors.New("el correo es obligatorio")
	}

	return &User{
		id:       id,
		name:     name,
		email:    email,
		password: password,
	}, nil
}

// Getters

func (u User) GetID() int {
	return u.id
}

func (u User) GetName() string {
	return u.name
}

func (u User) GetEmail() string {
	return u.email
}
