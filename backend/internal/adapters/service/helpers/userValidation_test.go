// internal/adapters/service/helpers/userValidation_test.go
package helpers

import (
	"errors"
	"testing"
	"time"

	"github.com/undndnwnkk/go-react-todoapp/internal/core/domain"
)

func TestCheckEmail(t *testing.T) {
	tests := []struct {
		name    string
		email   string
		wantErr bool
	}{
		{"valid email", "user@example.com", false},
		{"missing @", "userexample.com", true},
		{"empty email", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CheckEmail(tt.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("got err=%v, wantErr=%v", err, tt.wantErr)
			}
		})
	}
}

func TestCheckPasswordLength(t *testing.T) {
	tests := []struct {
		name     string
		password string
		wantErr  bool
	}{
		{"valid password", "12345678", false},
		{"short password", "1234", true},
		{"empty password", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CheckPasswordLength(tt.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("got err=%v, wantErr=%v", err, tt.wantErr)
			}
		})
	}
}

func TestCheckDateOfBirth(t *testing.T) {
	tests := []struct {
		name        string
		dateOfBirth time.Time
		wantErr     bool
	}{
		{"valid date", time.Date(2026, time.April, 1, 1, 1, 1, 1, time.UTC), false},
		{"date too old", time.Date(1899, time.April, 1, 1, 1, 1, 1, time.UTC), true},
		{"date in future", time.Now().AddDate(1, 0, 0), true},
		{"empty date", time.Time{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CheckDateOfBirth(tt.dateOfBirth)
			if (err != nil) != tt.wantErr {
				t.Errorf("got err=%v, wantErr=%v", err, tt.wantErr)
			}
		})
	}
}

func TestCheckUserCreateRequest(t *testing.T) {
	dob := func(year int) *time.Time {
		t := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
		return &t
	}

	tests := []struct {
		name      string
		request   domain.UserCreateRequest
		wantErr   bool
		wantErrIs error
	}{
		{
			name: "valid request without dob",
			request: domain.UserCreateRequest{
				Name:     "Lev",
				LastName: "Bykov",
				Email:    "lev@example.com",
				Password: "strongpass123",
			},
			wantErr: false,
		},
		{
			name: "valid request with dob",
			request: domain.UserCreateRequest{
				Name:        "Lev",
				LastName:    "Bykov",
				Email:       "lev@example.com",
				Password:    "strongpass123",
				DateOfBirth: dob(1995),
			},
			wantErr: false,
		},
		{
			name: "invalid email",
			request: domain.UserCreateRequest{
				Name:     "Lev",
				LastName: "Bykov",
				Email:    "notanemail",
				Password: "strongpass123",
			},
			wantErr:   true,
			wantErrIs: domain.ErrInvalidEmail,
		},
		{
			name: "short password",
			request: domain.UserCreateRequest{
				Name:     "Lev",
				LastName: "Bykov",
				Email:    "lev@example.com",
				Password: "123",
			},
			wantErr:   true,
			wantErrIs: domain.ErrShortPassword,
		},
		{
			name: "invalid dob - too old",
			request: domain.UserCreateRequest{
				Name:        "Lev",
				LastName:    "Bykov",
				Email:       "lev@example.com",
				Password:    "strongpass123",
				DateOfBirth: dob(1800),
			},
			wantErr:   true,
			wantErrIs: domain.ErrInvalidDateOfBirth,
		},
		{
			name: "invalid dob - future",
			request: domain.UserCreateRequest{
				Name:        "Lev",
				LastName:    "Bykov",
				Email:       "lev@example.com",
				Password:    "strongpass123",
				DateOfBirth: dob(time.Now().Year() + 2),
			},
			wantErr:   true,
			wantErrIs: domain.ErrInvalidDateOfBirth,
		},
		{
			name: "email checked before password",
			request: domain.UserCreateRequest{
				Name:     "Lev",
				LastName: "Bykov",
				Email:    "notanemail",
				Password: "123",
			},
			wantErr:   true,
			wantErrIs: domain.ErrInvalidEmail,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CheckUserCreateRequest(tt.request)

			if (err != nil) != tt.wantErr {
				t.Errorf("got err=%v, wantErr=%v", err, tt.wantErr)
				return
			}
			if tt.wantErrIs != nil && !errors.Is(err, tt.wantErrIs) {
				t.Errorf("got err=%v, want err=%v", err, tt.wantErrIs)
			}
		})
	}
}
