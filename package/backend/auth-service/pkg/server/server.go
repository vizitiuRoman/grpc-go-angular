package server

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "github.com/auth-service/grpc-proto/auth"
	"github.com/auth-service/pkg/controllers"
	"github.com/auth-service/pkg/models"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Server struct {
	logger    *zap.SugaredLogger
	port      string
	interrupt chan os.Signal
	listen    chan error
}

func NewServer() *Server {
	return &Server{
		port:      "9090",
		logger:    newLogger(),
		interrupt: make(chan os.Signal, 1),
		listen:    make(chan error, 1),
	}
}

func (srv *Server) Init() {
	err := models.InitRedis()
	if err != nil {
		srv.logger.Fatalf("Error on init redis: %v", err)
	}
}

func (srv *Server) StartGRPC() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	listener, err := net.Listen("tcp", ":"+srv.port)
	if err != nil {
		srv.logger.Fatalf("Listen error: %v", err)
	}

	gRPCServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(gRPCServer, controllers.NewController(srv.logger))

	go func(listen chan error) {
		srv.logger.Info("Service started on port: " + srv.port)
		listen <- gRPCServer.Serve(listener)
	}(srv.listen)

	signal.Notify(srv.interrupt, syscall.SIGINT, syscall.SIGTERM)
	srv.waitShutdown(srv.listen, srv.interrupt)
}

func (srv Server) waitShutdown(listen chan error, interrupt chan os.Signal) {
	for {
		select {
		case err := <-listen:
			if err != nil {
				srv.logger.Fatalf("Listener error: %v", err)
			}
			os.Exit(0)
		case err := <-interrupt:
			srv.logger.Fatalf("Shutdown signal: %v", err.String())
		}
	}
}
