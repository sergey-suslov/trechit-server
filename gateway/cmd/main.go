package main

import (
	"log"

	godotenv "github.com/joho/godotenv"
	"github.com/sergey-suslov/trechit-server/gateway/internal/db"
	"github.com/sergey-suslov/trechit-server/gateway/internal/routing"
	"github.com/sergey-suslov/trechit-server/gateway/utils"
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
