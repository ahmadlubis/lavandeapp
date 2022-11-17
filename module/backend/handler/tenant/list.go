package tenant

import (
	"encoding/json"
	"github.com/ahmadlubis/lavandeapp/module/backend/entity"
	"github.com/ahmadlubis/lavandeapp/module/backend/handler"
	"github.com/ahmadlubis/lavandeapp/module/backend/model"
	"github.com/ahmadlubis/lavandeapp/module/backend/model/request"
	"github.com/ahmadlubis/lavandeapp/module/backend/usecase"
	"net/http"
	"strconv"
)

type tenantListHandler struct {
	validator usecase.UnitOwnerVerificationUsecase
	usecase   usecase.TenantListUsecase
}

func NewTenantListHandler(validator usecase.UnitOwnerVerificationUsecase, usecase usecase.TenantListUsecase) handler.Handler {
	return &tenantListHandler{validator: validator, usecase: usecase}
}

func (h *tenantListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) error {
	temp := r.Context().Value(handler.RequestSubjectContextKey)
	var user, ok = temp.(entity.User)
	if !ok {
		return model.InvalidTokenError
	}

	query := r.URL.Query()

	var unitID, limit, offset uint64
	var err error
	if unitID, err = strconv.ParseUint(query.Get("unit_id"), 10, 64); err != nil {
		return model.NewExpectedError("unit_id must be a number", "TENANT_INVALID", http.StatusBadRequest, "")
	}
	if limit, err = strconv.ParseUint(query.Get("limit"), 10, 64); err != nil {
		return model.NewExpectedError("limit must be a number", "TENANT_INVALID", http.StatusBadRequest, "")
	}
	if offset, err = strconv.ParseUint(query.Get("offset"), 10, 64); err != nil {
		return model.NewExpectedError("offset must be a number", "TENANT_INVALID", http.StatusBadRequest, "")
	}

	// Verify that requester is the owner of target Unit
	err = h.validator.VerifyOwner(r.Context(), unitID, user.ID)
	if err != nil {
		return err
	}

	// Force only search by unitID
	req := request.ListTenantRequest{
		UnitID:     unitID,
		ActiveOnly: true,
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
