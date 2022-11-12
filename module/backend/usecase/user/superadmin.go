package user

import (
	"context"
	"errors"
	"fmt"
	"net/http"

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

func (u *superadminUsecase) SetUserAsSuperadmin(ctx context.Context, userID int) error {
	if err := u.db.Where("id = ?", userID).Update("role", entity.UserRoleAdmin).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.NewExpectedError("userID not found", "USERID_NOT_FOUND",
				http.StatusBadRequest, fmt.Sprintf("%d", userID))
		}
		return model.NewUnknownError(fmt.Sprintf("%d", userID), err)
	}

	return nil
}
