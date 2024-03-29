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

type userUpdateHandler struct {
	usecase usecase.UserUpdateUsecase
}

func NewUserUpdateHandler(usecase usecase.UserUpdateUsecase) handler.Handler {
	return &userUpdateHandler{usecase: usecase}
}

func (h *userUpdateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) error {
	reqBody, _ := io.ReadAll(r.Body)

	var req request.AdminUpdateUserRequest
	err := json.Unmarshal(reqBody, &req)
	if err != nil {
		return model.NewExpectedError("bad request format", "UNIT_INVALID", http.StatusBadRequest, "")
	}

	unit, err := h.usecase.AdminUpdate(r.Context(), req)
	if err != nil {
		return err
	}

	err = json.NewEncoder(w).Encode(unit)
	if err != nil {
		return err
	}

	return nil
}
