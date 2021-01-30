package main

import (
	"log"

	"github.com/user-service/pkg/config"
	"github.com/user-service/pkg/server"
	"github.com/user-service/pkg/store"
)

func main() {
	config.Init()
	s, err := store.NewStore()
	if err != nil {
		log.Fatal(err)
	}
	srv := server.NewServer(config.Get().Port)
	srv.StartGRPC(s)
}
