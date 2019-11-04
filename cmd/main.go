package main

import (
	godotenv "github.com/joho/godotenv"
	db "github.com/sergey-suslov/trechit-server/internal/db"
	routing "github.com/sergey-suslov/trechit-server/internal/routing"
	"log"
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
	routing.Init()
}
