package handler

import (
	"encoding/json"
	"github.com/ahmadlubis/lavandeapp/module/backend/entity"
	"github.com/ahmadlubis/lavandeapp/module/backend/model"
	"net/http"
)

type userSelfInfoHandler struct{}

func NewUserSelfInfoHandler() Handler {
	return &userSelfInfoHandler{}
}

func (h *userSelfInfoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) error {
	var user, ok = r.Context().Value(RequestSubjectContextKey).(entity.User)
	if !ok {
		return model.InvalidTokenError
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		return err
	}

	return nil
}
