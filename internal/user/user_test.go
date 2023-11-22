package user

import (
	"testing"

	"github.com/codepnw/go-tdd/models"
)

func TestCreateUser(t *testing.T) {
	h := NewUserHandler(nil)
	oldPassword := "password"

	user := &models.User{
		Email:    "test@mail.com",
		Password: oldPassword,
		Name:     "John Cena",
	}

	createdUser, err := h.CreateUser(user)
	if err != nil {
		t.Fatal(err)
	}

	if createdUser.ID == 0 {
		t.Errorf("want id not to be zero")
	}

	if user.Name != createdUser.Name {
		t.Errorf("expected %q: got %q", user.Name, createdUser.Name)
	}

	if createdUser.Password == oldPassword {
		t.Error("password was not hashed")
	}
}
