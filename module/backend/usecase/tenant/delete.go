package tenant

import (
	"context"
	"errors"
	"github.com/ahmadlubis/lavandeapp/module/backend/entity"
	"github.com/ahmadlubis/lavandeapp/module/backend/model"
	"github.com/ahmadlubis/lavandeapp/module/backend/model/request"
	"github.com/ahmadlubis/lavandeapp/module/backend/usecase"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

type tenantDeletionUsecase struct {
	db *gorm.DB
}

func NewTenantDeletionUsecase(db *gorm.DB) usecase.TenantDeletionUsecase {
	return &tenantDeletionUsecase{db: db}
}

func (u *tenantDeletionUsecase) Delete(ctx context.Context, req request.DeleteTenantRequest) error {
	var tenant entity.Tenant
	now := time.Now()

	// If the target unit and user already has no active relationship, raise error
	if err := u.db.Where("unit_id = ? AND user_id = ? AND (end_at IS NULL OR end_at > ?)", req.UnitID, req.UserID, now).First(&tenant).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.NewExpectedError("target tenant not found", "TENANT_NOT_FOUND", http.StatusNotFound, "")
		}
		return model.NewUnknownError(strconv.FormatUint(req.UserID, 10), err)
	}

	tenant.EndAt = &now
	if result := u.db.Save(&tenant); result.Error != nil {
		return model.NewUnknownError("", result.Error)
	}

	return nil
}
