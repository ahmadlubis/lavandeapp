package admin

import (
	"encoding/json"
	"github.com/ahmadlubis/lavandeapp/module/backend/handler"
	"github.com/ahmadlubis/lavandeapp/module/backend/model"
	"github.com/ahmadlubis/lavandeapp/module/backend/model/request"
	"github.com/ahmadlubis/lavandeapp/module/backend/usecase"
	"io"
	"net/http"
)

type tenantCreationHandler struct {
	usecase usecase.TenantCreationUsecase
}

func NewTenantCreationHandler(usecase usecase.TenantCreationUsecase) handler.Handler {
	return &tenantCreationHandler{usecase: usecase}
}

func (h *tenantCreationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) error {
	reqBody, _ := io.ReadAll(r.Body)

	var req request.CreateTenantRequest
	err := json.Unmarshal(reqBody, &req)
	if err != nil {
		return model.NewExpectedError("bad request format", "TENANT_INVALID", http.StatusBadRequest, "")
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
