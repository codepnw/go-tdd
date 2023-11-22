package user

import (
	"errors"
	"testing"

	"github.com/codepnw/go-tdd/models"
)

var (
	oldSqlCreateUser = sqlCreateUser
	oldSqlDeleteUserByID = sqlDeleteUserByID
	oldSqlFindUserByEmail = sqlFindUserByEmail
	oldSqlFindUserByID = sqlFindUserByID
)

func TestCreateUser(t *testing.T) {
	h := NewUserHandler(testDB)
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

	sqlCreateUser = "invalid"
	_, err = h.CreateUser(user)
	if err == nil {
		t.Errorf("expected error not to be nil for invalid CreateUser sql")
	}
	sqlCreateUser = oldSqlCreateUser

	err = h.DeleteUserByID(createdUser.ID)
	if err != nil {
		t.Errorf("expected nil error during DeleteUserByID; got %q", err)
	}

	sqlDeleteUserByID = "invalid"
	err = h.DeleteUserByID(createdUser.ID)
	if err == nil {
		t.Errorf("expected error not to be nil for invalid DeleteUserByID sql")
	}
	sqlDeleteUserByID = oldSqlDeleteUserByID
}

func TestFindUserByEmail(t *testing.T) {
	h := NewUserHandler(testDB)

	user := &models.User{
		Email:    "testaa@mail.com",
		Password: "password",
		Name:     "John Cena",
	}

	createdUser, err := h.CreateUser(user)
	if err != nil {
		t.Fatal(err)
	}

	uByEmail, err := h.FindUserByEmail(createdUser.Email)
	if err != nil {
		t.Errorf("expected no error; got %q", err)
	}

	if uByEmail.Email != createdUser.Email {
		t.Errorf("expected %q; got %q", createdUser.Email, uByEmail.Email)
	}

	_, err = h.FindUserByEmail("invalid@mail.com")
	if err == nil {
		t.Errorf("want error; got nil for invalid email")
	}

	if err != nil && !errors.Is(err, ErrUserNotFound) {
		t.Errorf("want ErrUserNotFound error; got %q", err)
	}

	sqlFindUserByEmail = "invalid"
	_, err = h.FindUserByEmail(createdUser.Email)
	if err == nil {
		t.Errorf("want error; got nil for invalid FindUserByEmail %q", err)
	}
	sqlFindUserByEmail = oldSqlFindUserByEmail

	_ = h.DeleteUserByID(createdUser.ID)
}

func TestFindUserByID(t *testing.T) {
	h := NewUserHandler(testDB)

	user := &models.User{
		Email:    "testaa@mail.com",
		Password: "password",
		Name:     "John Cena",
	}

	createdUser, err := h.CreateUser(user)
	if err != nil {
		t.Fatal(err)
	}

	uByID, err := h.FindUserByID(createdUser.ID)
	if err != nil {
		t.Errorf("expected no error; got %q", err)
	}

	if uByID.ID != createdUser.ID {
		t.Errorf("expected %q; got %q", createdUser.ID, uByID.ID)
	}

	_, err = h.FindUserByID(-1)
	if err == nil {
		t.Errorf("want error; got nil for invalid id")
	}

	if err != nil && !errors.Is(err, ErrUserNotFound) {
		t.Errorf("want ErrUserNotFound error; got %q", err)
	}

	sqlFindUserByID = "invalid"
	_, err = h.FindUserByID(createdUser.ID)
	if err == nil {
		t.Errorf("want error; got nil for invalid FindUserByID %q", err)
	}
	sqlFindUserByID = oldSqlFindUserByID

	_ = h.DeleteUserByID(createdUser.ID)
}