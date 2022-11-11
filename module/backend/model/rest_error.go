package model

import "net/http"

var (
	InvalidTokenError = NewExpectedError("session expired, please login again", "USER_UNAUTHORIZED", http.StatusUnauthorized, "")
	UserNotFoundError = NewExpectedError("user not found", "USER_NOT_FOUND", http.StatusNotFound, "")
	NonAdminError     = NewExpectedError("forbidden", "USER_FORBIDDEN", http.StatusForbidden, "")
)

type RestError struct {
	Message       string `json:"error_message"`
	Code          string `json:"error_code"`
	Status        int    `json:"-"`
	TrackId       string `json:"-"`
	OriginalError error  `json:"-"`
}

func (r *RestError) Error() string {
	return r.Message
}

func (r *RestError) WithTrackId(trackId string) error {
	r.TrackId = trackId
	return r
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
