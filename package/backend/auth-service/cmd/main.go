package main

import (
	"log"
	"os"

	"github.com/auth-service/pkg/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	srv := server.NewServer(os.Getenv("PORT"))
	srv.Init()
	srv.StartGRPC()
}
