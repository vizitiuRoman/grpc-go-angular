package server

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "github.com/auth-service/grpc-proto/auth"
	"github.com/auth-service/pkg/controller"
	"github.com/auth-service/pkg/logger"
	"github.com/auth-service/pkg/models"
	"google.golang.org/grpc"
)

type server struct {
	port string
}

func NewServer(port string) *server {
	return &server{
		port: port,
	}
}

func (srv *server) Init() {
	err := models.InitRedis()
	if err != nil {
		log.Fatalf("Error on init redis: %v", err)
	}
}

func (srv *server) StartGRPC() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	listenCh := make(chan error, 1)
	interruptCh := make(chan os.Signal, 1)

	zapLogger := logger.NewLogger()


	listener, err := net.Listen("tcp", ":"+srv.port)
	if err != nil {
		zapLogger.Fatalf("Listen error: %v", err)
	}

	gRPCServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(
		gRPCServer,
		controller.NewController(os.Getenv("USER_SVC_ADDR"), zapLogger),
	)

	go func(listen chan error) {
		zapLogger.Info("Service started on port: " + srv.port)
		listen <- gRPCServer.Serve(listener)
	}(listenCh)

	signal.Notify(interruptCh, syscall.SIGINT, syscall.SIGTERM)
	waitShutdown(listenCh, interruptCh)
}

func waitShutdown(listen chan error, interrupt chan os.Signal) {
	for {
		select {
		case err := <-listen:
			if err != nil {
				log.Fatalf("Listener error: %v", err)
			}
			os.Exit(0)
		case err := <-interrupt:
			log.Fatalf("Shutdown signal: %v", err.String())
		}
	}
}

