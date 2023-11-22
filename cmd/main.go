package main

import (
	"log"

	"github.com/codepnw/go-tdd/database"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	_, err = database.ConnectDB()
	if err != nil {
		log.Fatal("connection database failed: ", err)
	}

	log.Println("Hello..")
}