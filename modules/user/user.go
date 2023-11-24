package user

import "github.com/gofiber/fiber/v2"

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type CreateUserReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type CreateUserRes struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type IUsersRepository interface {
	CreateUser(user *User) (*User, error)
	DeleteUserByID(id int64) error
	FindUserByEmail(email string) (*User, error)
	FindUserByID(id int64) (*User, error)
}

type IUsersUsecase interface {
	CreateUser(req *CreateUserReq) (*CreateUserRes, error)
	DeleteUserByID(id int64) error
	FindUserByEmail(email string) (*User, error)
	FindUserByID(id int64) (*User, error)
}

type IUsersHandler interface {
	CreateUser(c *fiber.Ctx) error
	FindUserByEmail(c *fiber.Ctx) error
	FindUserByID(c *fiber.Ctx) error
	DeleteUserByID(c *fiber.Ctx) error
}
