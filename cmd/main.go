package main

import (
	"log"

	"github.com/codepnw/go-tdd/cmd/server"
	"github.com/codepnw/go-tdd/database"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	dbConn := database.ConnectDB()
	dbConn.Close()

	server.NewServer(dbConn.GetDB()).Start()
}
