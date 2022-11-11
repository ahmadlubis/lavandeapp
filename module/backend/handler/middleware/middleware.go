package middleware

import (
	"github.com/ahmadlubis/lavandeapp/module/backend/handler"
	"github.com/ahmadlubis/lavandeapp/module/backend/usecase"
	"net/http"
)

func WithDefaultAdminMiddlewares(usecase usecase.UserTokenVerificationUsecase, h handler.Handler) http.Handler {
	return handler.NewErrorFormatter(WithDefaultHeaders(WithLogging(WithLoggedInAdmin(h, usecase))))
}

func WithDefaultMiddlewares(usecase usecase.UserTokenVerificationUsecase, h handler.Handler) http.Handler {
	return handler.NewErrorFormatter(WithDefaultHeaders(WithLogging(WithLoggedInUser(h, usecase))))
}

func WithDefaultNoAuthMiddlewares(h handler.Handler) http.Handler {
	return handler.NewErrorFormatter(WithDefaultHeaders(WithLogging(h)))
}
