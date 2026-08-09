package users

// UserRepository define las operaciones que
// cualquier repositorio de usuarios debe implementar.
type UserRepository interface {
	Save(user User) error
	GetAll() ([]User, error)
}
