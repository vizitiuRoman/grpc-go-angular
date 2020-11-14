package main

import (
	"github.com/auth-service/pkg/server"
	"github.com/auth-service/pkg/settings"
)

func main() {
	settings.Init()
	srv := server.NewServer()
	srv.Init()
	srv.StartGRPC()
}
