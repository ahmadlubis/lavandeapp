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

type unitCreationHandler struct {
	usecase usecase.UnitCreationUsecase
}

func NewUnitCreationHandler(usecase usecase.UnitCreationUsecase) handler.Handler {
	return &unitCreationHandler{usecase: usecase}
}

func (h *unitCreationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) error {
	reqBody, _ := io.ReadAll(r.Body)

	var req request.CreateUnitRequest
	err := json.Unmarshal(reqBody, &req)
	if err != nil {
		return model.NewExpectedError("bad request format", "UNIT_INVALID", http.StatusBadRequest, "")
	}

	unit, err := h.usecase.Create(r.Context(), req)
	if err != nil {
		return err
	}

	err = json.NewEncoder(w).Encode(unit)
	if err != nil {
		return err
	}

	return nil
}
