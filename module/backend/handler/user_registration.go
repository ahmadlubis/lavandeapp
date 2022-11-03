package handler

import (
	"encoding/json"
	"github.com/ahmadlubis/lavandeapp/module/backend/model"
	"github.com/ahmadlubis/lavandeapp/module/backend/model/request"
	"github.com/ahmadlubis/lavandeapp/module/backend/usecase"
	"io"
	"net/http"
)

type userRegistrationHandler struct {
	usecase usecase.UserRegistrationUsecase
}

func NewUserRegistrationHandler(usecase usecase.UserRegistrationUsecase) Handler {
	return &userRegistrationHandler{usecase: usecase}
}

func (h *userRegistrationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) error {
	reqBody, _ := io.ReadAll(r.Body)

	var req request.RegisterUserRequest
	err := json.Unmarshal(reqBody, &req)
	if err != nil {
		return model.NewExpectedError("bad request format", "USER_INVALID", http.StatusBadRequest, "")
	}

	user, err := h.usecase.RegisterUser(r.Context(), req)
	if err != nil {
		return err
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		return err
	}

	return nil
}