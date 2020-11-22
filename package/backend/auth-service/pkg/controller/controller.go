package controller

import (
	"context"
	"strings"

	pb "github.com/auth-service/grpc-proto/auth"
	userpb "github.com/auth-service/grpc-proto/user"
	"github.com/auth-service/pkg/auth"
	. "github.com/auth-service/pkg/models"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Controller struct {
	userAddr string
	logger   *zap.SugaredLogger
}

func NewController(userAddr string, logger *zap.SugaredLogger) *Controller {
	return &Controller{userAddr, logger}
}

func (c *Controller) Auth(ctx context.Context, req *pb.AuthReq) (*pb.AuthRes, error) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(c.userAddr, grpc.WithInsecure())
	if err != nil {
		c.logger.Errorf("Auth error dial connection: %v", err)
	}
	defer conn.Close()

	client := userpb.NewUserServiceClient(conn)

	userReq := &userpb.VerifyUserReq{Email: req.Email, Password: req.Password}

	response, err := client.VerifyUser(ctx, userReq)
	if err != nil {
		c.logger.Errorf("Auth error verify user: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	token, err := auth.CreateToken(ctx, response.Id)
	if err != nil {
		c.logger.Errorf("Auth error create token: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	return &pb.AuthRes{
		Token:        token.AToken,
		RefreshToken: token.RToken,
	}, nil
}

func (c *Controller) Register(ctx context.Context, req *pb.RegisterReq) (*pb.AuthRes, error) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(c.userAddr, grpc.WithInsecure())
	if err != nil {
		c.logger.Errorf("Register error dial connection: %v", err)
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	defer conn.Close()

	client := userpb.NewUserServiceClient(conn)

	userReq := &userpb.CreateUserReq{Email: req.Email, Password: req.Password}

	response, err := client.CreateUser(ctx, userReq)
	if err != nil {
		c.logger.Errorf("Register error create user: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	token, err := auth.CreateToken(ctx, response.Id)
	if err != nil {
		c.logger.Errorf("Register error create token: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	return &pb.AuthRes{
		Token:        token.AToken,
		RefreshToken: token.RToken,
	}, nil
}

func (c *Controller) UpdateAuth(ctx context.Context, req *pb.UpdateAuthReq) (*pb.UpdateAuthRes, error) {
	token, err := extractMetadata(ctx, "authorization")
	if err != nil {
		c.logger.Errorf("UpdateAuth error extract metadata: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	extractedAToken, err := auth.ExtractAtMetadata(token)
	if err != nil {
		c.logger.Errorf("Error extract access token metadata: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	rtUUID := auth.GetRTokenUUID(extractedAToken.AtUUID, extractedAToken.UserID)

	extractedRToken, err := auth.ExtractRtMetadata(req.RefreshToken)
	if err != nil {
		c.logger.Errorf("UpdateAuth error extract refresh token metadata: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	if extractedRToken.RtUUID != rtUUID {
		c.logger.Error("UpdateAuth error extract refresh token not equal rt uuid")
		return nil, status.Errorf(codes.InvalidArgument, "Invalid argument")
	}

	var td TokenDetails
	err = td.DeleteByUUID(ctx, extractedAToken.AtUUID, rtUUID)
	if err != nil {
		c.logger.Errorf("UpdateAuth error delete token by uuid: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	createdToken, err := auth.CreateToken(ctx, extractedAToken.UserID)
	if err != nil {
		c.logger.Errorf("UpdateAuth error create token: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	return &pb.UpdateAuthRes{
		Token:        createdToken.AToken,
		RefreshToken: createdToken.RToken,
	}, nil
}

func (c *Controller) Logout(ctx context.Context, req *pb.Stub) (*pb.Stub, error) {
	token, err := extractMetadata(ctx, "authorization")
	if err != nil {
		c.logger.Errorf("Logout error extract metadata: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	extractedAToken, err := auth.ExtractAtMetadata(token)
	if err != nil {
		c.logger.Errorf("Logout error extract access token metadata: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	rtUUID := auth.GetRTokenUUID(extractedAToken.AtUUID, extractedAToken.UserID)

	var td TokenDetails
	err = td.DeleteByUUID(ctx, extractedAToken.AtUUID, rtUUID)
	if err != nil {
		c.logger.Errorf("Logout error delete token by uuid: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	return &pb.Stub{}, nil
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
