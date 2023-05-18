package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	configPath string = "../config.json"
)

func Test_ReadConfigFromFile(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    Config
		wantErr bool
	}{
		{
			name: "Correct config path",
			args: args{
				path: configPath,
			},
			want: Config{
				Port:             8443,
				ConnectionString: "postgres://root:root@localhost:5432/data-persistence?sslmode=disable",
				Environment:      Dev,
			},
			wantErr: false,
		},
		{
			name: "Incorrect config path",
			args: args{
				path: "configuration.json",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadConfigFromFile(tt.args.path)
			//assert.NotEqual(t, err != nil, tt.wantErr)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadConfigFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_EnvironmentConfig(t *testing.T) {
	type args struct {
		configPath string
		Env        Environment
	}

	tests := []struct {
		name string
		args args
		want Config
	}{
		{
			name: "Test config in local dev environment",
			args: args{
				configPath: configPath,
				Env:        Dev,
			},
			want: Config{
				Port:             8443,
				ConnectionString: "postgres://root:root@localhost:5432/data-persistence?sslmode=disable",
				Environment:      Dev,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := ReadConfigFromFile(test.args.configPath)
			assert.Nil(t, err)
			assert.Equal(t, test.want, got)
		})
	}
}
