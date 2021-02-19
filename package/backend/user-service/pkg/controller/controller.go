package controller

import (
	"context"

	pb "github.com/user-service/grpc-proto/user"
	. "github.com/user-service/pkg/domain"
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

func (ctrl *UserController) CreateUser(ctx context.Context, req *pb.CreateUserReq) (*pb.UserRes, error) {
	err := req.Validate()
	if err != nil {
		ctrl.logger.Errorf("[ctrl.CreateUser] error validate user: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	createdUser, err := ctrl.services.User.CreateUser(&User{
		Email: req.Email, Password: req.Password,
	})
	if err != nil {
		ctrl.logger.Errorf("[ctrl.CreateUser] error create user: %v", err)
		return nil, status.Errorf(codes.AlreadyExists, err.Error())
	}

	return &pb.UserRes{Id: createdUser.ID, Email: createdUser.Email}, nil
}

func (ctrl *UserController) UpdateUser(ctx context.Context, req *pb.UpdateUserReq) (*pb.UserRes, error) {
	err := req.Validate()
	if err != nil {
		ctrl.logger.Errorf("[ctrl.UpdateUser] error validate user: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	updatedUser, err := ctrl.services.User.UpdateUser(&User{
		ID: req.Id, Email: req.Email, Password: req.Password,
	})
	if err != nil {
		ctrl.logger.Errorf("[ctrl.UpdateUser] error update user: %v", err)
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.UserRes{Id: updatedUser.ID, Email: updatedUser.Email}, nil
}

func (ctrl *UserController) DeleteUser(ctx context.Context, req *pb.UserReq) (*pb.Stub, error) {
	err := ctrl.services.User.DeleteUser(req.Id)
	if err != nil {
		ctrl.logger.Errorf("[ctrl.DeleteUser] error delete by id: %v", err)
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.Stub{}, nil
}

func (ctrl *UserController) VerifyUser(ctx context.Context, req *pb.VerifyUserReq) (*pb.UserRes, error) {
	err := req.Validate()
	if err != nil {
		ctrl.logger.Errorf("[ctrl.VerifyUser] error validate user: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	verifiedUser, err := ctrl.services.User.VerifyUser(req.Email, req.Password)
	if err != nil {
		ctrl.logger.Errorf("[ctrl.VerifyUser] error verify password: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, "Invalid email or password")
	}

	return &pb.UserRes{Id: verifiedUser.ID, Email: verifiedUser.Email}, nil
}

func (ctrl *UserController) GetUser(ctx context.Context, req *pb.UserReq) (*pb.UserRes, error) {
	foundUser, err := ctrl.services.User.GetUser(req.Id)
	if err != nil {
		ctrl.logger.Errorf("[ctrl.GetUser] error find user by id: %v", err)
		return nil, status.Errorf(codes.NotFound, "Not found user")
	}
	return &pb.UserRes{Id: foundUser.ID, Email: foundUser.Email}, nil
}

func (ctrl *UserController) GetUsers(ctx context.Context, stub *pb.Stub) (*pb.UsersRes, error) {
	return &pb.UsersRes{}, nil
}
