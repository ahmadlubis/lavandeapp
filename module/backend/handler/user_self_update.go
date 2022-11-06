package handler

import (
	"encoding/json"
	"github.com/ahmadlubis/lavandeapp/module/backend/entity"
	"github.com/ahmadlubis/lavandeapp/module/backend/model"
	"github.com/ahmadlubis/lavandeapp/module/backend/model/request"
	"github.com/ahmadlubis/lavandeapp/module/backend/usecase"
	"io"
	"net/http"
)

type userSelfUpdateHandler struct {
	usecase usecase.UserSelfUpdateUsecase
}

func NewUserSelfUpdateHandler(usecase usecase.UserSelfUpdateUsecase) Handler {
	return &userSelfUpdateHandler{
		usecase: usecase,
	}
}

func (h *userSelfUpdateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) error {
	var user, ok = r.Context().Value(RequestSubjectContextKey).(entity.User)
	if !ok {
		return model.InvalidTokenError
	}

	reqBody, _ := io.ReadAll(r.Body)

	var req request.SelfUpdateUserRequest
	err := json.Unmarshal(reqBody, &req)
	if err != nil {
		return model.NewExpectedError("bad request format", "USER_INVALID", http.StatusBadRequest, "")
	}
	req.TargetEmail = user.Email

	user, err = h.usecase.SelfUpdate(r.Context(), req)
	if err != nil {
		return err
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		return err
	}

	return nil
}
