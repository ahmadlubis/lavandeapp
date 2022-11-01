package usecase

import (
	"github.com/ahmadlubis/lavandeapp/module/backend/entity"
	"github.com/ahmadlubis/lavandeapp/module/backend/model/request"
)

type UserRegistrationUsecase interface {
	RegisterUser(request request.RegisterUserRequest) (entity.User, error)
}
