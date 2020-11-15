package main

import (
	"github.com/auth-service/pkg/server"
	"github.com/auth-service/pkg/config"
)

func main() {
	config.Init()
	srv := server.NewServer()
	srv.Init()
	srv.StartGRPC()
}
