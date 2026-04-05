package domain

import "errors"

var (
	// Database errors
	UserDatabaseNotFoundError = errors.New("dont found user with this id")
	UserDatabaseAddError      = errors.New("cant add user to user database")

	// Env errors
	EnvNotFoundError    = errors.New(".env not found")
	JwtTooShortError    = errors.New("jwt secret is too small")
	JwtAccessLongError  = errors.New("jwt access ttl must be less than 15 minutes")
	JwtRefreshLongError = errors.New("jwt refresh ttl must be less than 30 days")
)
