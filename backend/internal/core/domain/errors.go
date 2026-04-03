package domain

import "errors"

var (
	UserDatabaseNotFoundError = errors.New("dont found user with this id")
	UserDatabaseAddError      = errors.New("cant add user to user database")
)
