package users

import "fmt"

// UserService contiene la lógica de negocio
// relacionada con la gestión de usuarios
type UserService struct {
	repository UserRepository
}

// NewUserService crea una nueva instancia
// del servicio de usuarios.
func NewUserService(repository UserRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}

// RegisterUser valida y registra un usuario
// utilizando el repositorio configurado.
func (s *UserService) RegisterUser(user User) error {
	if user.GetName() == "" {
		return fmt.Errorf("nombre invalido")
	}

	return s.repository.Save(user)
}

// MemoryRepository simula una base de datos
// almacenando usuarios temporalmente en memoria.
type MemoryRepository struct {
	users []User
}

// Save almacena un usuario en memoria.
func (r *MemoryRepository) Save(user User) error {
	r.users = append(r.users, user)
	return nil
}

// GetAll devuelve todos los usuarios registrados.
func (r *MemoryRepository) GetAll() ([]User, error) {
	return r.users, nil
}
