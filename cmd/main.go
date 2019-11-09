package main

import (
	"log"

	godotenv "github.com/joho/godotenv"
	"github.com/sergey-suslov/trechit-server/internal/db"
	"github.com/sergey-suslov/trechit-server/internal/routing"
	"github.com/sergey-suslov/trechit-server/utils"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	_, err = db.Init()
	if err != nil {
		log.Println("Couldn't connect to DB")
	}
	utils.InitValidator()
	routing.Init()
}
