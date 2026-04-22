package helpers

import (
	"errors"
	"testing"

	"github.com/undndnwnkk/go-react-todoapp/internal/core/domain"
)

func TestCheckCategoryCreateRequest(t *testing.T) {
	tests := []struct {
		name      string
		request   domain.CategoryCreateRequest
		wantErr   bool
		wantErrIs error
	}{
		{
			name: "valid request",
			request: domain.CategoryCreateRequest{
				Name:  "Work",
				Color: "#FF5733",
			},
			wantErr: false,
		},
		{
			name: "name exactly two chars - boundary",
			request: domain.CategoryCreateRequest{
				Name:  "Ab",
				Color: "#FF5733",
			},
			wantErr: false,
		},
		{
			name: "color exactly two chars - boundary",
			request: domain.CategoryCreateRequest{
				Name:  "Work",
				Color: "Ab",
			},
			wantErr: false,
		},
		{
			name: "empty name",
			request: domain.CategoryCreateRequest{
				Name:  "",
				Color: "#FF5733",
			},
			wantErr:   true,
			wantErrIs: domain.ErrShortCategoryName,
		},
		{
			name: "one char name",
			request: domain.CategoryCreateRequest{
				Name:  "W",
				Color: "#FF5733",
			},
			wantErr:   true,
			wantErrIs: domain.ErrShortCategoryName,
		},
		{
			name: "empty color",
			request: domain.CategoryCreateRequest{
				Name:  "Work",
				Color: "",
			},
			wantErr:   true,
			wantErrIs: domain.ErrShortCategoryColor,
		},
		{
			name: "one char color",
			request: domain.CategoryCreateRequest{
				Name:  "Work",
				Color: "#",
			},
			wantErr:   true,
			wantErrIs: domain.ErrShortCategoryColor,
		},
		{
			name: "name checked before color",
			request: domain.CategoryCreateRequest{
				Name:  "W", // оба невалидны —
				Color: "#", // должна вернуться ошибка name (первая)
			},
			wantErr:   true,
			wantErrIs: domain.ErrShortCategoryName,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CheckCategoryCreateRequest(tt.request)

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
