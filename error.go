package gopherpanic

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/ulphidius/iterago"
)

type Format uint

const (
	GNU Format = iota
	GNUWithTraces
	Custom
	CustomWithTraces
)

var GopherpanicFormat Format = GNU

func init() {
	format, err := strconv.Atoi(os.Getenv("GOPHERPANIC_FORMAT"))
	if err != nil || format > 3 || format < 0 {
		return
	}

	GopherpanicFormat = Format(format)
}

// Representation of an error
type Error struct {
	Code     Code     `json:"code"`             // Kind of error (Internal, client, etc.)
	Message  string   `json:"message"`          // Message which describe the user error
	Position Position `json:"position"`         // Where the Error is spawns in the user code (Auto generation if New or Wrap is used)
	Traces   []Trace  `json:"traces,omitempty"` // Wrapped parent errors
}

// Create a new error with the user parameters and current spawn position
func New(code Code, message string, traces ...Trace) *Error {
	return &Error{
		Code:     code,
		Message:  message,
		Position: Position{}.spawn(2),
		Traces:   traces,
	}
}

// Create a new error that wraps an existing error.
//
// The build behavior is equivalent to New function.
func Wrap(code Code, message string, err *Error) *Error {
	newErr := ErrorBuilder{}.New().
		WithCode(code).
		WithMessage(message).
		WithPosition(Position{}.spawn(2)).
		WithTraces(append([]Trace{err.IntoTrace()}, err.Traces...)...).
		Build()
	return &newErr
}

func (err Error) IntoTrace() Trace {
	return Trace{
		Message:  err.Message,
		Position: err.Position,
	}
}

// The output changes depending of the GOPHERPANIC_FORMAT value
//
// - GNU(0): sample.go:50: Error: 0:failed to perform task:sample error
//
// - GNUWithTraces(1): sample.go:50: Error: 0:failed to perform task:sample error
//
// - Custom(2): code id: 0; description: failed to perform task\n\terror message: sample error
//
// - CustomWithTraces(3): code id: 0; description: failed to perform task\n\terror message: sample error
func (err Error) Error() string {
	switch GopherpanicFormat {
	case Custom:
		return err.Format(true, true)
	case CustomWithTraces:
		return err.FormatWithTraces(true)
	case GNU:
		return err.Format(false, true)
	case GNUWithTraces:
		return err.FormatWithTraces(false)
	default:
		return err.Format(false, true)
	}
}

// Convert into string the Error structure without Traces.
// Can remove the position data.
//
// Allowed formats:
//
// - gopherpanic format
//
// - GNU format
func (err Error) Format(custom bool, withInnerData bool) string {
	if custom {
		if !withInnerData {
			return fmt.Sprintf(
				"code id: %d; description: %s\n\terror message: %s",
				err.Code.ID,
				err.Code.Description,
				err.Message,
			)
		}

		return fmt.Sprintf(
			"code id: %d; description: %s\n\terror message: %s; in file: %s; at line: %d",
			err.Code.ID,
			err.Code.Description,
			err.Message,
			err.Position.File,
			err.Position.Line,
		)
	}

	if !withInnerData {
		return fmt.Sprintf(
			"Error: %d:%s:%s",
			err.Code.ID,
			err.Code.Description,
			err.Message,
		)
	}

	return fmt.Sprintf(
		"%s:%d: Error: %d:%s:%s",
		err.Position.File,
		err.Position.Line,
		err.Code.ID,
		err.Code.Description,
		err.Message,
	)
}

// Convert into string the Error structure with Traces.
//
// Allowed formats:
//
// - gopherpanic format
//
// - GNU format
func (err Error) FormatWithTraces(custom bool) string {
	if custom {
		return iterago.Fold(err.Traces, err.Format(custom, true), func(acc string, trace Trace) string {
			return acc + fmt.Sprintf("\n\t\t%s", trace.Format(custom))
		})
	}

	return iterago.Fold(err.Traces, err.Format(custom, true), func(acc string, trace Trace) string {
		return acc + fmt.Sprintf("\n%s", trace.Format(custom))
	})
}

// Convert into JSON string (with or without indentation)
func (err Error) FormatJSON(indent bool) string {
	var data []byte

	if indent {
		data, _ = json.MarshalIndent(err, "", "\t")
		return string(data)
	}

	data, _ = json.Marshal(err)
	return string(data)
}

// Representation of a parent error
type Trace struct {
	Message  string   `json:"message"`  // Message which describe the user error. Retrived from the Error structucture
	Position Position `json:"position"` // Where the Error is spawns in the user code (Auto generation if New or Wrap is used). Retrived from the Error structure
}

func (trace Trace) IntoError() Error {
	return Error{
		Message:  trace.Message,
		Position: trace.Position,
	}
}

// Convert into string
//
// Allowed formats:
//
// - gopherpanic format
//
// - GNU format
func (trace Trace) Format(custom bool) string {
	if custom {
		return fmt.Sprintf("trace message: %s; in file: %s; at line: %d", trace.Message, trace.Position.File, trace.Position.Line)
	}

	return fmt.Sprintf("%s:%d: Error: %s", trace.Position.File, trace.Position.Line, trace.Message)
}
