package utility

import (
	"log"
	"net/http"
	"regexp"

	"github.com/ahmadlubis/lavandeapp/module/backend/model"
)

const asciiRegex = "^[\\x00-\\x7F]+$"

func ValidateUserPassword(passwd, trackId string) error {
	log.Println("New Pass -> ", passwd)
	if len(passwd) < 8 || len(passwd) > 32 {
		return model.NewExpectedError("password must be between 8 to 32 characters long", "USER_INVALID", http.StatusBadRequest, trackId)
	}
	if match, _ := regexp.MatchString(asciiRegex, passwd); !match {
		return model.NewExpectedError("password can't contains non-standard characters", "USER_INVALID", http.StatusBadRequest, trackId)
	}
	return nil
}
