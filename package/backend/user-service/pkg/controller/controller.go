package controller

import (
	"context"

	pb "github.com/user-service/grpc-proto/user"
	. "github.com/user-service/pkg/models"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
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
		c.logger.Errorf("CreateUser error validate user: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	user := &User{Email: req.Email, Password: req.Password}
	createdUser, err := user.Create()
	if err != nil {
		c.logger.Errorf("CreateUser error create user: %v", err)
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.UserRes{Id: createdUser.ID, Email: createdUser.Email}, nil
}

func (c *Controller) UpdateUser(ctx context.Context, req *pb.UpdateUserReq) (*pb.UserRes, error) {
	err := req.Validate()
	if err != nil {
		c.logger.Errorf("UpdateUser error validate user: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	user := User{ID: req.Id, Email: req.Email}
	err = user.Update()
	if err != nil {
		c.logger.Errorf("UpdateUser error update user: %v", err)
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.UserRes{Id: req.Id, Email: req.Email}, nil
}

func (c *Controller) DeleteUser(ctx context.Context, req *pb.UserReq) (*pb.Stub, error) {
	var user User
	err := user.DeleteByID(req.Id)
	if err != nil {
		c.logger.Errorf("DeleteUser error delete by id: %v", err)
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	return &pb.Stub{}, nil
}

func (c *Controller) VerifyUser(ctx context.Context, req *pb.VerifyUserReq) (*pb.UserRes, error) {
	err := req.Validate()
	if err != nil {
		c.logger.Errorf("VerifyUser error validate user: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	var user User
	foundUser, err := user.FindByEmail(req.Email)
	if err != nil {
		c.logger.Errorf("VerifyUser error find by email user: %v", err)
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	err = VerifyPassword(foundUser.Password, req.Password)
	if err != nil {
		c.logger.Errorf("VerifyUser error verify password: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, "Invalid email or password")
	}

	return &pb.UserRes{Id: foundUser.ID, Email: foundUser.Email}, nil
}

func (c *Controller) GetUser(ctx context.Context, req *pb.UserReq) (*pb.UserRes, error) {
	var user User
	foundUser, err := user.FindByID(req.Id)
	if err != nil {
		c.logger.Errorf("GetUser error find user by id: %v", err)
		return nil, status.Errorf(codes.NotFound, "Not found user")
	}
	return &pb.UserRes{Id: foundUser.ID, Email: foundUser.Email}, nil
}

func (c *Controller) GetUsers(ctx context.Context, stub *pb.Stub) (*pb.UsersRes, error) {
	return &pb.UsersRes{}, nil
}
