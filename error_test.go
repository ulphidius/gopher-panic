package gopherpanic

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	type args struct {
		message string
		traces  []Trace
	}

	tests := []struct {
		name string
		args args
		want Error
	}{
		{
			name: "OK",
			args: args{
				message: "sample error",
				traces: []Trace{
					{
						Message: "inner 1",
						Position: Position{
							File: "inner_1.go",
							Line: 10,
						},
					},
					{
						Message: "inner 2",
						Position: Position{
							File: "inner_2.go",
							Line: 20,
						},
					},
					{
						Message: "inner 3",
						Position: Position{
							File: "inner_3.go",
							Line: 30,
						},
					},
				},
			},
			want: Error{
				Message: "sample error",
				Position: Position{
					File: "error_test.go",
					Line: 97,
				},
				Traces: []Trace{
					{
						Message: "inner 1",
						Position: Position{
							File: "inner_1.go",
							Line: 10,
						},
					},
					{
						Message: "inner 2",
						Position: Position{
							File: "inner_2.go",
							Line: 20,
						},
					},
					{
						Message: "inner 3",
						Position: Position{
							File: "inner_3.go",
							Line: 30,
						},
					},
				},
			},
		},
		{
			name: "nil traces",
			args: args{
				message: "sample error",
			},
			want: Error{
				Message: "sample error",
				Position: Position{
					File: "error_test.go",
					Line: 97,
				},
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := New(testCase.args.message, testCase.args.traces) // Error check based on the current line
			files := strings.Split(result.Position.File, "/")
			result.Position.File = files[len(files)-1]
			assert.Equal(t, testCase.want, result)
		})
	}
}

func TestWrap(t *testing.T) {
	type args struct {
		message string
		err     Error
	}

	tests := []struct {
		name string
		args args
		want Error
	}{
		{
			name: "OK",
			args: args{
				message: "sample error",
				err: Error{
					Message: "error message",
					Position: Position{
						File: "error.go",
						Line: 10,
					},
					Traces: []Trace{
						{
							Message: "inner 1",
							Position: Position{
								File: "inner_1.go",
								Line: 10,
							},
						},
						{
							Message: "inner 2",
							Position: Position{
								File: "inner_2.go",
								Line: 20,
							},
						},
						{
							Message: "inner 3",
							Position: Position{
								File: "inner_3.go",
								Line: 30,
							},
						},
					},
				},
			},
			want: Error{
				Message: "sample error",
				Position: Position{
					File: "error_test.go",
					Line: 193,
				},
				Traces: []Trace{
					{
						Message: "error message",
						Position: Position{
							File: "error.go",
							Line: 10,
						},
					},
					{
						Message: "inner 1",
						Position: Position{
							File: "inner_1.go",
							Line: 10,
						},
					},
					{
						Message: "inner 2",
						Position: Position{
							File: "inner_2.go",
							Line: 20,
						},
					},
					{
						Message: "inner 3",
						Position: Position{
							File: "inner_3.go",
							Line: 30,
						},
					},
				},
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := Wrap(testCase.args.message, testCase.args.err) // Error check based on the current line
			files := strings.Split(result.Position.File, "/")
			result.Position.File = files[len(files)-1]
			assert.Equal(t, testCase.want, result)
		})
	}
}

func TestErrorIntoTrace(t *testing.T) {
	tests := []struct {
		name   string
		fields Error
		want   Trace
	}{
		{
			name: "OK",
			fields: Error{
				Message: "sample error",
				Position: Position{
					File: "sample.go",
					Line: 50,
				},
			},
			want: Trace{
				Message: "sample error",
				Position: Position{
					File: "sample.go",
					Line: 50,
				},
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := testCase.fields.IntoTrace()
			assert.Equal(t, testCase.want, result)
		})
	}
}

