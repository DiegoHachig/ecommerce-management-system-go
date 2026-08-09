package users

import "github.com/DiegoHachig/ecommerce-management-system-go/database"

// MySQLRepository implementa el acceso a datos
// utilizando MySQL y GORM para la persistencia
// de usuarios.
type MySQLRepository struct{}

// Save almacena un usuario en la base de datos
// convirtiendo el modelo de negocio en una entidad GORM.
func (r *MySQLRepository) Save(user User) error {
	// Conversión del modelo de dominio
	// a entidad persistente de base de datos.
	entity := UserEntity{
		Name:     user.GetName(),
		Email:    user.GetEmail(),
		Password: "",
	}

	return database.DB.Create(&entity).Error
}

// GetAll recupera todos los usuarios almacenados
// en MySQL y los transforma en objetos User.
func (r *MySQLRepository) GetAll() ([]User, error) {

	var entities []UserEntity

	result := database.DB.Find(&entities)

	if result.Error != nil {
		return nil, result.Error
	}

	var usersList []User
	// Conversión de entidades GORM
	// a objetos de negocio User.
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
