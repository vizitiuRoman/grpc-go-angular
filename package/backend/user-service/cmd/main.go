package main

import (
	"github.com/user-service/pkg/server"
	"github.com/user-service/pkg/settings"
)

func main() {
	settings.Init()
	srv := server.NewServer()
	srv.Init()
	srv.StartGRPC()
}
