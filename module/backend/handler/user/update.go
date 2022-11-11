package user

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

type userUpdateHandler struct {
	usecase usecase.UserUpdateUsecase
}

func NewUserUpdateHandler(usecase usecase.UserUpdateUsecase) handler.Handler {
	return &userUpdateHandler{
		usecase: usecase,
	}
}

func (h *userUpdateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) error {
	var user, ok = r.Context().Value(handler.RequestSubjectContextKey).(entity.User)
	if !ok {
		return model.InvalidTokenError
	}

	reqBody, _ := io.ReadAll(r.Body)

	var req request.UpdateUserRequest
	err := json.Unmarshal(reqBody, &req)
	if err != nil {
		return model.NewExpectedError("bad request format", "USER_INVALID", http.StatusBadRequest, "")
	}

	req.TargetEmail = user.Email // Force currently logged-in users as target
	user, err = h.usecase.Update(r.Context(), req)
	if err != nil {
		return err
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		return err
	}

	return nil
}
