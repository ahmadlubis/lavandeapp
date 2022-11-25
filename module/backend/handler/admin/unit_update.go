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

type unitUpdateHandler struct {
	usecase usecase.UnitUpdateUsecase
}

func NewUnitUpdateHandler(usecase usecase.UnitUpdateUsecase) handler.Handler {
	return &unitUpdateHandler{usecase: usecase}
}

func (h *unitUpdateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) error {
	reqBody, _ := io.ReadAll(r.Body)

	var req request.AdminUpdateUnitRequest
	err := json.Unmarshal(reqBody, &req)
	if err != nil {
		return model.NewExpectedError("bad request format", "UNIT_INVALID", http.StatusBadRequest, "")
	}

	unit, err := h.usecase.Update(r.Context(), request.UpdateUnitRequest{
		ID:     req.ID,
		GovID:  req.GovID,
		Tower:  req.Tower,
		Floor:  req.Floor,
		UnitNo: req.UnitNo,
	})
	if err != nil {
		return err
	}

	err = json.NewEncoder(w).Encode(unit)
	if err != nil {
		return err
	}

	return nil
}
