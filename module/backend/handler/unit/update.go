package unit

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

type unitUpdateHandler struct {
	validator usecase.UnitOwnerVerificationUsecase
	usecase   usecase.UnitUpdateUsecase
}

func NewUnitUpdateHandler(validator usecase.UnitOwnerVerificationUsecase, usecase usecase.UnitUpdateUsecase) handler.Handler {
	return &unitUpdateHandler{
		validator: validator,
		usecase:   usecase,
	}
}

func (h *unitUpdateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) error {
	temp := r.Context().Value(handler.RequestSubjectContextKey)
	var user, ok = temp.(entity.User)
	if !ok {
		return model.InvalidTokenError
	}

	reqBody, _ := io.ReadAll(r.Body)
	var req request.UpdateUnitRequest
	err := json.Unmarshal(reqBody, &req)
	if err != nil {
		return model.NewExpectedError("bad request format", "UNIT_INVALID", http.StatusBadRequest, "")
	}

	// Verify that requester is the owner of target Unit
	err = h.validator.VerifyOwner(r.Context(), req.ID, user.ID)
	if err != nil {
		return err
	}

	unit, err := h.usecase.Update(r.Context(), req)
	if err != nil {
		return err
	}

	err = json.NewEncoder(w).Encode(unit)
	if err != nil {
		return err
	}

	return nil
}
