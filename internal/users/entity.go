package users

// UserEntity representa la estructura de la tabla
// de usuarios almacenada en MySQL mediante GORM.
import "gorm.io/gorm"

type UserEntity struct {
	gorm.Model

	Name     string
	Email    string
	Password string
}
