package user_data

import (
	"errors"
	authentication_middleware "wabustock/pkg/middleware/authentication-middleware"
	"wabustock/pkg/utils/paseto-token"
)

func DecryptToken(maker *paseto_token.PasetoMaker) (*paseto_token.Payload, error) {
	payload := &paseto_token.Payload{}

	token, extractTokenError := authentication_middleware.ExtractPasetoTokenFromHeader()

	if extractTokenError != nil {
		return payload, errors.New(extractTokenError.Error())
	}

	err := maker.Paseto.Decrypt(token, maker.SymmetricKey, payload, nil)
	if err != nil {
		return payload, errors.New(err.Error())

	}

	err = payload.Valid()
	if err != nil {
		return payload, errors.New(err.Error())
	}

	return payload, nil
}

func GetUserId() (string, error) {
	payload, err := DecryptToken(paseto_token.TokenMaker)
	if err != nil {
		return "", err
	}
	return payload.UserId, nil
}
