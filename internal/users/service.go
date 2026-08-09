package users

import "fmt"

type UserService struct {
	repository UserRepository
}

func NewUserService(repository UserRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (s *UserService) RegisterUser(user User) error {
	if user.GetName() == "" {
		return fmt.Errorf("nombre invalido")
	}

	return s.repository.Save(user)
}

type MemoryRepository struct {
	users []User
}

func (r *MemoryRepository) Save(user User) error {
	r.users = append(r.users, user)
	return nil
}

func (r *MemoryRepository) GetAll() ([]User, error) {
	return r.users, nil
}
