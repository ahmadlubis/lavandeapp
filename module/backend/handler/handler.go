package handler

import (
	"encoding/json"
	"github.com/ahmadlubis/lavandeapp/module/backend/model"
	"net/http"
)

// HandleWithError and Handler will be used instead of http.Handler for all middleware and request handler
// This is done so that the error object is accessible on all Middleware, and only the last middleware need to format the error to the responseWriter
type HandleWithError func(w http.ResponseWriter, r *http.Request) error

type Handler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request) error
}

// ErrorFormatter will receive our custom Handler and output http.Handler,
// This middleware will be the last on the stack, and will format the error to the responseWriter if exists
type ErrorFormatter struct {
	handler Handler
}

func (l *ErrorFormatter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := l.handler.ServeHTTP(w, r)
	if err != nil {
		restErr, ok := err.(*model.RestError)
		if !ok {
			restErr = model.NewUnknownError("", err)
		}

		w.WriteHeader(restErr.Status)
		_ = json.NewEncoder(w).Encode(restErr)
	}
}

func NewErrorFormatter(handlerToWrap Handler) http.Handler {
	return &ErrorFormatter{handlerToWrap}
}
