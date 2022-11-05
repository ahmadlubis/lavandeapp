package response

import "time"

type UserAccessTokenResponse struct {
	AccessToken string    `json:"access_token"`
	ExpiredAt   time.Time `json:"expired_at"`
}
