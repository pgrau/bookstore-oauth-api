package access_token

import (
	"github.com/pgrau/bookstore-oauth-api/lib/error"
	"github.com/pgrau/bookstore-oauth-api/src/repository/db"
	"github.com/pgrau/bookstore-oauth-api/src/repository/rest"
)

type Service interface {
	GetById(string) (*AccessToken, *error.RestErr)
	Create(request AccessTokenRequest) (*AccessToken, *error.RestErr)
	UpdateExpirationTime(AccessToken) *error.RestErr
}

type service struct {
	userRepository   rest.RestUserRepository
	accessRepository db.DbRepository
}

func NewService(userRepo rest.RestUserRepository, accessTokenRepo db.DbRepository) Service {
	return &service{
		userRepository:   userRepo,
		accessRepository: accessTokenRepo,
	}
}

func (s *service) GetById(accessTokenId string) (*AccessToken, *error.RestErr) {
	if len(accessTokenId) == 0 {
		return nil, error.BadRequest("Invalid access token id")
	}
	accessToken, err := s.accessRepository.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}

	return accessToken, nil
}

func (s *service) Create(request AccessTokenRequest) (*AccessToken, *error.RestErr) {
	if err := request.Validate(); err != nil {
		return nil, err
	}

	//TODO: Support both grant types: client_credentials and password

	// Authenticate the user against the Users API:
	user, err := s.userRepository.LoginUser(request.Username, request.Password)
	if err != nil {
		return nil, err
	}

	// Generate a new access token:
	at := GetNewAccessToken(user.Id)
	at.Generate()

	// Save the new access token in Cassandra:
	if err := s.accessRepository.Create(at); err != nil {
		return nil, err
	}
	return &at, nil
}

func (s *service) UpdateExpirationTime(at AccessToken) *error.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}

	return s.accessRepository.UpdateExpirationTime(at)
}