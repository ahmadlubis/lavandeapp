package middleware

import (
	"github.com/ahmadlubis/lavandeapp/module/backend/handler"
	"github.com/ahmadlubis/lavandeapp/module/backend/usecase"
	"net/http"
)

func WithDefaultMiddlewares(h handler.Handler, usecase usecase.UserTokenVerificationUsecase) http.Handler {
	return handler.NewErrorFormatter(WithDefaultHeaders(WithLogging(WithLoggedInUser(h, usecase))))
}

func WithDefaultNoAuthMiddlewares(h handler.Handler) http.Handler {
	return handler.NewErrorFormatter(WithDefaultHeaders(WithLogging(h)))
}
