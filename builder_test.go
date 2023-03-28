package gopherpanic

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorBuilderNew(t *testing.T) {
	tests := []struct {
		name   string
		fields ErrorBuilder
		want   ErrorBuilder
	}{
		{
			name:   "OK",
			fields: ErrorBuilder{},
			want:   ErrorBuilder{},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := ErrorBuilder{}.New()
			assert.Equal(t, testCase.want, result)
		})
	}
}

func TestErrorBuilderDefault(t *testing.T) {
	tests := []struct {
		name   string
		fields ErrorBuilder
		want   ErrorBuilder
	}{
		{
			name:   "OK",
			fields: ErrorBuilder{},
			want: ErrorBuilder{
				message:  "an unexpected error occured",
				position: Position{File: "builder_test.go", Line: 49},
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := testCase.fields.Default() // Error check based on the current line
			files := strings.Split(result.position.File, "/")
			result.position.File = files[len(files)-1]
			assert.Equal(t, testCase.want, result)
		})
	}
}

func TestErrorBuilderWithMessage(t *testing.T) {
	tests := []struct {
		name   string
		fields ErrorBuilder
		args   string
		want   ErrorBuilder
	}{
		{
			name:   "OK",
			fields: ErrorBuilder{},
			args:   "sample error",
			want:   ErrorBuilder{message: "sample error"},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := ErrorBuilder{}.WithMessage(testCase.args)
			assert.Equal(t, testCase.want, result)
		})
	}
}

func TestErrorBuilderWithPosition(t *testing.T) {
	tests := []struct {
		name   string
		fields ErrorBuilder
		args   Position
		want   ErrorBuilder
	}{
		{
			name:   "OK",
			fields: ErrorBuilder{},
			args:   Position{File: "sample.go", Line: 50},
			want:   ErrorBuilder{position: Position{File: "sample.go", Line: 50}},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := ErrorBuilder{}.WithPosition(testCase.args)
			assert.Equal(t, testCase.want, result)
		})
	}
}

func TestErrorBuilderWithTraces(t *testing.T) {
	tests := []struct {
		name   string
		fields ErrorBuilder
		args   []Trace
		want   ErrorBuilder
	}{
		{
			name:   "OK",
			fields: ErrorBuilder{},
			args: []Trace{
				{
					Message: "sample 1",
					Position: Position{
						File: "sample_1.go",
						Line: 10,
					},
				},
				{
					Message: "sample 2",
					Position: Position{
						File: "sample_2.go",
						Line: 20,
					},
				},
				{
					Message: "sample 3",
					Position: Position{
						File: "sample_3.go",
						Line: 30,
					},
				},
			},
			want: ErrorBuilder{traces: []Trace{
				{
					Message: "sample 1",
					Position: Position{
						File: "sample_1.go",
						Line: 10,
					},
				},
				{
					Message: "sample 2",
					Position: Position{
						File: "sample_2.go",
						Line: 20,
					},
				},
				{
					Message: "sample 3",
					Position: Position{
						File: "sample_3.go",
						Line: 30,
					},
				},
			}},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := ErrorBuilder{}.WithTraces(testCase.args...)
			assert.Equal(t, testCase.want, result)
		})
	}
}

func TestErrorBuilderBuild(t *testing.T) {
	tests := []struct {
		name   string
		fields ErrorBuilder
		want   Error
	}{
		{
			name: "OK",
			fields: ErrorBuilder{
				message:  "sample error",
				position: Position{File: "sample.go", Line: 50},
				traces: []Trace{
					{
						Message: "sample 1",
						Position: Position{
							File: "sample_1.go",
							Line: 10,
						},
					},
					{
						Message: "sample 2",
						Position: Position{
							File: "sample_2.go",
							Line: 20,
						},
					},
					{
						Message: "sample 3",
						Position: Position{
							File: "sample_3.go",
							Line: 30,
						},
					},
				},
			},
			want: Error{
				Message:  "sample error",
				Position: Position{File: "sample.go", Line: 50},
				Traces: []Trace{
					{
						Message: "sample 1",
						Position: Position{
							File: "sample_1.go",
							Line: 10,
						},
					},
					{
						Message: "sample 2",
						Position: Position{
							File: "sample_2.go",
							Line: 20,
						},
					},
					{
						Message: "sample 3",
						Position: Position{
							File: "sample_3.go",
							Line: 30,
						},
					},
				},
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := testCase.fields.Build()
			assert.Equal(t, testCase.want, result)
		})
	}
}
