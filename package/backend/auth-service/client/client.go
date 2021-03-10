package main

import (
	"context"
	"log"
	"os"

	pb "github.com/auth-service/grpc-proto/auth"
	"google.golang.org/grpc"
)

func main() {
	port := os.Getenv("PORT")

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
