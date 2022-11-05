package usecase

import (
	"context"
	"github.com/ahmadlubis/lavandeapp/module/backend/entity"
	"github.com/ahmadlubis/lavandeapp/module/backend/model/request"
	"github.com/ahmadlubis/lavandeapp/module/backend/model/response"
)

type UserRegistrationUsecase interface {
	RegisterUser(ctx context.Context, request request.RegisterUserRequest) (entity.User, error)
}

type UserLoginUsecase interface {
	Login(ctx context.Context, request request.LoginUserRequest) (response.UserAccessTokenResponse, error)
}
