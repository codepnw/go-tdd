package user

import (
	"database/sql"
	"errors"
)

type usersRepository struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) IUsersRepository {
	return &usersRepository{
		db: db,
	}
}

var (
	sqlCreateUser = `INSERT INTO users (name, email, password)
		VALUES ($1, $2, $3) RETURNING id, name, email, password;`
	sqlDeleteUserByID  = `DELETE FROM users WHERE id = $1;`
	sqlFindUserByEmail = `SELECT id, name, email, password FROM users WHERE email = $1;`
	sqlFindUserByID    = `SELECT id, name, email, password FROM users WHERE id = $1;`

	ErrUserNotFound = errors.New("user not found")
)

func (u *usersRepository) CreateUser(user *User) (*User, error) {
	err := u.db.QueryRow(sqlCreateUser, user.Name, user.Email, user.Password).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *usersRepository) DeleteUserByID(id int64) error {
	_, err := u.db.Exec(sqlDeleteUserByID, id)
	if err != nil {
		return err
	}
	return nil
}

func (u *usersRepository) FindUserByEmail(email string) (*User, error) {
	user := &User{}

	err := u.db.QueryRow(sqlFindUserByEmail, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return user, nil
}

func (u *usersRepository) FindUserByID(id int64) (*User, error) {
	user := &User{}

	err := u.db.QueryRow(sqlFindUserByID, id).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return user, nil
}
