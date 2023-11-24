package server

import (
	"database/sql"
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
)

type IServer interface {
	Start()
}

type server struct {
	app *fiber.App
	db  *sql.DB
}

func NewServer(db *sql.DB) IServer {
	return &server{
		db:  db,
		app: fiber.New(),
	}
}

func (s *server) GetServer() *server {
	return s
}

func (s *server) Start() {
	v1 := s.app.Group("/v1")
	modules := InitModule(v1, s)

	modules.MonitorModule()
	modules.UsersModule().Init()

	// Graceful Shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		log.Println("server is shutting down....")
		_ = s.app.Shutdown()
	}()

	log.Printf("server is starting on %v", os.Getenv("PORT"))
	s.app.Listen(os.Getenv("PORT"))
}
