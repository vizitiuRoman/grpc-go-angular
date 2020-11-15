package main

import (
	"github.com/user-service/pkg/config"
	"github.com/user-service/pkg/server"
)

func main() {
	config.Init()
	srv := server.NewServer()
	srv.Init()
	srv.StartGRPC()
}
