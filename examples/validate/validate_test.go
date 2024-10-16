package main

import (
	"testing"

	userpb "buf/example/gen/user/v1"
	"github.com/bufbuild/protovalidate-go"
)

func TestUserValidation(t *testing.T) {
	validator, err := protovalidate.New()
	if err != nil {
		t.Fatalf("failed to create validator: %v", err)
	}

	tests := []struct {
		name    string
		user    *userpb.User
		wantErr bool
	}{
		{
			name: "Valid user",
			user: &userpb.User{
				Id:       "1",
				Name:     "Dilmurat",
				Email:    "dilmurat@example.com",
				Password: "password123",
			},
			wantErr: false,
		},
		{
			name: "Invalid user with empty name",
			user: &userpb.User{
				Id:       "2",
				Name:     "",
				Email:    "validemail@example.com",
				Password: "password123",
			},
			wantErr: true,
		},
		{
			name: "Invalid user with bad email",
			user: &userpb.User{
				Id:       "3",
				Name:     "ValidName",
				Email:    "invalid-email",
				Password: "password123",
			},
			wantErr: true,
		},
		{
			name: "Invalid user with short password",
			user: &userpb.User{
				Id:       "4",
				Name:     "ValidName",
				Email:    "validemail@example.com",
				Password: "short",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.Validate(tt.user)

			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
