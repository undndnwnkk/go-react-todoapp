package helpers

import (
	"strings"
	"time"

	"github.com/undndnwnkk/go-react-todoapp/internal/core/domain"
)

func CheckEmail(email string) error {
	if !strings.Contains(email, "@") {
		return domain.ErrInvalidEmail
	}

	return nil
}

func CheckPasswordLength(password string) error {
	if len(password) < 8 {
		return domain.ErrShortPassword
	}
	return nil
}

func CheckDateOfBirth(dateOfBirth time.Time) error {

	min := time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)
	now := time.Now()
	max := time.Date(now.Year()+1, 1, 1, 0, 0, 0, 0, time.UTC)

	if dateOfBirth.Before(min) || !dateOfBirth.Before(max) {
		return domain.ErrInvalidDateOfBirth
	}

	return nil
}

func CheckUserCreateRequest(request domain.UserCreateRequest) error {
	if err := CheckEmail(request.Email); err != nil {
		return err
	}

	if err := CheckPasswordLength(request.Password); err != nil {
		return err
	}

	if request.DateOfBirth != nil {
		if err := CheckDateOfBirth(*request.DateOfBirth); err != nil {
			return err
		}
	}

	return nil
}
