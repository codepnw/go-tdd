package user

import (
	"database/sql"
	"errors"

	"github.com/codepnw/go-tdd/models"
	"github.com/codepnw/go-tdd/utils/password"
)

type User interface {
	CreateUser(user *User) (*models.User, error)
	DeleteUserByID(id int64) error
	FindUserByEmail(email string) (*models.User, error)
	FindUserByID(email string) (*models.User, error)
}

type userHandler struct {
	db *sql.DB
}

func NewUserHandler(db *sql.DB) *userHandler {
	return &userHandler{db: db}
}

var (
	sqlCreateUser = `INSERT INTO users (name, email, password)
		VALUES ($1, $2, $3) RETURNING id, name, email, password;`
	sqlDeleteUserByID = `DELETE FROM users WHERE id = $1;`
	sqlFindUserByEmail = `SELECT id, name, email, password FROM users WHERE email = $1;`
	sqlFindUserByID = `SELECT id, name, email, password FROM users WHERE id = $1;`

	ErrUserNotFound = errors.New("user not found")
)

func (u *userHandler) CreateUser(user *models.User) (*models.User, error) {
	user.Password, _ = password.PasswordHash(user.Password)
	err := u.db.QueryRow(sqlCreateUser, user.Name, user.Email, user.Password).Scan(&user.ID, &user.Name, &user.Email, &user.Password)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userHandler) DeleteUserByID(id int64) error {
	_, err := u.db.Exec(sqlDeleteUserByID, id)
	if err != nil {
		return err
	}

	return nil
}

func (u *userHandler) FindUserByEmail(email string) (*models.User, error) {
	user := &models.User{}

	err := u.db.QueryRow(sqlFindUserByEmail, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return user, nil
}

func (u *userHandler) FindUserByID(id int64) (*models.User, error) {
	user := &models.User{}

	err := u.db.QueryRow(sqlFindUserByID, id).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return user, nil
}
