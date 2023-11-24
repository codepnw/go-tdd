package server

import (
	"github.com/codepnw/go-tdd/modules/user"
)

type IUsersModule interface {
	Init()
	Repository() user.IUsersRepository
	Usecase() user.IUsersUsecase
	Handler() user.IUsersHandler
}

type usersModule struct {
	*moduleFactory
	repo    user.IUsersRepository
	usecase user.IUsersUsecase
	handler user.IUsersHandler
}

func (m *moduleFactory) UsersModule() IUsersModule {
	uRepo := user.NewUsersRepository(m.s.db)
	uUsecase := user.NewUsersUsecase(uRepo)
	uHandler := user.NewUsersHandler(uUsecase)

	return &usersModule{
		moduleFactory: m,
		repo:          uRepo,
		usecase:       uUsecase,
		handler:       uHandler,
	}
}

func (u *usersModule) Init() {
	router := u.r.Group("/users")
	router.Post("/create", u.handler.CreateUser)
}

func (u *usersModule) Repository() user.IUsersRepository { return u.repo }
func (u *usersModule) Usecase() user.IUsersUsecase       { return u.usecase }
func (u *usersModule) Handler() user.IUsersHandler       { return u.handler }
