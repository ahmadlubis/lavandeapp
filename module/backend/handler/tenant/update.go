package tenant

import (
	"encoding/json"
	"github.com/ahmadlubis/lavandeapp/module/backend/entity"
	"github.com/ahmadlubis/lavandeapp/module/backend/handler"
	"github.com/ahmadlubis/lavandeapp/module/backend/model"
	"github.com/ahmadlubis/lavandeapp/module/backend/model/request"
	"github.com/ahmadlubis/lavandeapp/module/backend/usecase"
	"io"
	"net/http"
)

type tenantCreationHandler struct {
	validator usecase.UnitOwnerVerificationUsecase
	usecase   usecase.TenantCreationUsecase
}

func NewTenantCreationHandler(validator usecase.UnitOwnerVerificationUsecase, usecase usecase.TenantCreationUsecase) handler.Handler {
	return &tenantCreationHandler{
		validator: validator,
		usecase:   usecase,
	}
}

func (h *tenantCreationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) error {
	var user, ok = r.Context().Value(handler.RequestSubjectContextKey).(entity.User)
	if !ok {
		return model.InvalidTokenError
	}

	reqBody, _ := io.ReadAll(r.Body)
	var req request.CreateTenantRequest
	err := json.Unmarshal(reqBody, &req)
	if err != nil {
		return model.NewExpectedError("bad request format", "TENANT_INVALID", http.StatusBadRequest, "")
	}

	// Verify that requester is the owner of target Tenant
	err = h.validator.VerifyOwner(r.Context(), req.UnitID, user.ID)
	if err != nil {
		return err
	}

	tenant, err := h.usecase.Create(r.Context(), req)
	if err != nil {
		return err
	}

	err = json.NewEncoder(w).Encode(tenant)
	if err != nil {
		return err
	}

	return nil
}
