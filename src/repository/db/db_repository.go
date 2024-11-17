package db

import (
	accesstoken "BOOKSTORE_OAUTH-API/src/domain/access_token"
	"BOOKSTORE_OAUTH-API/src/utils/errors"
)

func New() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*accesstoken.AccessToken, *errors.RestErr)
}

type dbRepository struct {
}

func (r *dbRepository) GetById(id string) (*accesstoken.AccessToken, *errors.RestErr) {
	return nil, errors.NewInternalServerError("database connection not implemented yet")
}
