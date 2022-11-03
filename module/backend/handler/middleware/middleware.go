package middleware

import (
	"github.com/ahmadlubis/lavandeapp/module/backend/handler"
	"net/http"
)

func WithDefaultMiddlewares(h handler.Handler) http.Handler {
	return handler.NewErrorFormatter(WithDefaultHeaders(WithLogging(h)))
}
