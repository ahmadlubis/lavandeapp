package admin

import (
	"encoding/json"
	"github.com/ahmadlubis/lavandeapp/module/backend/handler"
	"github.com/ahmadlubis/lavandeapp/module/backend/model"
	"github.com/ahmadlubis/lavandeapp/module/backend/model/request"
	"github.com/ahmadlubis/lavandeapp/module/backend/usecase"
	"net/http"
	"strconv"
)

type userListHandler struct {
	usecase usecase.UserListUsecase
}

func NewUserListHandler(usecase usecase.UserListUsecase) handler.Handler {
	return &userListHandler{usecase: usecase}
}

func (h *userListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) error {
	query := r.URL.Query()

	limit, err := strconv.ParseUint(query.Get("limit"), 10, 64)
	if err != nil {
		return model.NewExpectedError("limit must be a number", "USER_INVALID", http.StatusBadRequest, "")
	}
	offset, err := strconv.ParseUint(query.Get("offset"), 10, 64)
	if err != nil {
		return model.NewExpectedError("offset must be a number", "USER_INVALID", http.StatusBadRequest, "")
	}

	req := request.ListUserRequest{
		Name:     query.Get("name"),
		NIK:      query.Get("nik"),
		Email:    query.Get("email"),
		PhoneNo:  query.Get("phone_no"),
		Status:   query.Get("status"),
		Religion: query.Get("religion"),
		Limit:    limit,
		Offset:   offset,
	}
	resp, err := h.usecase.List(r.Context(), req)
	if err != nil {
		return err
	}

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		return err
	}

	return nil
}