func TestErrorError(t *testing.T) {
	tests := []struct {
		name   string
		fields Error
		want   string
	}{
		{
			name: "OK",
			fields: Error{
				Message: "sample error",
				Position: Position{
					File: "sample.go",
					Line: 50,
				},
			},
			want: "error: sample error; in file: sample.go; at line: 50",
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := testCase.fields.Error()
			assert.Equal(t, testCase.want, result)
		})
	}
}

func TestErrorFormat(t *testing.T) {
	tests := []struct {
		name   string
		fields Error
		want   string
	}{
		{
			name: "OK",
			fields: Error{
				Message: "sample error",
				Position: Position{
					File: "sample.go",
					Line: 50,
				},
			},
			want: "error: sample error; in file: sample.go; at line: 50",
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := testCase.fields.Format()
			assert.Equal(t, testCase.want, result)
		})
	}
}

func TestErrorFormatWithTraces(t *testing.T) {
	tests := []struct {
		name   string
		fields Error
		want   string
	}{
		{
			name: "OK",
			fields: Error{
				Message: "sample error",
				Position: Position{
					File: "sample.go",
					Line: 50,
				},
				Traces: []Trace{
					{
						Message: "inner 1",
						Position: Position{
							File: "inner_1.go",
							Line: 10,
						},
					},
					{
						Message: "inner 2",
						Position: Position{
							File: "inner_2.go",
							Line: 20,
						},
					},
					{
						Message: "inner 3",
						Position: Position{
							File: "inner_3.go",
							Line: 30,
						},
					},
				},
			},
			want: "error: sample error; in file: sample.go; at line: 50\n\terror: inner 1; in file: inner_1.go; at line: 10\n\terror: inner 2; in file: inner_2.go; at line: 20\n\terror: inner 3; in file: inner_3.go; at line: 30",
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := testCase.fields.FormatWithTraces()
			assert.Equal(t, testCase.want, result)
		})
	}
}

func TestErrorFormatJSON(t *testing.T) {
	tests := []struct {
		name   string
		fields Error
		want   string
	}{
		{
			name: "OK",
			fields: Error{
				Message: "sample error",
				Position: Position{
					File: "sample.go",
					Line: 50,
				},
				Traces: []Trace{
					{
						Message: "inner 1",
						Position: Position{
							File: "inner_1.go",
							Line: 10,
						},
					},
					{
						Message: "inner 2",
						Position: Position{
							File: "inner_2.go",
							Line: 20,
						},
					},
					{
						Message: "inner 3",
						Position: Position{
							File: "inner_3.go",
							Line: 30,
						},
					},
				},
			},
			want: "{\"message\":\"sample error\",\"position\":{\"file\":\"sample.go\",\"line\":50},\"traces\":[{\"message\":\"inner 1\",\"position\":{\"file\":\"inner_1.go\",\"line\":10}},{\"message\":\"inner 2\",\"position\":{\"file\":\"inner_2.go\",\"line\":20}},{\"message\":\"inner 3\",\"position\":{\"file\":\"inner_3.go\",\"line\":30}}]}",
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := testCase.fields.FormatJSON()
			assert.Equal(t, testCase.want, result)
		})
	}
}

func TestTraceIntoError(t *testing.T) {
	tests := []struct {
		name   string
		fields Trace
		want   Error
	}{
		{
			name: "OK",
			fields: Trace{
				Message: "sample error",
				Position: Position{
					File: "sample.go",
					Line: 50,
				},
			},
			want: Error{
				Message: "sample error",
				Position: Position{
					File: "sample.go",
					Line: 50,
				},
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := testCase.fields.IntoError()
			assert.Equal(t, testCase.want, result)
		})
	}
}

func TestTraceFormat(t *testing.T) {
	tests := []struct {
		name   string
		fields Trace
		want   string
	}{
		{
			name: "OK",
			fields: Trace{
				Message: "sample error",
				Position: Position{
					File: "sample.go",
					Line: 50,
				},
			},
			want: "error: sample error; in file: sample.go; at line: 50",
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := testCase.fields.Format()
			assert.Equal(t, testCase.want, result)
		})
	}
}
