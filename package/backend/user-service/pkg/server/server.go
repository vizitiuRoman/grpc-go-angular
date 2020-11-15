package server

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "github.com/user-service/grpc-proto/user"
	"github.com/user-service/pkg/controllers"
	"github.com/user-service/pkg/models"
	"github.com/user-service/pkg/settings"
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
		port:      settings.Get().Port,
		logger:    newLogger(),
		interrupt: make(chan os.Signal, 1),
		listen:    make(chan error, 1),
	}
}

func (srv *Server) Init() {
	err := models.InitDatabase()
	if err != nil {
		srv.logger.Fatalf("Error on init db: %v", err)
	}
}

func (srv *Server) StartGRPC() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	listener, err := net.Listen("tcp", ":"+srv.port)
	if err != nil {
		srv.logger.Fatalf("Listen error: %v", err)
	}

	gRPCServer := grpc.NewServer()
	pb.RegisterUserServiceServer(gRPCServer, controllers.NewController(srv.logger))

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
