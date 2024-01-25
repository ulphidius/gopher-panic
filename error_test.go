package gopherpanic

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleNew() {
	err := New(InternalError, "message fail to compute the statistics")
	filename_without_path := strings.Split(err.Position.File, "/")
	err.Position.File = filename_without_path[len(filename_without_path)-1]
	d, _ := json.Marshal(err)
	fmt.Println(string(d))
	// Output: {"code":{"id":3,"description":"failed to perform application task"},"message":"message fail to compute the statistics","position":{"file":"error_test.go","line":13}}
}

func ExampleWrap() {
	err := New(InternalError, "message fail to compute the statistics")
	filename_without_path := strings.Split(err.Position.File, "/")
	err.Position.File = filename_without_path[len(filename_without_path)-1]

	newErr := Wrap(InternalError, "fail to fetch statistics data", err)
	filename_without_path = strings.Split(newErr.Position.File, "/")
	newErr.Position.File = filename_without_path[len(filename_without_path)-1]
	d, _ := json.Marshal(newErr)
	fmt.Println(string(d))
	// Output: {"code":{"id":3,"description":"failed to perform application task"},"message":"fail to fetch statistics data","position":{"file":"error_test.go","line":26},"traces":[{"message":"message fail to compute the statistics","position":{"file":"error_test.go","line":22}}]}
}

func ExampleError_IntoTrace() {
	trace := New(InternalError, "message fail to compute the statistics").IntoTrace()
	filename_without_path := strings.Split(trace.Position.File, "/")
	trace.Position.File = filename_without_path[len(filename_without_path)-1]
	d, _ := json.Marshal(trace)
	fmt.Println(string(d))
	// Output: {"message":"message fail to compute the statistics","position":{"file":"error_test.go","line":35}}

}

func ExampleError_Error() {
	err := New(InternalError, "message fail to compute the statistics")
	filename_without_path := strings.Split(err.Position.File, "/")
	err.Position.File = filename_without_path[len(filename_without_path)-1]
	fmt.Println(err.Error())
	// Output: error_test.go:45: Error: 3:failed to perform application task:message fail to compute the statistics
}

func ExampleError_Format() {
	err := New(InternalError, "message fail to compute the statistics")
	filename_without_path := strings.Split(err.Position.File, "/")
	err.Position.File = filename_without_path[len(filename_without_path)-1]

	newErr := Wrap(InternalError, "fail to fetch statistics data", err)
	filename_without_path = strings.Split(newErr.Position.File, "/")
	newErr.Position.File = filename_without_path[len(filename_without_path)-1]

	fmt.Println(newErr.Format(true, true))
	// Output:
	// code id: 3; description: failed to perform application task
	// 	error message: fail to fetch statistics data; in file: error_test.go; at line: 57
}

func ExampleError_Format_withoutInnerData() {
	err := New(InternalError, "message fail to compute the statistics")
	filename_without_path := strings.Split(err.Position.File, "/")
	err.Position.File = filename_without_path[len(filename_without_path)-1]

	newErr := Wrap(InternalError, "fail to fetch statistics data", err)
	filename_without_path = strings.Split(newErr.Position.File, "/")
	newErr.Position.File = filename_without_path[len(filename_without_path)-1]

	fmt.Println(newErr.Format(true, false))
	// Output:
	// code id: 3; description: failed to perform application task
	//	error message: fail to fetch statistics data
}

func ExampleError_Format_GNUWithInnerData() {
	err := New(InternalError, "message fail to compute the statistics")
	filename_without_path := strings.Split(err.Position.File, "/")
	err.Position.File = filename_without_path[len(filename_without_path)-1]

	newErr := Wrap(InternalError, "fail to fetch statistics data", err)
	filename_without_path = strings.Split(newErr.Position.File, "/")
	newErr.Position.File = filename_without_path[len(filename_without_path)-1]

	fmt.Println(newErr.Format(false, true))
	// Output: error_test.go:87: Error: 3:failed to perform application task:fail to fetch statistics data
}

func ExampleError_Format_GNUWithoutInnerData() {
	err := New(InternalError, "message fail to compute the statistics")
	filename_without_path := strings.Split(err.Position.File, "/")
	err.Position.File = filename_without_path[len(filename_without_path)-1]

	newErr := Wrap(InternalError, "fail to fetch statistics data", err)
	filename_without_path = strings.Split(newErr.Position.File, "/")
	newErr.Position.File = filename_without_path[len(filename_without_path)-1]

	fmt.Println(newErr.Format(false, false))
	// Output: Error: 3:failed to perform application task:fail to fetch statistics data
}

