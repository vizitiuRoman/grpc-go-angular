package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	pb "github.com/user-service/grpc-proto/user"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	var conn *grpc.ClientConn
	conn, err = grpc.Dial("localhost:"+port, grpc.WithInsecure())
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
