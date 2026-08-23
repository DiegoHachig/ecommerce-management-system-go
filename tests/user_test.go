package tests

import (
	"testing"

	"github.com/DiegoHachig/ecommerce-management-system-go/internal/users"
)

func TestCreateUser(t *testing.T) {

	user, err := users.NewUser(
		1,
		"Diego Hachig",
		"diego@email.com",
		"1234567890",
	)

	if err != nil {
		t.Errorf("No se pudo crear el usuario: %v", err)
	}

	if user.GetID() != 1 {
		t.Errorf("ID incorrecto")
	}

	if user.GetName() != "Diego Hachig" {
		t.Errorf("Nombre incorrecto")
	}
}
