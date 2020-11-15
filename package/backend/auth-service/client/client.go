package main

import (
	"context"
	"log"

	pb "github.com/auth-service/grpc-proto/auth"
	"github.com/auth-service/pkg/config"
	"google.golang.org/grpc"
)

func main() {
	config.Init()
	port := config.Get().Port

	var conn *grpc.ClientConn
	conn, err := grpc.Dial("localhost"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	client := pb.NewAuthServiceClient(conn)

	req := &pb.AuthReq{Email: "roma", Password: "qwewq"}

	response, err := client.Auth(context.Background(), req)
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server aToken: %s", response.Token)
	log.Printf("Response from server rToken: %s", response.RefreshToken)
}
