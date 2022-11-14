package user

import (
	"context"
	"errors"
	"strconv"

	"github.com/ahmadlubis/lavandeapp/module/backend/entity"
	"github.com/ahmadlubis/lavandeapp/module/backend/model"
	"github.com/ahmadlubis/lavandeapp/module/backend/usecase"
	"gorm.io/gorm"
)

type superadminUsecase struct {
	db *gorm.DB
}

func NewSuperAdminUsecase(db *gorm.DB) usecase.SuperadminUsecase {
	return &superadminUsecase{db: db}
}

func (s *superadminUsecase) SetUserAsSuperadmin(ctx context.Context, userID int) error {
	var user entity.User
	if err := s.db.Where("id = ?", userID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.UserNotFoundError
		}
		return model.NewUnknownError(strconv.FormatUint(uint64(userID), 10), err)
	}

	user.Role = entity.UserRoleAdmin
	result := s.db.Save(&user)

	return result.Error
}
