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

	if err = u.validateCreateTenantRequest(ctx, req, role); err != nil {
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

	if result := u.db.Create(&tenant); result.Error != nil {
		if mysqlErr, ok := result.Error.(*mysql.MySQLError); ok {
			if mysqlErr.Number == utility.MysqlErrorConflictNumber {
				return entity.Tenant{}, model.NewExpectedError("another account with the same GovID / (Tower, Floor, TenantNo) already exists", "TENANT_CONFLICT", http.StatusConflict, strconv.FormatUint(req.UnitID, 10))
			}
		}
		return entity.Tenant{}, model.NewUnknownError(strconv.FormatUint(req.UnitID, 10), result.Error)
	}

	return tenant, nil
}

func (u *tenantCreationUsecase) validateCreateTenantRequest(_ context.Context, req request.CreateTenantRequest, role entity.TenantRole) error {
	var user entity.User
	if err := u.db.Where("id = ?", req.UserID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.NewExpectedError("target user not found", "TENANT_INVALID", http.StatusBadRequest, strconv.FormatUint(req.UserID, 10))
		}
		return model.NewUnknownError(strconv.FormatUint(req.UserID, 10), err)
	}

	var unit entity.Unit
	if err := u.db.Where("id = ?", req.UnitID).First(&unit).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.NewExpectedError("target unit not found", "TENANT_INVALID", http.StatusBadRequest, strconv.FormatUint(req.UnitID, 10))
		}
		return model.NewUnknownError(strconv.FormatUint(req.UserID, 10), err)
	}

	return nil
}
