package unit

import (
	"context"
	"github.com/ahmadlubis/lavandeapp/module/backend/entity"
	"github.com/ahmadlubis/lavandeapp/module/backend/model"
	"github.com/ahmadlubis/lavandeapp/module/backend/model/request"
	"github.com/ahmadlubis/lavandeapp/module/backend/model/response"
	"github.com/ahmadlubis/lavandeapp/module/backend/usecase"
	"gorm.io/gorm"
	"net/http"
	"strings"
	"time"
)

const (
	minLimit = 10
	maxLimit = 50
)

type unitListUsecase struct {
	db *gorm.DB
}

func NewUnitListUsecase(db *gorm.DB) usecase.UnitListUsecase {
	return &unitListUsecase{db: db}
}

func (u *unitListUsecase) List(ctx context.Context, request request.ListUnitRequest) (response.ListUnitResponse, error) {
	var units []entity.Unit
	var total int64

	req, err := u.normalizeListRequest(request)
	if err != nil {
		return response.ListUnitResponse{}, err
	}

	// Build WHERE query and its parameters
	var conditions []string
	var params []interface{}
	if req.ID != 0 {
		conditions = append(conditions, "id = ?")
		params = append(params, req.ID)
	}
	if req.OwnerID != 0 {
		unitIDs, err := u.getOwnedUnitIDs(ctx, req.OwnerID)
		if err != nil {
			return response.ListUnitResponse{}, err
		}
		conditions = append(conditions, "id IN ?")
		params = append(params, unitIDs)
	}
	if req.GovID != "" {
		conditions = append(conditions, "gov_id = ?")
		params = append(params, req.GovID)
	}
	if req.Tower != "" {
		conditions = append(conditions, "tower = ?")
		params = append(params, req.Tower)
	}
	if req.Floor != "" {
		conditions = append(conditions, "floor = ?")
		params = append(params, req.Floor)
	}
	if req.UnitNo != "" {
		conditions = append(conditions, "unit_no = ?")
		params = append(params, req.UnitNo)
	}

	query := strings.Join(conditions, " AND ")

	// Fetch the list of units considering pagination
	if result := u.db.Where(query, params...).Offset(int(req.Offset)).Limit(int(req.Limit)).Order("id").Find(&units); result.Error != nil {
		return response.ListUnitResponse{}, model.NewUnknownError("", result.Error)
	}

	// Fetch the total of all units with NO regards to the pagination
	if result := u.db.Model(&units).Where(query, params...).Count(&total); result.Error != nil {
		return response.ListUnitResponse{}, model.NewUnknownError("", result.Error)
	}

	return response.ListUnitResponse{
		Data: units,
		Meta: response.PaginationMeta{
			Limit:  req.Limit,
			Offset: req.Offset,
			Count:  uint64(len(units)),
			Total:  uint64(total),
		},
	}, nil
}

func (u *unitListUsecase) getOwnedUnitIDs(_ context.Context, ownerID uint64) ([]uint64, error) {
	var ids []uint64
	if result := u.db.Model(&entity.Tenant{}).Where("user_id = ? AND role = ? AND (end_at IS NULL OR end_at > ?)", ownerID, entity.TenantRoleOwner, time.Now()).Pluck("unit_id", &ids); result.Error != nil {
		return ids, model.NewUnknownError("", result.Error)
	}
	return ids, nil
}

func (u *unitListUsecase) normalizeListRequest(req request.ListUnitRequest) (request.ListUnitRequest, error) {
	if req.Tower == "" && (req.Floor != "" || req.UnitNo != "") {
		return req, model.NewExpectedError("tower filter needs to be present to search by floor / unit_no", "UNIT_INVALID", http.StatusBadRequest, "")
	}
	if req.Floor == "" && req.UnitNo != "" {
		return req, model.NewExpectedError("tower and floor filter need to be present to search by unit_no", "UNIT_INVALID", http.StatusBadRequest, "")
	}

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