func ExampleError_FormatWithTraces() {
	err := New(InternalError, "message fail to compute the statistics")
	filename_without_path := strings.Split(err.Position.File, "/")
	err.Position.File = filename_without_path[len(filename_without_path)-1]

	newErr := Wrap(InternalError, "fail to fetch statistics data", err)
	filename_without_path = strings.Split(newErr.Position.File, "/")
	newErr.Position.File = filename_without_path[len(filename_without_path)-1]

	newErr2 := Wrap(InternalError, "fail to fetch statistics data", newErr)
	filename_without_path = strings.Split(newErr2.Position.File, "/")
	newErr2.Position.File = filename_without_path[len(filename_without_path)-1]

	newErr3 := Wrap(InternalError, "fail to fetch statistics data", newErr2)
	filename_without_path = strings.Split(newErr3.Position.File, "/")
	newErr3.Position.File = filename_without_path[len(filename_without_path)-1]

	fmt.Println(newErr3.FormatWithTraces(true))
	// Output:
	// code id: 3; description: failed to perform application task
	// 	error message: fail to fetch statistics data; in file: error_test.go; at line: 121
	// 		trace message: fail to fetch statistics data; in file: error_test.go; at line: 117
	// 		trace message: fail to fetch statistics data; in file: error_test.go; at line: 113
	// 		trace message: message fail to compute the statistics; in file: error_test.go; at line: 109
}

func ExampleError_FormatWithTraces_GNU() {
	err := New(InternalError, "message fail to compute the statistics")
	filename_without_path := strings.Split(err.Position.File, "/")
	err.Position.File = filename_without_path[len(filename_without_path)-1]

	newErr := Wrap(InternalError, "fail to fetch statistics data", err)
	filename_without_path = strings.Split(newErr.Position.File, "/")
	newErr.Position.File = filename_without_path[len(filename_without_path)-1]

	newErr2 := Wrap(InternalError, "fail to fetch statistics data", newErr)
	filename_without_path = strings.Split(newErr2.Position.File, "/")
	newErr2.Position.File = filename_without_path[len(filename_without_path)-1]

	newErr3 := Wrap(InternalError, "fail to fetch statistics data", newErr2)
	filename_without_path = strings.Split(newErr3.Position.File, "/")
	newErr3.Position.File = filename_without_path[len(filename_without_path)-1]

	fmt.Println(newErr3.FormatWithTraces(false))
	// Output:
	// error_test.go:147: Error: 3:failed to perform application task:fail to fetch statistics data
	// error_test.go:143: Error: fail to fetch statistics data
	// error_test.go:139: Error: fail to fetch statistics data
	// error_test.go:135: Error: message fail to compute the statistics
}

func ExampleError_FormatJSON() {
	err := New(InternalError, "message fail to compute the statistics")
	filename_without_path := strings.Split(err.Position.File, "/")
	err.Position.File = filename_without_path[len(filename_without_path)-1]

	newErr := Wrap(InternalError, "fail to fetch statistics data", err)
	filename_without_path = strings.Split(newErr.Position.File, "/")
	newErr.Position.File = filename_without_path[len(filename_without_path)-1]

	newErr2 := Wrap(InternalError, "fail to fetch statistics data", newErr)
	filename_without_path = strings.Split(newErr2.Position.File, "/")
	newErr2.Position.File = filename_without_path[len(filename_without_path)-1]

	newErr3 := Wrap(InternalError, "fail to fetch statistics data", newErr2)
	filename_without_path = strings.Split(newErr3.Position.File, "/")
	newErr3.Position.File = filename_without_path[len(filename_without_path)-1]

	fmt.Println(newErr3.FormatJSON(false))
	// Output: {"code":{"id":3,"description":"failed to perform application task"},"message":"fail to fetch statistics data","position":{"file":"error_test.go","line":172},"traces":[{"message":"fail to fetch statistics data","position":{"file":"error_test.go","line":168}},{"message":"fail to fetch statistics data","position":{"file":"error_test.go","line":164}},{"message":"message fail to compute the statistics","position":{"file":"error_test.go","line":160}}]}
}

func ExampleTrace_Format() {
	trace := Trace{Message: "error database", Position: Position{File: "error_test.go", Line: 828}}
	fmt.Println(trace.Format(true))
	// Output: trace message: error database; in file: error_test.go; at line: 828
}

func ExampleTrace_Format_GNU() {
	trace := Trace{Message: "error database", Position: Position{File: "error_test.go", Line: 828}}
	fmt.Println(trace.Format(false))
	// Output: error_test.go:828: Error: error database
}

