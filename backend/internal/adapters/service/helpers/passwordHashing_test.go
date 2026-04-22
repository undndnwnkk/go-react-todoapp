package helpers

import "testing"

func TestHashPassword(t *testing.T) {
	password := "12345678acbde"
	hashed, _ := HashPassword(password)
	if hashed == "" {
		t.Error("password don't hashed")
	}

	if hashed == password {
		t.Error("password and hash are equal")
	}
}

func TestCheckPasswordHash(t *testing.T) {
	hashedBasePassword, _ := HashPassword("12345678abcde")

	tests := []struct {
		name     string
		password string
		wantErr  bool
	}{
		{"correct password", "12345678abcde", false},
		{"incorrect password", "09876543sjkd", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isEqual := CheckPasswordHash(tt.password, hashedBasePassword)
			if isEqual == tt.wantErr {
				t.Errorf("got res=%v, wantErr=%v", isEqual, tt.wantErr)
			}
		})
	}
}
