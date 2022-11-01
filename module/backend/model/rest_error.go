package model

import "net/http"

type RestError struct {
	Message       string
	Code          string
	Status        int
	TrackId       string
	OriginalError error
}

func (r *RestError) Error() string {
	return r.Message
}

func NewExpectedError(message, code string, status int, trackId string) *RestError {
	return &RestError{
		Message: message,
		Code:    code,
		Status:  status,
		TrackId: trackId,
	}
}

func NewUnknownError(trackId string, err error) *RestError {
	return &RestError{
		Message:       "Internal Server Error",
		Code:          "UNKNOWN",
		Status:        http.StatusInternalServerError,
		TrackId:       trackId,
		OriginalError: err,
	}
}
