package server

import "github.com/gofiber/fiber/v2"

type IModuleFactory interface {
	MonitorModule()
	UsersModule() IUsersModule
}

type moduleFactory struct {
	r fiber.Router
	s *server
}

func InitModule(r fiber.Router, s *server) IModuleFactory {
	return &moduleFactory{
		r: r,
		s: s,
	}
}

func (m *moduleFactory) MonitorModule() {
	m.r.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON("OK")
	})
}

