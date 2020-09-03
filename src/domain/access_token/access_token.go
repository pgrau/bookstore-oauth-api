package access_token

import (
	"fmt"
	"github.com/pgrau/bookstore-oauth-api/lib/crypto"
	"github.com/pgrau/bookstore-oauth-api/lib/error"
	"strings"
	"time"
)

const (
	expirationTime             = 24
	grantTypePassword          = "password"
	grandTypeClientCredentials = "client_credentials"
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

type AccessTokenRequest struct {
	GrantType    string `json:"grant_type"`
	Scope        string `json:"scope"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (at *AccessTokenRequest) Validate() *error.RestErr {
	switch at.GrantType {
	case grantTypePassword:
	case grandTypeClientCredentials:
		break

	default:
		return error.BadRequest("invalid grant_type parameter")
	}

	return nil
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

func GetNewAccessToken(userId int64) AccessToken {
	return AccessToken{
		UserId:  userId,
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}

func (at *AccessToken) Generate() {
	at.AccessToken = crypto.GetMd5(fmt.Sprintf("at-%d-%d-ran", at.UserId, at.Expires))
}