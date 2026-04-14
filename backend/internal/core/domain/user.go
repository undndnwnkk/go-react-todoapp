package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID  `json:"id"`
	Name         string     `json:"name"`
	LastName     string     `json:"last_name"`
	Email        string     `json:"email"`
	DateOfBirth  *time.Time `json:"date_of_birth"`
	PasswordHash string     `json:"password_hash"`
	CreatedAt    time.Time  `json:"created_at"`
}

type UserCreateRequest struct {
	Name        string     `json:"name" validate:"required,min=2,max=30"`
	LastName    string     `json:"last_name" validate:"required,min=2,max=30"`
	Email       string     `json:"email" validate:"required,email"`
	Password    string     `json:"password" validate:"required,min=8,max=100"`
	DateOfBirth *time.Time `json:"date_of_birth" validate:"omitempty"`
}

type UserUpdateRequest struct {
	Name         string `json:"name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	DateOfBirth  string `json:"date_of_birth"`
}

type UserPatchRequest struct {
	Name         *string `json:"name"`
	LastName     *string `json:"last_name"`
	Email        *string `json:"email"`
	PasswordHash *string `json:"password_hash"`
	DateOfBirth  *string `json:"date_of_birth"`
}

type UserLoginRequest struct {
	Email    string
	Password string
}

type UserIdResponse struct {
	ID uuid.UUID `json:"id"`
}

func NewUser() *User {
	return &User{}
}
