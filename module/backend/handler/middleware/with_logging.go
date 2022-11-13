package middleware

import (
	"github.com/ahmadlubis/lavandeapp/module/backend/handler"
	"github.com/ahmadlubis/lavandeapp/module/backend/model"
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
	if err != nil {
		if restError, ok := err.(*model.RestError); ok && restError.OriginalError != nil {
			log.Printf("%s %s %v; ERROR: %s", r.Method, r.URL.Path, time.Since(start), restError.OriginalError.Error())
		} else {
			log.Printf("%s %s %v; ERROR: %s", r.Method, r.URL.Path, time.Since(start), err.Error())
		}
	} else {
		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
	}

	return err
}

func WithLogging(handlerToWrap handler.Handler) handler.Handler {
	return &loggingMiddleware{handlerToWrap}
}
