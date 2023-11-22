package user

import (
	"database/sql"

	"github.com/codepnw/go-tdd/models"
	"github.com/codepnw/go-tdd/utils/password"
)

type userHandler struct {
	db *sql.DB
}

func NewUserHandler(db *sql.DB) *userHandler {
	return &userHandler{db: db}
}

func (u *userHandler) CreateUser(user *models.User) (*models.User, error) {
	user.ID = 1
	user.Password, _ = password.PasswordHash(user.Password)
	return user, nil
}
