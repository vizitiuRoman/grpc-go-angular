package auth

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	. "github.com/auth-service/pkg/models"
	. "github.com/auth-service/pkg/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
)

type (
	createToken struct {
		AToken string
		RToken string
	}
	rtDetails struct {
		RtUUID string
		UserID uint64
	}
	atDetails struct {
		AtUUID string
		UserID uint64
	}
)

var secret = "feqwe"

func prepareToken(extractedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(extractedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return &jwt.Token{}, err
	}
	return token, nil
}

func CreateToken(ctx context.Context, userID uint64) (*createToken, error) {
	atUUID := uuid.NewV4().String()
	atClaims := jwt.MapClaims{}
	atClaims[UserID] = userID
	atClaims[AtUUID] = atUUID
	atClaims["exp"] = time.Now().Add(AtExpires).Unix()
	aToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims).SignedString([]byte(secret))
	if err != nil {
		return &createToken{}, err
	}

	rtUUID := atUUID + "++" + strconv.Itoa(int(userID))
	rtClaims := jwt.MapClaims{}
	rtClaims[UserID] = userID
	rtClaims[RtUUID] = rtUUID
	rtClaims["exp"] = time.Now().Add(RtExpires).Unix()
	rToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims).SignedString([]byte(secret))
	if err != nil {
		return &createToken{}, err
	}

	td := &TokenDetails{
		AToken:    aToken,
		RToken:    rToken,
		AtUUID:    atUUID,
		RtUUID:    rtUUID,
		AtExpires: time.Now().Add(AtExpires).Unix(),
		RtExpires: time.Now().Add(RtExpires).Unix(),
	}
	err = td.Create(ctx, userID)
	if err != nil {
		return &createToken{}, err
	}
	return &createToken{aToken, rToken}, nil
}

func ExtractAtMetadata(extractedToken string) (*atDetails, error) {
	token, err := prepareToken(extractedToken)
	if err != nil {
		return &atDetails{}, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		atUUID, ok := claims[AtUUID].(string)
		if !ok {
			return &atDetails{}, errors.New("Cannot get access uuid")
		}
		userID, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims[UserID]), 10, 32)
		if err != nil {
			return &atDetails{}, errors.New("Cannot get user id")
		}
		return &atDetails{
			AtUUID: atUUID,
			UserID: userID,
		}, nil
	}
	return &atDetails{}, errors.New("ExtractAtMetadata error")
}

func ExtractRtMetadata(rToken string) (*rtDetails, error) {
	token, err := prepareToken(rToken)
	if err != nil {
		return &rtDetails{}, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		rtUUID, ok := claims[RtUUID].(string)
		if !ok {
			return &rtDetails{}, errors.New("Cannot get refresh uuid")
		}
		userID, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims[UserID]), 10, 32)
		if err != nil {
			return &rtDetails{}, errors.New("Cannot get user id")
		}
		return &rtDetails{
			RtUUID: rtUUID,
			UserID: userID,
		}, nil
	}
	return &rtDetails{}, errors.New("ExtractRtMetadata error")
}

func GetRTokenUUID(aUUID string, userID uint64) string {
	return fmt.Sprintf("%s++%d", aUUID, userID)
}
