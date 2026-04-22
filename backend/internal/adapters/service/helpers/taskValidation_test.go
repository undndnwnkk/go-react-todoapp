package helpers

import (
	"errors"
	"testing"

	"github.com/undndnwnkk/go-react-todoapp/internal/core/domain"
)

func TestValidateCreateRequest(t *testing.T) {
	tests := []struct {
		name      string
		request   domain.TaskCreateRequest
		wantErr   bool
		wantErrIs error
	}{
		{
			name:    "valid title",
			request: domain.TaskCreateRequest{Title: "Buy milk"},
			wantErr: false,
		},
		{
			name:      "empty title",
			request:   domain.TaskCreateRequest{Title: ""},
			wantErr:   true,
			wantErrIs: domain.ErrShortTitle,
		},
		{
			name:      "one char title",
			request:   domain.TaskCreateRequest{Title: "A"},
			wantErr:   true,
			wantErrIs: domain.ErrShortTitle,
		},
		{
			name:    "exactly two chars - boundary",
			request: domain.TaskCreateRequest{Title: "Ab"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateCreateRequest(tt.request)

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
