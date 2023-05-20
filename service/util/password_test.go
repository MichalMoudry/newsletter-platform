package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ComparePasswordHash(t *testing.T) {
	type args struct {
		password string
		hash     string
	}
	tests := []struct {
		name      string
		args      args
		wantMatch bool
		wantErr   bool
	}{
		{
			name: "Test password with corresponding hash",
			args: args{
				password: "Password1",
				hash:     "$argon2id$v=19$m=65536,t=3,p=2$ZjgxOTQwM2UtNWJkZi00ZDczLWEyZjMtZWUwNWMyY2UzN2Iz$MuhmXORZzxzFMSMxj5HeRjroON5T0CICPLoGEYcjM5g",
			},
			wantMatch: true,
			wantErr:   false,
		},
		{
			name: "Test incorrect password against a hash",
			args: args{
				password: "Password1.",
				hash:     "$argon2id$v=19$m=65536,t=3,p=2$ZjgxOTQwM2UtNWJkZi00ZDczLWEyZjMtZWUwNWMyY2UzN2Iz$MuhmXORZzxzFMSMxj5HeRjroON5T0CICPLoGEYcjM5g",
			},
			wantMatch: false,
			wantErr:   false,
		},
		{
			name: "Test hash with incorrect metadata",
			args: args{
				password: "Password1",
				hash:     "$argon2id$v=18$m=65536,t=3,p=2$ZjgxOTQwM2UtNWJkZi00ZDczLWEyZjMtZWUwNWMyY2UzN2Iz$MuhmXORZzxzFMSMxj5HeRjroON5T0CICPLoGEYcjM5g",
			},
			wantMatch: false,
			wantErr:   true,
		},
		{
			name: "Test hash comparison with missing salt",
			args: args{
				password: "Password1",
				hash:     "$argon2id$v=19$m=65536,t=3,p=2$MuhmXORZzxzFMSMxj5HeRjroON5T0CICPLoGEYcjM5g",
			},
			wantMatch: false,
			wantErr:   true,
		},
		{
			name: "Test hash comparison with missing parameters",
			args: args{
				password: "Password1",
				hash:     "$argon2id$v=19$MuhmXORZzxzFMSMxj5HeRjroON5T0CICPLoGEYcjM5g",
			},
			wantMatch: false,
			wantErr:   true,
		},
		{
			name: "Test empty hash",
			args: args{
				password: "Password1",
				hash:     "",
			},
			wantMatch: false,
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMatch, err := ComparePasswordHash(tt.args.password, tt.args.hash)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, gotMatch, tt.wantMatch)
		})
	}
}
