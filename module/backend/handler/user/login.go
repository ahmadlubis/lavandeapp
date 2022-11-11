package user

import (
	"encoding/json"
	"github.com/ahmadlubis/lavandeapp/module/backend/handler"
	"github.com/ahmadlubis/lavandeapp/module/backend/model"
	"github.com/ahmadlubis/lavandeapp/module/backend/model/request"
	"github.com/ahmadlubis/lavandeapp/module/backend/usecase"
	"io"
	"net/http"
)

type userLoginHandler struct {
	usecase usecase.UserLoginUsecase
}

func NewUserLoginHandler(usecase usecase.UserLoginUsecase) handler.Handler {
	return &userLoginHandler{usecase: usecase}
}

func (h *userLoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) error {
	reqBody, _ := io.ReadAll(r.Body)

	var req request.LoginUserRequest
	err := json.Unmarshal(reqBody, &req)
	if err != nil {
		return model.NewExpectedError("bad request format", "USER_INVALID", http.StatusBadRequest, "")
	}

	user, err := h.usecase.Login(r.Context(), req)
	if err != nil {
		return err
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		return err
	}

	return nil
}
