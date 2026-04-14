package domain

import "errors"

var (
	// Service errors
	ErrInvalidEmail       = errors.New("invalid email")
	ErrShortPassword      = errors.New("short password")
	ErrInvalidDateOfBirth = errors.New("invalid date of birth")
	ErrInvalidPassword    = errors.New("invalid password")
	ErrUserAlreadyExists  = errors.New("user already exists")

	ErrShortTitle = errors.New("short title")

	ErrShortCategoryName  = errors.New("invalid name")
	ErrShortCategoryColor = errors.New("invalid color")

	// Database errors
	ErrUserNotFound  = errors.New("user not found")
	ErrAddToDatabase = errors.New("cant add user to user database")

	ErrCategoryNotFound  = errors.New("dont found category with this id")
	ErrCategoryForbidden = errors.New("you dont have access to this category")

	ErrTaskNotFound  = errors.New("task not found")
	ErrTaskForbidden = errors.New("you dont have access to this task")

	// Env errors
	ErrEnvNotFound    = errors.New(".env not found")
	ErrShortJwtSecret = errors.New("jwt secret is too small")
	ErrJwtAccessLong  = errors.New("jwt access ttl must be less than 15 minutes")
	ErrJwtRefreshLong = errors.New("jwt refresh ttl must be less than 30 days")
)
