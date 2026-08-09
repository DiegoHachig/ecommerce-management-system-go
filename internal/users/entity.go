package users

import "gorm.io/gorm"

type UserEntity struct {
	gorm.Model

	Name     string
	Email    string
	Password string
}
