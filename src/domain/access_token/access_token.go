package access_token

import (
	"github.com/pgrau/bookstore-oauth-api/lib/error"
	"strings"
	"time"
)

const (
	expirationTime = 24
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func (at *AccessToken) Validate() *error.RestErr {
	if len(strings.TrimSpace(at.AccessToken)) == 0 {
		return error.BadRequest("invalid access token id")
	}

	if at.UserId <= 0 {
		return error.BadRequest("invalid user id")
	}

	if at.ClientId <= 0 {
		return error.BadRequest("invalid client id")
	}

	if at.Expires <= 0 {
		return error.BadRequest("invalid expiration time")
	}

	return nil;
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}