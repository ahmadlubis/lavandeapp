package middleware

import (
	"context"
	"github.com/ahmadlubis/lavandeapp/module/backend/handler"
	"github.com/ahmadlubis/lavandeapp/module/backend/model"
	"github.com/ahmadlubis/lavandeapp/module/backend/usecase"
	"net/http"
	"strings"
)

type authMiddleware struct {
	handler handler.Handler
	usecase usecase.UserTokenVerificationUsecase
}

func (l *authMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) error {
	tokens := strings.Split(r.Header.Get("Authorization"), "Bearer ")
	if len(tokens) != 2 {
		return model.InvalidTokenError
	}

	user, err := l.usecase.VerifyToken(r.Context(), tokens[1])
	if err != nil {
		return err
	}

	newCtx := context.WithValue(r.Context(), handler.RequestSubjectContextKey, user)

	return l.handler.ServeHTTP(w, r.WithContext(newCtx))
}

func WithLoggedInUser(handlerToWrap handler.Handler, usecase usecase.UserTokenVerificationUsecase) handler.Handler {
	return &authMiddleware{
		handler: handlerToWrap,
		usecase: usecase,
	}
}
