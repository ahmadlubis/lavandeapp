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

type tenantDeletionHandler struct {
	validator usecase.UnitOwnerVerificationUsecase
	usecase   usecase.TenantDeletionUsecase
}

func NewTenantDeletionHandler(validator usecase.UnitOwnerVerificationUsecase, usecase usecase.TenantDeletionUsecase) handler.Handler {
	return &tenantDeletionHandler{
		validator: validator,
		usecase:   usecase,
	}
}

func (h *tenantDeletionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) error {
	var user, ok = r.Context().Value(handler.RequestSubjectContextKey).(entity.User)
	if !ok {
		return model.InvalidTokenError
	}

	reqBody, _ := io.ReadAll(r.Body)
	var req request.DeleteTenantRequest
	err := json.Unmarshal(reqBody, &req)
	if err != nil {
		return model.NewExpectedError("bad request format", "TENANT_INVALID", http.StatusBadRequest, "")
	}

	// Verify that requester is the owner of target Tenant
	err = h.validator.VerifyOwner(r.Context(), req.UnitID, user.ID)
	if err != nil {
		return err
	}

	// Verify that the terget userID is NOT the current_user
	if user.ID == req.UserID {
		return model.NewExpectedError("can't delete your own role", "TENANT_INVALID", http.StatusBadRequest, "")
	}

	err = h.usecase.Delete(r.Context(), req)
	if err != nil {
		return err
	}

	err = json.NewEncoder(w).Encode(map[string]string{"message": "tenant deleted"})
	if err != nil {
		return err
	}

	return nil
}
