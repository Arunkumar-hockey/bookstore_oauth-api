package db

import (
	"auth/src/domain/access_token"
	"auth/src/utils/errors"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}

type dbRepository struct {
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	//TODO: implement get access token from CassandraDB.
	return nil, errors.NewInternalServerError("database connection not implemented yet")
}
