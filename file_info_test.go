package gopherpanic

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPositionSpawn(t *testing.T) {
	tests := []struct {
		name   string
		fields Position
		want   Position
	}{
		{
			name:   "OK",
			fields: Position{},
			want: Position{
				File: "file_info_test.go",
				Line: 28,
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := testCase.fields.Spawn() // Error check based on the current line
			files := strings.Split(result.File, "/")
			result.File = files[len(files)-1]
			assert.Equal(t, testCase.want, result)
		})
	}
}

func TestPositionPrivateSpawn(t *testing.T) {
	tests := []struct {
		name   string
		fields Position
		args   int
		want   Position
	}{
		{
			name:   "OK - lib spawn level",
			fields: Position{},
			args:   0,
			want: Position{
				File: "file_info.go",
				Line: 19,
			},
		},
		{
			name:   "OK - lib test spawn level",
			fields: Position{},
			args:   1,
			want: Position{
				File: "file_info_test.go",
				Line: 65,
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := testCase.fields.spawn(testCase.args) // Error check based on the current line
			files := strings.Split(result.File, "/")
			result.File = files[len(files)-1]
			assert.Equal(t, testCase.want, result)
		})
	}
}
