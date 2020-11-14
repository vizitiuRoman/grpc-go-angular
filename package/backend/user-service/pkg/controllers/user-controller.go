package controllers

import (
	"context"
	"strings"

	pb "github.com/user-service/grpc-proto/user"
	. "github.com/user-service/pkg/models"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Controller struct {
	logger *zap.SugaredLogger
}

func NewController(logger *zap.SugaredLogger) *Controller {
	return &Controller{logger}
}

func (c *Controller) CreateUser(ctx context.Context, req *pb.CreateUserReq) (*pb.UserRes, error) {
	err := req.Validate()
	if err != nil {
		c.logger.Errorf("Error validate user: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	user := &User{Email: req.Email, Password: req.Password}
	createdUser, err := user.Create()
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.UserRes{Id: createdUser.ID, Email: createdUser.Email, Name: createdUser.Email}, nil
}

func (c *Controller) UpdateUser(ctx context.Context, req *pb.UpdateUserReq) (*pb.UserRes, error) {
	panic("implement me")
}

func (c *Controller) DeleteUser(ctx context.Context, req *pb.UserReq) (*pb.Stub, error) {
	panic("implement me")
}

func (c *Controller) VerifyUser(ctx context.Context, req *pb.VerifyUserReq) (*pb.UserRes, error) {
	panic("implement me")
}

func (c *Controller) GetUser(ctx context.Context, req *pb.UserReq) (*pb.UserRes, error) {
	panic("implement me")
}

func (c *Controller) GetUsers(ctx context.Context, stub *pb.Stub) (*pb.UsersRes, error) {
	panic("implement me")
}

func extractMetadata(ctx context.Context, key string) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "nil", status.Errorf(codes.DataLoss, "Failed to get metadata")
	}
	extractedMetadata := md[key]
	if len(extractedMetadata) == 0 {
		return "nil", status.Errorf(codes.InvalidArgument, "Missing authorization header")
	}
	if strings.Trim(extractedMetadata[0], " ") == "" {
		return "nil", status.Errorf(codes.InvalidArgument, "Empty authorization header")
	}
	return extractedMetadata[0], nil
}
