package util

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RemoveAllNewLineChars(t *testing.T) {
	type args struct {
		input  string
		osType string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test positive windows case",
			args: args{
				input:  "Test input\r\n",
				osType: "windows",
			},
			want: "Test input",
		},
		{
			name: "Test positive unix case",
			args: args{
				input:  "Test input\n",
				osType: "linux",
			},
			want: "Test input",
		},
	}

	runtimeOs := runtime.GOOS
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.osType == runtimeOs {
				got := RemoveAllNewLineChars(tt.args.input)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
