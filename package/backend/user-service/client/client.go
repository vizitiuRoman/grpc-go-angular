package main

import (
	"context"
	"log"

	pb "github.com/user-service/grpc-proto/user"
	"github.com/user-service/pkg/settings"
	"google.golang.org/grpc"
)

func main() {
	settings.Init()
	port := settings.Get().Port

	var conn *grpc.ClientConn
	conn, err := grpc.Dial("localhost:"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	req := &pb.CreateUserReq{Email: "roma@gmail.com", Password: "qweqwe"}

	response, err := client.CreateUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server aToken: %s", response.Email)
	log.Printf("Response from server aToken: %d", response.Id)
}
