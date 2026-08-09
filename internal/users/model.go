package users

import "errors"

// User representa un usuario del sistema.
// dentro del sistema de e-commerce.
type User struct {
	id       int
	name     string
	email    string
	password string
}

// Constructor
// NewUser crea una nueva instancia de User
// validando los datos recibidos.
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
// GetID devuelve el identificador del usuario.
func (u User) GetID() int {
	return u.id
}

// GetName devuelve el nombre del usuario.
func (u User) GetName() string {
	return u.name
}

// GetEmail devuelve el email del usuario.
func (u User) GetEmail() string {
	return u.email
}
