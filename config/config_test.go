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
				SecurityString:   "132456789-/ef-/eqw-f",
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
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadConfigFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
