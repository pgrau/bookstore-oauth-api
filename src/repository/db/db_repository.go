package db

import (
	"fmt"
	"github.com/gocql/gocql"
	"github.com/pgrau/bookstore-oauth-api/lib/error"
	"github.com/pgrau/bookstore-oauth-api/src/client/cassandra"
	"github.com/pgrau/bookstore-oauth-api/src/domain/access_token"
)

const (
	queryGetAccessToken    = "SELECT access_token, user_id, client_id, expires FROM access_token WHERE access_token=?"
	queryCreateAccessToken = "INSERT INTO access_token(access_token, user_id, client_id, expires) VALUES (?, ?, ?, ?);"
	queryUpdateExpires     = "UPDATE access_token SET expires=? WHERE access_token=?;"
)

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *error.RestErr)
	Create(access_token.AccessToken) *error.RestErr
	UpdateExpirationTime(access_token.AccessToken) *error.RestErr
}

type dbRepository struct {
}

func NewRepository() DbRepository {
	return &dbRepository{}
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *error.RestErr) {
	var result access_token.AccessToken
	if err := cassandra.GetSession().Query(queryGetAccessToken, id).Scan(
		&result.AccessToken,
		&result.UserId,
		&result.ClientId,
		&result.Expires,
	); err != nil {
		if err == gocql.ErrNotFound {
			return nil, error.NotFound(fmt.Sprintf("no access token found with given id (%s)", id))
		}
		return nil, error.InternalServerError(err.Error())
	}
	return &result, nil
}

func (r *dbRepository) Create(at access_token.AccessToken) *error.RestErr {
	if err := cassandra.GetSession().Query(queryCreateAccessToken,
		&at.AccessToken,
		&at.UserId,
		&at.ClientId,
		&at.Expires,
	).Exec(); err != nil {
		return error.InternalServerError("error when trying to save access token in database")
	}
	return nil
}

func (r *dbRepository) UpdateExpirationTime(at access_token.AccessToken) *error.RestErr {
	if err := cassandra.GetSession().Query(queryUpdateExpires,
		&at.Expires,
		&at.AccessToken,
	).Exec(); err != nil {
		return error.InternalServerError("error when trying to update current resource")
	}
	return nil
}




