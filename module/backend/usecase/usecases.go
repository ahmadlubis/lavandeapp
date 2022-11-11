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

type UnitCreationUsecase interface {
	Create(ctx context.Context, req request.CreateUnitRequest) (entity.Unit, error)
}

type TenantCreationUsecase interface {
	Create(ctx context.Context, req request.CreateTenantRequest) (entity.Tenant, error)
}
