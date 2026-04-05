package domain

import "errors"

var (
	// Database errors
	ErrUserNotFound  = errors.New("dont found user with this id")
	ErrAddToDatabase = errors.New("cant add user to user database")

	ErrCategoryNotFound  = errors.New("dont found category with this id")
	ErrCategoryForbidden = errors.New("you dont have access to this category")

	ErrTaskNotFound  = errors.New("dont found task with this id")
	ErrTaskForbidden = errors.New("you dont have access to this task")

	// Env errors
	ErrEnvNotFound    = errors.New(".env not found")
	ErrShortJwtSecret = errors.New("jwt secret is too small")
	ErrJwtAccessLong  = errors.New("jwt access ttl must be less than 15 minutes")
	ErrJwtRefreshLong = errors.New("jwt refresh ttl must be less than 30 days")
)
