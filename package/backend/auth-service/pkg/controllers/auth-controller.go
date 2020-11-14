package controllers

import (
	"context"
	"strings"

	pb "github.com/auth-service/grpc-proto/auth"
	"github.com/auth-service/pkg/auth"
	. "github.com/auth-service/pkg/models"
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

func (c *Controller) Auth(ctx context.Context, req *pb.AuthReq) (*pb.AuthRes, error) {
	token, err := auth.CreateToken(context.Background(), 1)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	return &pb.AuthRes{
		UserId:       1,
		Token:        token.AToken,
		RefreshToken: token.RToken,
	}, nil
}

func (c *Controller) UpdateAuth(ctx context.Context, req *pb.UpdateAuthReq) (*pb.UpdateAuthRes, error) {
	token, err := extractMetadata(ctx, "authorization")
	if err != nil {
		c.logger.Errorf("Error extract metadata: %v", err)
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
		c.logger.Errorf("Error extract refresh token metadata: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	if extractedRToken.RtUUID != rtUUID {
		c.logger.Error("Error extract refresh token not equal rt uuid")
		return nil, status.Errorf(codes.InvalidArgument, "Invalid argument")
	}

	var td TokenDetails
	err = td.DeleteByUUID(ctx, extractedAToken.AtUUID, rtUUID)
	if err != nil {
		c.logger.Errorf("Error delete token by uuid: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	createdToken, err := auth.CreateToken(ctx, extractedAToken.UserID)
	if err != nil {
		c.logger.Errorf("Error create token: %v", err)
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
		c.logger.Errorf("Error extract metadata: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	extractedAToken, err := auth.ExtractAtMetadata(token)
	if err != nil {
		c.logger.Errorf("Error extract access token metadata: %v", err)
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	rtUUID := auth.GetRTokenUUID(extractedAToken.AtUUID, extractedAToken.UserID)

	var td TokenDetails
	err = td.DeleteByUUID(ctx, extractedAToken.AtUUID, rtUUID)
	if err != nil {
		c.logger.Errorf("Error delete token by uuid: %v", err)
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
