package gopherpanic

import (
	"encoding/json"
	"fmt"

	"github.com/ulphidius/iterago"
)

type Error struct {
	Code     Code     `json:"code"`
	Message  string   `json:"message"`
	Position Position `json:"position"`
	Traces   []Trace  `json:"traces,omitempty"`
}

func New(code Code, message string, traces ...Trace) *Error {
	return &Error{
		Code:     code,
		Message:  message,
		Position: Position{}.spawn(2),
		Traces:   traces,
	}
}

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

func (err Error) Error() string {
	return err.Format(false, false)
}

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

func (err Error) FormatJSON(indent bool) string {
	var data []byte

	if indent {
		data, _ = json.MarshalIndent(err, "", "\t")
		return string(data)
	}

	data, _ = json.Marshal(err)
	return string(data)
}

type Trace struct {
	Message  string   `json:"message"`
	Position Position `json:"position"`
}

func (trace Trace) IntoError() Error {
	return Error{
		Message:  trace.Message,
		Position: trace.Position,
	}
}

func (trace Trace) Format(custom bool) string {
	if custom {
		return fmt.Sprintf("trace message: %s; in file: %s; at line: %d", trace.Message, trace.Position.File, trace.Position.Line)
	}

	return fmt.Sprintf("%s:%d: Error: %s", trace.Position.File, trace.Position.Line, trace.Message)
}
