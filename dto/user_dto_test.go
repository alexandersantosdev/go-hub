package dto_test

import (
	"testing"

	"github.com/alexandersantosdev/go-hub/dto"
	"github.com/stretchr/testify/assert"
)

func TestValidateUser(t *testing.T) {
	tests := []struct {
		name    string
		input   dto.UserDTO
		wantErr bool
	}{
		{
			name: "Valid user",
			input: dto.UserDTO{
				Name:            "John Doe",
				Email:           "john@example.com",
				Password:        "secret123",
				ConfirmPassword: "secret123",
			},
			wantErr: false,
		},
		{
			name: "Missing name",
			input: dto.UserDTO{
				Email:           "john@example.com",
				Password:        "secret123",
				ConfirmPassword: "secret123",
			},
			wantErr: true,
		},
		{
			name: "Invalid email",
			input: dto.UserDTO{
				Name:            "John",
				Email:           "invalid-email",
				Password:        "secret123",
				ConfirmPassword: "secret123",
			},
			wantErr: true,
		},
		{
			name: "Password mismatch",
			input: dto.UserDTO{
				Name:            "John",
				Email:           "john@example.com",
				Password:        "secret123",
				ConfirmPassword: "wrongpass",
			},
			wantErr: true,
		},
		{
			name: "Short password",
			input: dto.UserDTO{
				Name:            "John",
				Email:           "john@example.com",
				Password:        "123",
				ConfirmPassword: "123",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.input.ValidateUser()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestValidateUserLogin(t *testing.T) {
	tests := []struct {
		name    string
		input   dto.UserLoginDTO
		wantErr bool
	}{
		{
			name: "Valid login",
			input: dto.UserLoginDTO{
				Email:    "user@example.com",
				Password: "password123",
			},
			wantErr: false,
		},
		{
			name: "Missing email",
			input: dto.UserLoginDTO{
				Password: "password123",
			},
			wantErr: true,
		},
		{
			name: "Missing password",
			input: dto.UserLoginDTO{
				Email: "user@example.com",
			},
			wantErr: true,
		},
		{
			name: "Invalid email format",
			input: dto.UserLoginDTO{
				Email:    "invalid",
				Password: "password123",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.input.ValidateUserLogin()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
