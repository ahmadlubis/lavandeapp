package usecase

import (
	"context"

	"github.com/ahmadlubis/lavandeapp/module/backend/entity"
	"github.com/ahmadlubis/lavandeapp/module/backend/model"
	"github.com/ahmadlubis/lavandeapp/module/backend/model/request"
)

type UserRegistrationUsecase interface {
	Register(ctx context.Context, request request.RegisterUserRequest) (entity.User, error)
}

type UserLoginUsecase interface {
	Login(ctx context.Context, request request.LoginUserRequest) (model.AccessToken, error)
}

type UserTokenVerificationUsecase interface {
	VerifyToken(ctx context.Context, token string) (entity.User, error)
}

type UserSelfUpdateUsecase interface {
	SelfUpdate(ctx context.Context, request request.SelfUpdateUserRequest) (entity.User, error)
}

type SuperadminUsecase interface {
	SetUserAsSuperadmin(ctx context.Context, userID int) error
}
