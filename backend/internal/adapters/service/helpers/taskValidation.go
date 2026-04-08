package helpers

import "github.com/undndnwnkk/go-react-todoapp/internal/core/domain"

func ValidateCreateRequest(request domain.TaskCreateRequest) error {
	if len(request.Title) < 2 {
		return domain.ErrShortTitle
	}

	return nil
}
