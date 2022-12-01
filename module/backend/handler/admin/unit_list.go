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

type unitListHandler struct {
	usecase usecase.UnitListUsecase
}

func NewUnitListHandler(usecase usecase.UnitListUsecase) handler.Handler {
	return &unitListHandler{usecase: usecase}
}

func (h *unitListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) error {
	query := r.URL.Query()

	var ID, ownerID, limit, offset uint64
	var err error
	if query.Get("id") != "" {
		if ID, err = strconv.ParseUint(query.Get("id"), 10, 64); err != nil {
			return model.NewExpectedError("id must be a number", "UNIT_INVALID", http.StatusBadRequest, "")
		}
	}
	if query.Get("owner_id") != "" {
		if ownerID, err = strconv.ParseUint(query.Get("owner_id"), 10, 64); err != nil {
			return model.NewExpectedError("owner_id must be a number", "UNIT_INVALID", http.StatusBadRequest, "")
		}
	}
	if limit, err = strconv.ParseUint(query.Get("limit"), 10, 64); err != nil {
		return model.NewExpectedError("limit must be a number", "UNIT_INVALID", http.StatusBadRequest, "")
	}
	if offset, err = strconv.ParseUint(query.Get("offset"), 10, 64); err != nil {
		return model.NewExpectedError("offset must be a number", "UNIT_INVALID", http.StatusBadRequest, "")
	}

	req := request.ListUnitRequest{
		ID:		 ID,
		OwnerID: ownerID,
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
