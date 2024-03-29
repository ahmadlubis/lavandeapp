package usecase

import (
	"context"

	"github.com/ahmadlubis/lavandeapp/module/backend/entity"
	"github.com/ahmadlubis/lavandeapp/module/backend/model"
	"github.com/ahmadlubis/lavandeapp/module/backend/model/request"
	"github.com/ahmadlubis/lavandeapp/module/backend/model/response"
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

type UserUpdateUsecase interface {
	Update(ctx context.Context, request request.UpdateUserRequest) (entity.User, error)
	AdminUpdate(ctx context.Context, request request.AdminUpdateUserRequest) (entity.User, error)
	ResetPassword(ctx context.Context, request request.UpdateUserRequest) (string, error)
}

type UserListUsecase interface {
	List(ctx context.Context, request request.ListUserRequest) (response.ListUserResponse, error)
}

type UnitCreationUsecase interface {
	Create(ctx context.Context, req request.CreateUnitRequest) (entity.Unit, error)
}

type UnitOwnerVerificationUsecase interface {
	VerifyOwner(ctx context.Context, unitID, userID uint64) error
	VerifyAnyOwner(ctx context.Context, userID uint64) error
}

type UnitUpdateUsecase interface {
	Update(ctx context.Context, request request.UpdateUnitRequest) (entity.Unit, error)
}

type UnitListUsecase interface {
	List(ctx context.Context, request request.ListUnitRequest) (response.ListUnitResponse, error)
}

type TenantCreationUsecase interface {
	Create(ctx context.Context, req request.CreateTenantRequest) (entity.Tenant, error)
}

type TenantDeletionUsecase interface {
	Delete(ctx context.Context, req request.DeleteTenantRequest) error
}

type TenantListUsecase interface {
	List(ctx context.Context, request request.ListTenantRequest) (response.ListTenantResponse, error)
}

type SuperadminUsecase interface {
	SetUserAsSuperadmin(ctx context.Context, userID int) error
}
