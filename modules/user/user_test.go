package user

import (
	"testing"
)

var (
	oldSqlCreateUser      = sqlCreateUser
	oldSqlDeleteUserByID  = sqlDeleteUserByID
	oldSqlFindUserByEmail = sqlFindUserByEmail
	oldSqlFindUserByID    = sqlFindUserByID
)

func TestCreateUser(t *testing.T) {

	repo := NewUsersRepository(testDB)
	h := NewUsersUsecase(repo)

	oldPassword := "password"

	user := &CreateUserReq{
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

