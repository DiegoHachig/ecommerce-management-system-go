package users

import "github.com/DiegoHachig/ecommerce-management-system-go/database"

type MySQLRepository struct{}

// Save guarda un usuario en MySQL.
func (r *MySQLRepository) Save(user User) error {

	entity := UserEntity{
		Name:     user.GetName(),
		Email:    user.GetEmail(),
		Password: "",
	}

	return database.DB.Create(&entity).Error
}

// GetAll obtiene todos los usuarios.
func (r *MySQLRepository) GetAll() ([]User, error) {

	var entities []UserEntity

	result := database.DB.Find(&entities)

	if result.Error != nil {
		return nil, result.Error
	}

	var usersList []User

	for _, entity := range entities {

		user, _ := NewUser(
			int(entity.ID),
			entity.Name,
			entity.Email,
			"",
		)

		usersList = append(usersList, *user)
	}

	return usersList, nil
}
