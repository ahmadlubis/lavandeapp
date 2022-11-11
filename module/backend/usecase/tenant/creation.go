package tenant

import (
	"context"
	"errors"
	"github.com/ahmadlubis/lavandeapp/module/backend/entity"
	"github.com/ahmadlubis/lavandeapp/module/backend/model"
	"github.com/ahmadlubis/lavandeapp/module/backend/model/request"
	"github.com/ahmadlubis/lavandeapp/module/backend/usecase"
	"github.com/ahmadlubis/lavandeapp/module/backend/utility"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

type tenantCreationUsecase struct {
	db *gorm.DB
}

func NewTenantCreationUsecase(db *gorm.DB) usecase.TenantCreationUsecase {
	return &tenantCreationUsecase{db: db}
}

func (u *tenantCreationUsecase) Create(ctx context.Context, req request.CreateTenantRequest) (entity.Tenant, error) {
	role, err := entity.ParseTenantRole(req.Role)
	if err != nil {
		return entity.Tenant{}, err
	}

	unit, err := u.validateCreateTenantRequest(ctx, req, role)
	if err != nil {
		return entity.Tenant{}, err
	}

	tenant := entity.Tenant{
		UnitID:  req.UnitID,
		UserID:  req.UserID,
		Role:    role,
		StartAt: req.StartAt,
		EndAt:   req.EndAt,
	}
	if err = tenant.Validate(); err != nil {
		return entity.Tenant{}, err
	}

	err = u.db.Transaction(func(tx *gorm.DB) error {
		// Store tenant information
		if result := u.db.Create(&tenant); result.Error != nil {
			if mysqlErr, ok := result.Error.(*mysql.MySQLError); ok {
				if mysqlErr.Number == utility.MysqlErrorConflictNumber {
					return model.NewExpectedError("another account with the same GovID / (Tower, Floor, TenantNo) already exists", "TENANT_CONFLICT", http.StatusConflict, strconv.FormatUint(req.UnitID, 10))
				}
			}
			return model.NewUnknownError(strconv.FormatUint(req.UnitID, 10), result.Error)
		}

		// Delete all uploaded documents on ownership changes
		if role == entity.TenantRoleOwner {
			unit.AJB = []byte{}
			unit.Akte = []byte{}
			if result := u.db.Save(&unit); result.Error != nil {
				return model.NewUnknownError(strconv.FormatUint(req.UnitID, 10), result.Error)
			}
		}

		return nil
	})
	if err != nil {
		return entity.Tenant{}, err
	}

	return tenant, nil
}

func (u *tenantCreationUsecase) validateCreateTenantRequest(_ context.Context, req request.CreateTenantRequest, role entity.TenantRole) (entity.Unit, error) {
	var user entity.User
	if err := u.db.Where("id = ?", req.UserID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Unit{}, model.NewExpectedError("target user not found", "TENANT_INVALID", http.StatusBadRequest, strconv.FormatUint(req.UserID, 10))
		}
		return entity.Unit{}, model.NewUnknownError(strconv.FormatUint(req.UserID, 10), err)
	}

	var unit entity.Unit
	if err := u.db.Where("id = ?", req.UnitID).First(&unit).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Unit{}, model.NewExpectedError("target unit not found", "TENANT_INVALID", http.StatusBadRequest, strconv.FormatUint(req.UnitID, 10))
		}
		return entity.Unit{}, model.NewUnknownError(strconv.FormatUint(req.UserID, 10), err)
	}

	var tenant entity.Tenant
	now := time.Now()

	// If the unit still has an active owner, raise error
	if role == entity.TenantRoleOwner {
		if err := u.db.Where("unit_id = ? AND role = ? AND (end_at IS NULL OR end_at > ?)", req.UnitID, entity.TenantRoleOwner, now).First(&tenant).Error; err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return entity.Unit{}, model.NewUnknownError(strconv.FormatUint(req.UserID, 10), err)
			}
		} else {
			return entity.Unit{}, model.NewExpectedError("target unit already has an active owner", "TENANT_CONFLICT", http.StatusConflict, strconv.FormatUint(req.UnitID, 10))
		}
	}

	// If the unit and user already has an active relationship, raise error
	if err := u.db.Where("unit_id = ? AND user_id = ? AND (end_at IS NULL OR end_at > ?)", req.UnitID, req.UserID, now).First(&tenant).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Unit{}, model.NewUnknownError(strconv.FormatUint(req.UserID, 10), err)
		}
	} else {
		return entity.Unit{}, model.NewExpectedError("unit and user already has an relationship", "TENANT_CONFLICT", http.StatusConflict, strconv.FormatUint(req.UnitID, 10))
	}

	return unit, nil
}
