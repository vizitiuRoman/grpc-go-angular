package controller

import (
	"context"

	pb "github.com/user-service/grpc-proto/user"
	. "github.com/user-service/pkg/models"
	"github.com/user-service/pkg/services"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserController struct {
	services *services.Manager
	logger   *zap.SugaredLogger
}

func NewUserController(services *services.Manager, logger *zap.SugaredLogger) *UserController {
	return &UserController{services, logger}
}

func (ctr *UserController) CreateUser(ctx context.Context, req *pb.CreateUserReq) (*pb.UserRes, error) {
	err := req.Validate()
	if err != nil {
		ctr.logger.Errorf("ctr.CreateUser error validate user: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	user := &User{Email: req.Email, Password: req.Password}
	createdUser, err := user.Create()
	if err != nil {
		ctr.logger.Errorf("ctr.CreateUser error create user: %v", err)
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.UserRes{Id: createdUser.ID, Email: createdUser.Email}, nil
}

func (ctr *UserController) UpdateUser(ctx context.Context, req *pb.UpdateUserReq) (*pb.UserRes, error) {
	err := req.Validate()
	if err != nil {
		ctr.logger.Errorf("ctr.UpdateUser error validate user: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	user := User{ID: req.Id, Email: req.Email}
	err = user.Update()
	if err != nil {
		ctr.logger.Errorf("ctr.UpdateUser error update user: %v", err)
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.UserRes{Id: req.Id, Email: req.Email}, nil
}

func (ctr *UserController) DeleteUser(ctx context.Context, req *pb.UserReq) (*pb.Stub, error) {
	var user User
	err := user.DeleteByID(req.Id)
	if err != nil {
		ctr.logger.Errorf("ctr.DeleteUser error delete by id: %v", err)
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	return &pb.Stub{}, nil
}

func (ctr *UserController) VerifyUser(ctx context.Context, req *pb.VerifyUserReq) (*pb.UserRes, error) {
	err := req.Validate()
	if err != nil {
		ctr.logger.Errorf("ctr.VerifyUser error validate user: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	var user User
	foundUser, err := user.FindByEmail(req.Email)
	if err != nil {
		ctr.logger.Errorf("ctr.VerifyUser error find by email user: %v", err)
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	err = services.VerifyPassword(foundUser.Password, req.Password)
	if err != nil {
		ctr.logger.Errorf("ctr.VerifyUser error verify password: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, "Invalid email or password")
	}

	return &pb.UserRes{Id: foundUser.ID, Email: foundUser.Email}, nil
}

func (ctr *UserController) GetUser(ctx context.Context, req *pb.UserReq) (*pb.UserRes, error) {
	var user User
	foundUser, err := user.FindByID(req.Id)
	if err != nil {
		ctr.logger.Errorf("ctr.GetUser error find user by id: %v", err)
		return nil, status.Errorf(codes.NotFound, "Not found user")
	}
	return &pb.UserRes{Id: foundUser.ID, Email: foundUser.Email}, nil
}

func (ctr *UserController) GetUsers(ctx context.Context, stub *pb.Stub) (*pb.UsersRes, error) {
	return &pb.UsersRes{}, nil
}