func TestNew(t *testing.T) {
	type args struct {
		code    Code
		message string
		traces  []Trace
	}

	tests := []struct {
		name string
		args args
		want *Error
	}{
		{
			name: "OK",
			args: args{
				code:    UnknownError,
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
			want: &Error{
				Code:    UnknownError,
				Message: "sample error",
				Position: Position{
					File: "error_test.go",
					Line: 284,
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
				code:    UnknownError,
				message: "sample error",
			},
			want: &Error{
				Code:    UnknownError,
				Message: "sample error",
				Position: Position{
					File: "error_test.go",
					Line: 284,
				},
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := New(testCase.args.code, testCase.args.message, testCase.args.traces...) // Error check based on the current line
			files := strings.Split(result.Position.File, "/")
			result.Position.File = files[len(files)-1]
			assert.Equal(t, testCase.want, result)
		})
	}
}

func TestWrap(t *testing.T) {
	type args struct {
		code    Code
		message string
		err     *Error
	}

	tests := []struct {
		name string
		args args
		want *Error
	}{
		{
			name: "OK",
			args: args{
				code:    UnknownError,
				message: "sample error",
				err: &Error{
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
			want: &Error{
				Code:    UnknownError,
				Message: "sample error",
				Position: Position{
					File: "error_test.go",
					Line: 383,
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
			result := Wrap(testCase.args.code, testCase.args.message, testCase.args.err) // Error check based on the current line
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
		args   Format
		fields Error
		want   string
	}{
		{
			name: "OK - GNU",
			args: GNU,
			fields: Error{
				Code:    UnknownError,
				Message: "sample error",
				Position: Position{
					File: "sample.go",
					Line: 50,
				},
				Traces: []Trace{
					{
						Message: "parent error",
						Position: Position{
							File: "sample2.go",
							Line: 76,
						},
					},
				},
			},
			want: "sample.go:50: Error: 0:failed to perform task:sample error",
		},
		{
			name: "OK - GNU With traces",
			args: GNUWithTraces,
			fields: Error{
				Code:    UnknownError,
				Message: "sample error",
				Position: Position{
					File: "sample.go",
					Line: 50,
				},
				Traces: []Trace{
					{
						Message: "parent error",
						Position: Position{
							File: "sample2.go",
							Line: 76,
						},
					},
				},
			},
			want: "sample.go:50: Error: 0:failed to perform task:sample error\nsample2.go:76: Error: parent error",
		},
		{
			name: "OK - Custom",
			args: Custom,
			fields: Error{
				Code:    UnknownError,
				Message: "sample error",
				Position: Position{
					File: "sample.go",
					Line: 50,
				},
				Traces: []Trace{
					{
						Message: "parent error",
						Position: Position{
							File: "sample2.go",
							Line: 76,
						},
					},
				},
			},
			want: "code id: 0; description: failed to perform task\n\terror message: sample error; in file: sample.go; at line: 50",
		},
		{
			name: "OK - Custom With traces",
			args: CustomWithTraces,
			fields: Error{
				Code:    UnknownError,
				Message: "sample error",
				Position: Position{
					File: "sample.go",
					Line: 50,
				},
				Traces: []Trace{
					{
						Message: "parent error",
						Position: Position{
							File: "sample2.go",
							Line: 76,
						},
					},
				},
			},
			want: "code id: 0; description: failed to perform task\n\terror message: sample error; in file: sample.go; at line: 50\n\t\ttrace message: parent error; in file: sample2.go; at line: 76",
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			GopherpanicFormat = testCase.args
			result := testCase.fields.Error()
			assert.Equal(t, testCase.want, result)
		})
		GopherpanicFormat = GNU
	}
}

func TestErrorFormat(t *testing.T) {
	type args struct {
		custom    bool
		withInner bool
	}
	tests := []struct {
		name   string
		fields Error
		args   args
		want   string
	}{
		{
			name: "OK - Custom",
			fields: Error{
				Code:    UnknownError,
				Message: "sample error",
				Position: Position{
					File: "sample.go",
					Line: 50,
				},
			},
			args: args{
				custom:    true,
				withInner: true,
			},
			want: "code id: 0; description: failed to perform task\n\terror message: sample error; in file: sample.go; at line: 50",
		},
		{
			name: "OK - Custom without inner data",
			fields: Error{
				Code:    UnknownError,
				Message: "sample error",
				Position: Position{
					File: "sample.go",
					Line: 50,
				},
			},
			args: args{
				custom:    true,
				withInner: false,
			},
			want: "code id: 0; description: failed to perform task\n\terror message: sample error",
		},
		{
			name: "OK - GNU Standard",
			fields: Error{
				Code:    UnknownError,
				Message: "sample error",
				Position: Position{
					File: "sample.go",
					Line: 50,
				},
			},
			args: args{
				custom:    false,
				withInner: true,
			},
			want: "sample.go:50: Error: 0:failed to perform task:sample error",
		},
		{
			name: "OK - GNU Standard inner data",
			fields: Error{
				Code:    UnknownError,
				Message: "sample error",
				Position: Position{
					File: "sample.go",
					Line: 50,
				},
			},
			args: args{
				custom:    false,
				withInner: false,
			},
			want: "Error: 0:failed to perform task:sample error",
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := testCase.fields.Format(testCase.args.custom, testCase.args.withInner)
			assert.Equal(t, testCase.want, result)
		})
	}
}

func TestErrorFormatWithTraces(t *testing.T) {
	tests := []struct {
		name   string
		fields Error
		args   bool
		want   string
	}{
		{
			name: "OK - Custom",
			fields: Error{
				Code:    UnknownError,
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
			args: true,
			want: "code id: 0; description: failed to perform task\n\terror message: sample error; in file: sample.go; at line: 50\n\t\ttrace message: inner 1; in file: inner_1.go; at line: 10\n\t\ttrace message: inner 2; in file: inner_2.go; at line: 20\n\t\ttrace message: inner 3; in file: inner_3.go; at line: 30",
		},
		{
			name: "OK - GNU Standard",
			fields: Error{
				Code:    UnknownError,
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
			args: false,
			want: "sample.go:50: Error: 0:failed to perform task:sample error\ninner_1.go:10: Error: inner 1\ninner_2.go:20: Error: inner 2\ninner_3.go:30: Error: inner 3",
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := testCase.fields.FormatWithTraces(testCase.args)
			assert.Equal(t, testCase.want, result)
		})
	}
}

func TestErrorFormatJSON(t *testing.T) {
	tests := []struct {
		name   string
		fields Error
		args   bool
		want   string
	}{
		{
			name: "OK",
			fields: Error{
				Code:    UnknownError,
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
			want: "{\"code\":{\"id\":0,\"description\":\"failed to perform task\"},\"message\":\"sample error\",\"position\":{\"file\":\"sample.go\",\"line\":50},\"traces\":[{\"message\":\"inner 1\",\"position\":{\"file\":\"inner_1.go\",\"line\":10}},{\"message\":\"inner 2\",\"position\":{\"file\":\"inner_2.go\",\"line\":20}},{\"message\":\"inner 3\",\"position\":{\"file\":\"inner_3.go\",\"line\":30}}]}",
		},
		{
			name: "OK - With Indent",
			args: true,
			fields: Error{
				Code:    UnknownError,
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
			want: "{\n\t\"code\": {\n\t\t\"id\": 0,\n\t\t\"description\": \"failed to perform task\"\n\t},\n\t\"message\": \"sample error\",\n\t\"position\": {\n\t\t\"file\": \"sample.go\",\n\t\t\"line\": 50\n\t},\n\t\"traces\": [\n\t\t{\n\t\t\t\"message\": \"inner 1\",\n\t\t\t\"position\": {\n\t\t\t\t\"file\": \"inner_1.go\",\n\t\t\t\t\"line\": 10\n\t\t\t}\n\t\t},\n\t\t{\n\t\t\t\"message\": \"inner 2\",\n\t\t\t\"position\": {\n\t\t\t\t\"file\": \"inner_2.go\",\n\t\t\t\t\"line\": 20\n\t\t\t}\n\t\t},\n\t\t{\n\t\t\t\"message\": \"inner 3\",\n\t\t\t\"position\": {\n\t\t\t\t\"file\": \"inner_3.go\",\n\t\t\t\t\"line\": 30\n\t\t\t}\n\t\t}\n\t]\n}",
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := testCase.fields.FormatJSON(testCase.args)
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
		args   bool
		want   string
	}{
		{
			name: "OK - Custom",
			fields: Trace{
				Message: "sample error",
				Position: Position{
					File: "sample.go",
					Line: 50,
				},
			},
			args: true,
			want: "trace message: sample error; in file: sample.go; at line: 50",
		},
		{
			name: "OK - GNU Standard",
			fields: Trace{
				Message: "sample error",
				Position: Position{
					File: "sample.go",
					Line: 50,
				},
			},
			args: false,
			want: "sample.go:50: Error: sample error",
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			result := testCase.fields.Format(testCase.args)
			assert.Equal(t, testCase.want, result)
		})
	}
}
