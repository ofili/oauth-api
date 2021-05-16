package db

import (
	"github.com/ofililewis/bookstore_oauth-api/src/domain/access_token"
	"github.com/ofililewis/bookstore_oauth-api/src/utils/errors"
)

func NewRepository() Repository {
	return &dbRepository{}
}

type Repository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct {

}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	//TODO: implement get access token from Cassandra DB

	return nil, errors.NewInternalServerError("database connection not implemented yet")
	// return nil, nil
}