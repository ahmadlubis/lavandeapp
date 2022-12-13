package admin

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/ahmadlubis/lavandeapp/module/backend/handler"
	"github.com/ahmadlubis/lavandeapp/module/backend/model"
	"github.com/ahmadlubis/lavandeapp/module/backend/model/request"
	"github.com/ahmadlubis/lavandeapp/module/backend/usecase"
)

type userResetPasswordHandler struct {
	usecase usecase.UserUpdateUsecase
}

func NewUserResetPasswordHandler(usecase usecase.UserUpdateUsecase) handler.Handler {
	return &userResetPasswordHandler{usecase: usecase}
}

func (h *userResetPasswordHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) error {
	reqBody, _ := io.ReadAll(r.Body)

	var req request.UpdateUserRequest
	err := json.Unmarshal(reqBody, &req)
	if err != nil {
		return model.NewExpectedError("bad request format", "UNIT_INVALID", http.StatusBadRequest, "")
	}

	unit, err := h.usecase.ResetPassword(r.Context(), req)
	if err != nil {
		return err
	}

	err = json.NewEncoder(w).Encode(unit)
	if err != nil {
		return err
	}

	return nil
}
