package util

import (
	"newsletter-platform/transport/model/contracts"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ContractValidation(t *testing.T) {
	type args struct {
		dto any
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test correct editor DTO validation",
			args: args{
				dto: contracts.RegisterUserRequestData{
					Email:    "test@test.com",
					UserName: "Test user 1",
					Password: "TestPass@1.2-",
				},
			},
			wantErr: false,
		},
		{
			name: "Test incorrect email in editor DTO",
			args: args{
				dto: contracts.RegisterUserRequestData{
					Email:    "test",
					UserName: "Test user 1",
					Password: "TestPass@1.2-",
				},
			},
			wantErr: true,
		},
		{
			name: "Test missing email in the editor DTO",
			args: args{
				dto: contracts.RegisterUserRequestData{
					UserName: "Test user 1",
					Password: "TestPass@1.2-",
				},
			},
			wantErr: true,
		},
		{
			name: "Test missing user name in editor DTO",
			args: args{
				dto: contracts.RegisterUserRequestData{
					Email:    "test@test.com",
					Password: "TestPass@1.2-",
				},
			},
			wantErr: true,
		},
		{
			name: "Test missing password in editor DTO",
			args: args{
				dto: contracts.RegisterUserRequestData{
					Email:    "test@test.com",
					UserName: "Test user 1",
				},
			},
			wantErr: true,
		},
		{
			name: "Test incorrect password length",
			args: args{
				dto: contracts.RegisterUserRequestData{
					Email:    "test@test.com",
					Password: "123",
					UserName: "Test user 1",
				},
			},
			wantErr: true,
		},
		// TODO: Add tests for checking validity of a password (e.g. length, characters included, ...)
		{
			name: "Test correct login request DTO",
			args: args{
				dto: contracts.LoginRequestData{
					Email:    "test@test.com",
					Password: "TestPass@1.2-",
				},
			},
			wantErr: false,
		},
		{
			name: "Test login request DTO with incorrect email",
			args: args{
				dto: contracts.LoginRequestData{
					Email: "test.com", Password: "TestPass@1.2-",
				},
			},
			wantErr: true,
		},
		{
			name: "Test login request DTO with incorrect password (incorrect length)",
			args: args{
				dto: contracts.LoginRequestData{
					Email: "test@test.com", Password: "test1.",
				},
			},
			wantErr: true,
		},
		{
			name: "Test password reset token request DTO",
			args: args{
				dto: contracts.PassResetTokenRequestData{
					Email: "test@test.com",
				},
			},
			wantErr: false,
		},
		{
			name: "Test password reset token request DTO with incorrect email",
			args: args{
				dto: contracts.PassResetTokenRequestData{
					Email: "test.com",
				},
			},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			isError := validate.Struct(test.args.dto) != nil
			assert.Equal(t, isError, test.wantErr)
		})
	}
}
