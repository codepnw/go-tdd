package user

import (
	"time"

	"github.com/codepnw/go-tdd/utils/password"
)

type usersUsecase struct {
	repo    IUsersRepository
	timeout time.Duration
}

func NewUsersUsecase(repo IUsersRepository) IUsersUsecase {
	return &usersUsecase{
		repo,
		time.Duration(2) * time.Second,
	}
}

func (u *usersUsecase) CreateUser(req *CreateUserReq) (*CreateUserRes, error) {
	hashedPassword, err := password.PasswordHash(req.Password)
	if err != nil {
		return nil, err
	}

	user := &User{
		Email:    req.Email,
		Password: hashedPassword,
		Name:     req.Name,
	}

	r, err := u.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	res := &CreateUserRes{
		ID:    r.ID,
		Email: r.Email,
		Name:  r.Name,
	}

	return res, nil
}

func (u *usersUsecase) DeleteUserByID(id int64) error {
	if err := u.repo.DeleteUserByID(id); err != nil {
		return err
	}
	return nil
}

func (u *usersUsecase) FindUserByEmail(email string) (*User, error) {
	user, err := u.repo.FindUserByEmail(email) 
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *usersUsecase) FindUserByID(id int64) (*User, error) {
	user, err := u.repo.FindUserByID(id) 
	if err != nil {
		return nil, err
	}
	return user, nil
}

