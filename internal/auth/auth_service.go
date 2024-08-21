package auth

import (
	"errors"
	"os"
	"time"
	"wabustock/internal/user"
	"wabustock/pkg/utils"
	"wabustock/pkg/utils/paseto-token"
)

func LoginService(authRequest AuthRequest) (AuthResponse, error) {
	switch authRequest.UserType {
	case "customer", "driver", "staff":
		return loginUser(authRequest)
	default:
		return AuthResponse{}, errors.New("invalid user_type_constants type")
	}
}

func loginUser(authRequest AuthRequest) (AuthResponse, error) {
	userDetails, err := user.FindUserByPhoneNumberRepo(authRequest.PhoneNumber)
	if err != nil {
		return AuthResponse{}, err
	}

	err = utils.VerifyPassword(*userDetails.Password, authRequest.Password)
	if err != nil {
		return AuthResponse{}, err
	}

	//paseto-token, err := createToken(string(userDetails.ID.String()))
	token, err := paseto_token.TokenMaker.CreateToken((userDetails.ID.String()), time.Hour)
	if err != nil {
		return AuthResponse{}, err
	}

	//return AuthResponse{
	//	Token:       paseto-token,
	//	Role:        *userDetails.Role,
	//	PhoneNumber: *userDetails.PhoneNumber,
	//}, nil
	return AuthResponse{
		Token:       token,
		Role:        "",
		PhoneNumber: *userDetails.PhoneNumber,
	}, nil
}

func createToken(userID string) (string, error) {
	privateKey := os.Getenv("ACCESS_TOKEN_PRIVATE_KEY")
	expireTime := os.Getenv("ACCESS_TOKEN_EXPIRED_IN")
	expireDuration, err := time.ParseDuration(expireTime)
	if err != nil {
		return "", err
	}

	return utils.CreateToken(expireDuration, userID, privateKey)
}
