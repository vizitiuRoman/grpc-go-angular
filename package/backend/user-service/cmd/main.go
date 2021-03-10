package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/user-service/pkg/server"
	"github.com/user-service/pkg/store"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	s, err := store.NewStore()
	if err != nil {
		log.Fatal(err)
	}
	srv := server.NewServer(os.Getenv("PORT"))
	srv.StartGRPC(s)
}
