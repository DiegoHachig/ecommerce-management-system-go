package users

import "fmt"

// ShowUsers muestra los usuarios registrados.
func ShowUsers(users []User) {
	for _, user := range users {
		fmt.Println(
			user.GetID(),
			user.GetName(),
			user.GetEmail(),
		)
	}
}
