package user

import (
	"encoding/json"
	"github.com/ahmadlubis/lavandeapp/module/backend/entity"
	"github.com/ahmadlubis/lavandeapp/module/backend/handler"
	"github.com/ahmadlubis/lavandeapp/module/backend/model"
	"github.com/ahmadlubis/lavandeapp/module/backend/model/request"
	"github.com/ahmadlubis/lavandeapp/module/backend/model/response"
	"github.com/ahmadlubis/lavandeapp/module/backend/usecase"
	"net/http"
	"strconv"
)

type userListHandler struct {
	validator usecase.UnitOwnerVerificationUsecase
	usecase   usecase.UserListUsecase
}

func NewUserListHandler(validator usecase.UnitOwnerVerificationUsecase, usecase usecase.UserListUsecase) handler.Handler {
	return &userListHandler{validator: validator, usecase: usecase}
}

func (h *userListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) error {
	var user, ok = r.Context().Value(handler.RequestSubjectContextKey).(entity.User)
	if !ok {
		return model.InvalidTokenError
	}

	// Verify that requester is the owner of any Unit
	var err = h.validator.VerifyAnyOwner(r.Context(), user.ID)
	if err != nil {
		return err
	}

	query := r.URL.Query()

	var id, limit, offset uint64

	if id, err = strconv.ParseUint(query.Get("id"), 10, 64); err != nil {
		id = 0
	}
	limit, err = strconv.ParseUint(query.Get("limit"), 10, 64)
	if err != nil {
		return model.NewExpectedError("limit must be a number", "USER_INVALID", http.StatusBadRequest, "")
	}
	offset, err = strconv.ParseUint(query.Get("offset"), 10, 64)
	if err != nil {
		return model.NewExpectedError("offset must be a number", "USER_INVALID", http.StatusBadRequest, "")
	}

	req := request.ListUserRequest{
		ID:     id,
		Name:   query.Get("name"),
		Email:  query.Get("email"),
		Limit:  limit,
		Offset: offset,
	}
	result, err := h.usecase.List(r.Context(), req)
	if err != nil {
		return err
	}

	var resp response.GenericListResponse
	for _, v := range result.Data {
		resp.Data = append(resp.Data, map[string]string{
			"id":    strconv.FormatUint(v.ID, 10),
			"email": v.Email,
			"name":  v.Name,
		})
	}
	resp.Meta = result.Meta
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		return err
	}

	return nil
}
