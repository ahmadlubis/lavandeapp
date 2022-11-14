package admin

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/ahmadlubis/lavandeapp/module/backend/handler"
	"github.com/ahmadlubis/lavandeapp/module/backend/model"
	"github.com/ahmadlubis/lavandeapp/module/backend/usecase"
)

type setSuperadminHandler struct {
	usecase usecase.SuperadminUsecase
}

func NewSuperAdminHandler(usecase usecase.SuperadminUsecase) handler.Handler {
	return &setSuperadminHandler{usecase: usecase}
}

func (h *setSuperadminHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) error {
	reqBody, _ := io.ReadAll(r.Body)

	var request = struct {
		UserID int `json:"user_id"`
	}{}

	err := json.Unmarshal(reqBody, &request)
	if err != nil {
		return model.NewExpectedError("bad request format", "TENANT_INVALID", http.StatusBadRequest, "")
	}

	return h.usecase.SetUserAsSuperadmin(r.Context(), request.UserID)
}
