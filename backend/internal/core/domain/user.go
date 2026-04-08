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
	Name        string     `json:"name"`
	LastName    string     `json:"last_name"`
	Email       string     `json:"email"`
	Password    string     `json:"password"`
	DateOfBirth *time.Time `json:"date_of_birth"`
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
