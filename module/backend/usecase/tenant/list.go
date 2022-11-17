package tenant

import (
	"context"
	"github.com/ahmadlubis/lavandeapp/module/backend/entity"
	"github.com/ahmadlubis/lavandeapp/module/backend/model"
	"github.com/ahmadlubis/lavandeapp/module/backend/model/request"
	"github.com/ahmadlubis/lavandeapp/module/backend/model/response"
	"github.com/ahmadlubis/lavandeapp/module/backend/usecase"
	"gorm.io/gorm"
	"strings"
	"time"
)

const (
	minLimit = 10
	maxLimit = 50
)

type tenantListUsecase struct {
	db *gorm.DB
}

func NewTenantListUsecase(db *gorm.DB) usecase.TenantListUsecase {
	return &tenantListUsecase{db: db}
}

func (u *tenantListUsecase) List(ctx context.Context, request request.ListTenantRequest) (response.ListTenantResponse, error) {
	var tenants []entity.Tenant
	var total int64

	req, err := u.normalizeListRequest(request)
	if err != nil {
		return response.ListTenantResponse{}, err
	}

	// Build WHERE query and its parameters
	var conditions []string
	var params []interface{}
	if req.UnitID != 0 {
		conditions = append(conditions, "unit_id = ?")
		params = append(params, req.UnitID)
	}
	if req.UserID != 0 {
		conditions = append(conditions, "user_id = ?")
		params = append(params, req.UserID)
	}
	if req.ActiveOnly {
		conditions = append(conditions, "(end_at IS NULL OR end_at > ?)")
		params = append(params, time.Now())
	}
	query := strings.Join(conditions, " AND ")

	// Fetch the list of tenants considering pagination
	if result := u.db.Where(query, params...).Offset(int(req.Offset)).Limit(int(req.Limit)).Order("id").Find(&tenants); result.Error != nil {
		return response.ListTenantResponse{}, model.NewUnknownError("", result.Error)
	}

	// Fetch the total of all tenants with NO regards to the pagination
	if result := u.db.Model(&tenants).Where(query, params...).Count(&total); result.Error != nil {
		return response.ListTenantResponse{}, model.NewUnknownError("", result.Error)
	}

	return response.ListTenantResponse{
		Data: tenants,
		Meta: response.PaginationMeta{
			Limit:  req.Limit,
			Offset: req.Offset,
			Count:  uint64(len(tenants)),
			Total:  uint64(total),
		},
	}, nil
}

func (u *tenantListUsecase) normalizeListRequest(req request.ListTenantRequest) (request.ListTenantRequest, error) {
	if req.Limit < minLimit {
		req.Limit = minLimit
	}
	if req.Limit > maxLimit {
		req.Limit = maxLimit
	}
	if req.Offset < 0 {
		req.Offset = 0
	}

	return req, nil
}
