package usecase

import (
	"context"
	"github.com/ahmadlubis/lavandeapp/module/backend/entity"
	"github.com/ahmadlubis/lavandeapp/module/backend/model/request"
)

type UserRegistrationUsecase interface {
	RegisterUser(ctx context.Context, request request.RegisterUserRequest) (entity.User, error)
}
