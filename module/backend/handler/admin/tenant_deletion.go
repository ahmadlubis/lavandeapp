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

type tenantDeletionHandler struct {
	usecase usecase.TenantDeletionUsecase
}

func NewTenantDeletionHandler(usecase usecase.TenantDeletionUsecase) handler.Handler {
	return &tenantDeletionHandler{usecase: usecase}
}

func (h *tenantDeletionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) error {
	reqBody, _ := io.ReadAll(r.Body)

	var req request.DeleteTenantRequest
	err := json.Unmarshal(reqBody, &req)
	if err != nil {
		return model.NewExpectedError("bad request format", "TENANT_INVALID", http.StatusBadRequest, "")
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
