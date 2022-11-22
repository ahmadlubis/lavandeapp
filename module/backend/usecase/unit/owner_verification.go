package unit

import (
	"context"
	"errors"
	"github.com/ahmadlubis/lavandeapp/module/backend/entity"
	"github.com/ahmadlubis/lavandeapp/module/backend/model"
	"github.com/ahmadlubis/lavandeapp/module/backend/usecase"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type unitOwnerVerificationUsecase struct {
	db *gorm.DB
}

func NewUnitOwnerVerificationUsecase(db *gorm.DB) usecase.UnitOwnerVerificationUsecase {
	return &unitOwnerVerificationUsecase{db: db}
}

// VerifyOwner raise error if userID is NOT the owner of unitID
func (u *unitOwnerVerificationUsecase) VerifyOwner(_ context.Context, unitID, userID uint64) error {
	var tenant entity.Tenant
	if err := u.db.Where("unit_id = ? AND user_id = ? AND role = ? AND (end_at IS NULL OR end_at > ?)", unitID, userID, entity.TenantRoleOwner, time.Now()).First(&tenant).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.ForbiddenError
		}
		return model.NewUnknownError(strconv.FormatUint(userID, 10), err)
	}
	return nil
}

// VerifyAnyOwner raise error if userID is NOT the owner of ANY unit
func (u *unitOwnerVerificationUsecase) VerifyAnyOwner(_ context.Context, userID uint64) error {
	var tenant entity.Tenant
	if err := u.db.Where("user_id = ? AND role = ? AND (end_at IS NULL OR end_at > ?)", userID, entity.TenantRoleOwner, time.Now()).First(&tenant).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.ForbiddenError
		}
		return model.NewUnknownError(strconv.FormatUint(userID, 10), err)
	}
	return nil
}
