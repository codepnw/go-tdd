package main

import (
	"log"

	"github.com/codepnw/go-tdd/cmd/server"
	"github.com/codepnw/go-tdd/database"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	dbConn := database.ConnectDB()
	defer dbConn.Close()

	server.NewServer(dbConn.GetDB()).Start()
}
