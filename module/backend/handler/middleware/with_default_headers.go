package middleware

import (
	"github.com/ahmadlubis/lavandeapp/module/backend/handler"
	"net/http"
)

type defaultHeadersMiddleware struct {
	handler handler.Handler
}

func (l *defaultHeadersMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) error {
	w.Header().Add("Content-Type", "application/json")
	err := l.handler.ServeHTTP(w, r)

	return err
}

func WithDefaultHeaders(handlerToWrap handler.Handler) handler.Handler {
	return &defaultHeadersMiddleware{handlerToWrap}
}
