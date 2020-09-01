package access_token

import (
	"github.com/pgrau/bookstore-oauth-api/lib/error"
)

type Repository interface {
	GetById(string) (*AccessToken, *error.RestErr)
	Create(AccessToken) *error.RestErr
	UpdateExpirationTime(AccessToken) *error.RestErr
}

type Service interface {
	GetById(string) (*AccessToken, *error.RestErr)
	Create(AccessToken) *error.RestErr
	UpdateExpirationTime(AccessToken) *error.RestErr
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(accessTokenId string) (*AccessToken, *error.RestErr) {
	if len(accessTokenId) == 0 {
		return nil, error.BadRequest("Invalid access token id")
	}
	accessToken, err := s.repository.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}

	return accessToken, nil
}

func (s *service) Create(at AccessToken) *error.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}

	return s.repository.Create(at)
}

func (s *service) UpdateExpirationTime(at AccessToken) *error.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}

	return s.repository.UpdateExpirationTime(at)
}