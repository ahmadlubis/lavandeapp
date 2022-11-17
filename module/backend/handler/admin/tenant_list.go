package admin

import (
	"encoding/json"
	"github.com/ahmadlubis/lavandeapp/module/backend/handler"
	"github.com/ahmadlubis/lavandeapp/module/backend/model"
	"github.com/ahmadlubis/lavandeapp/module/backend/model/request"
	"github.com/ahmadlubis/lavandeapp/module/backend/usecase"
	"net/http"
	"strconv"
	"strings"
)

type tenantListHandler struct {
	usecase usecase.TenantListUsecase
}

func NewTenantListHandler(usecase usecase.TenantListUsecase) handler.Handler {
	return &tenantListHandler{usecase: usecase}
}

func (h *tenantListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) error {
	query := r.URL.Query()

	var unitID, userID, limit, offset uint64
	var activeOnly bool
	var err error
	if query.Get("unit_id") != "" {
		if unitID, err = strconv.ParseUint(query.Get("unit_id"), 10, 64); err != nil {
			return model.NewExpectedError("unit_id must be a number", "TENANT_INVALID", http.StatusBadRequest, "")
		}
	}
	if query.Get("user_id") != "" {
		if userID, err = strconv.ParseUint(query.Get("user_id"), 10, 64); err != nil {
			return model.NewExpectedError("user_id must be a number", "TENANT_INVALID", http.StatusBadRequest, "")
		}
	}
	if strings.EqualFold(query.Get("active_only"), "true") {
		activeOnly = true
	}
	if limit, err = strconv.ParseUint(query.Get("limit"), 10, 64); err != nil {
		return model.NewExpectedError("limit must be a number", "TENANT_INVALID", http.StatusBadRequest, "")
	}
	if offset, err = strconv.ParseUint(query.Get("offset"), 10, 64); err != nil {
		return model.NewExpectedError("offset must be a number", "TENANT_INVALID", http.StatusBadRequest, "")
	}

	req := request.ListTenantRequest{
		UnitID:     unitID,
		UserID:     userID,
		ActiveOnly: activeOnly,
		Limit:      limit,
		Offset:     offset,
	}
	resp, err := h.usecase.List(r.Context(), req)
	if err != nil {
		return err
	}

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		return err
	}

	return nil
}
