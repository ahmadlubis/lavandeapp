package unit

import (
	"encoding/json"
	"github.com/ahmadlubis/lavandeapp/module/backend/entity"
	"github.com/ahmadlubis/lavandeapp/module/backend/handler"
	"github.com/ahmadlubis/lavandeapp/module/backend/model"
	"github.com/ahmadlubis/lavandeapp/module/backend/model/request"
	"github.com/ahmadlubis/lavandeapp/module/backend/usecase"
	"net/http"
	"strconv"
)

type unitListHandler struct {
	usecase usecase.UnitListUsecase
}

func NewUnitListHandler(usecase usecase.UnitListUsecase) handler.Handler {
	return &unitListHandler{usecase: usecase}
}

func (h *unitListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) error {
	temp := r.Context().Value(handler.RequestSubjectContextKey)
	var user, ok = temp.(entity.User)
	if !ok {
		return model.InvalidTokenError
	}

	query := r.URL.Query()

	var id, limit, offset uint64
	var err error
	if limit, err = strconv.ParseUint(query.Get("limit"), 10, 64); err != nil {
		return model.NewExpectedError("limit must be a number", "UNIT_INVALID", http.StatusBadRequest, "")
	}
	if offset, err = strconv.ParseUint(query.Get("offset"), 10, 64); err != nil {
		return model.NewExpectedError("offset must be a number", "UNIT_INVALID", http.StatusBadRequest, "")
	}
	if id, err = strconv.ParseUint(query.Get("id"), 10, 64); err != nil {
		id = 0
	}

	// FORCE filter by current_user owned units
	req := request.ListUnitRequest{
		ID:      id,
		OwnerID: user.ID,
		GovID:   query.Get("gov_id"),
		Tower:   query.Get("tower"),
		Floor:   query.Get("floor"),
		UnitNo:  query.Get("unit_no"),
		Limit:   limit,
		Offset:  offset,
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
