package rest

import (
	"encoding/json"
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/pgrau/bookstore-oauth-api/lib/error"
	"github.com/pgrau/bookstore-oauth-api/src/domain/user"
	"time"
)

var (
	httpClient = rest.RequestBuilder{
		BaseURL: "http://localhost:8081",
		Timeout: 100 * time.Millisecond,
	}
)

type RestUserRepository interface {
	LoginUser(string, string) (*user.User, *error.RestErr)
}

type userRepository struct{}

func NewRestUserRepository() RestUserRepository {
	return &userRepository{}
}

func (r *userRepository) LoginUser(email string, password string) (*user.User, *error.RestErr) {
	request := user.LoginRequest{
		Email:    email,
		Password: password,
	}

	response := httpClient.Post("/users/login", request)

	if response == nil || response.Response == nil {
		return nil, error.InternalServerError("invalid restclient response when trying to login user")
	}

	if response.StatusCode > 299 {
		var restErr error.RestErr
		err := json.Unmarshal(response.Bytes(), &restErr)
		if err != nil {
			return nil, error.InternalServerError("invalid error interface when trying to login user")
		}

		return nil, &restErr
	}

	var user user.User
	if err := json.Unmarshal(response.Bytes(), &user); err != nil {
		return nil, error.InternalServerError("error when trying to unmarshal users login response")
	}

	return &user, nil
}