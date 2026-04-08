package helpers

import "github.com/undndnwnkk/go-react-todoapp/internal/core/domain"

func CheckCategoryCreateRequest(request domain.CategoryCreateRequest) error {
	if len(request.Name) < 2 {
		return domain.ErrShortCategoryName
	}

	if len(request.Color) < 2 {
		return domain.ErrShortCategoryColor
	}

	return nil
}
