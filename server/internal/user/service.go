package user

import (
	"context"
	"errors"

	"github.com/nathanmintegui/edina/internal/auth"
	"github.com/nathanmintegui/edina/internal/user/dto"
)

type Service struct {
	Repository Repository
}

func (s Service) Login(ctx context.Context, request dto.LoginRequest) (string, error) {
	user, err := service.Repository.FindUserByLogin(ctx, request.Login)
	if err != nil {
		return "", err
	}

	if user.Id == 0 {
		return "", errors.New("user is not registered")
	}

	// NOTE: may add password max try 20 times
	isPasswordCorrect := auth.ComparePasswords(user.Password, []byte(request.Password))

	if !isPasswordCorrect {
		return "", errors.New("wrong password")
	}

	// TODO: retrieve secret from .env
	token, err := auth.CreateJWT([]byte("not so secret anymore, i'nit?"), user.Id)
	if err != nil {
		return "", err
	}

	return token, nil
}
