package auth

import (
	"errors"
	"os"
	"time"
	"wabustock/internal/user"
	"wabustock/pkg/utils"
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

	err = utils.VerifyPassword(userDetails.Password, authRequest.Password)
	if err != nil {
		return AuthResponse{}, err
	}

	token, err := createToken(string(userDetails.ID.String()))
	if err != nil {
		return AuthResponse{}, err
	}

	return AuthResponse{
		Token:       token,
		Role:        userDetails.Role,
		PhoneNumber: userDetails.PhoneNumber,
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
