package main

import (
	"github.com/joho/godotenv"
	"github.com/sergey-suslov/trechit-server/api/timer"
	"github.com/sergey-suslov/trechit-server/timer/internal/api"
	"github.com/sergey-suslov/trechit-server/timer/internal/db"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
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

	lis, err := net.Listen("tcp", "localhost:"+os.Getenv("PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	timer.RegisterTimerServer(grpcServer, api.TimerService{})
	err = grpcServer.Serve(lis)

}
