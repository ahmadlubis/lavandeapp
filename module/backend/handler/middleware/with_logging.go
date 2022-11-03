package middleware

import (
	"github.com/ahmadlubis/lavandeapp/module/backend/handler"
	"log"
	"net/http"
	"time"
)

type loggingMiddleware struct {
	handler handler.Handler
}

func (l *loggingMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) error {
	start := time.Now()
	err := l.handler.ServeHTTP(w, r)
	log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))

	return err
}

func WithLogging(handlerToWrap handler.Handler) handler.Handler {
	return &loggingMiddleware{handlerToWrap}
}
