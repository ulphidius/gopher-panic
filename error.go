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

func New(code Code, message string, traces []Trace) *Error {
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
	return err.Format()
}

func (err Error) Format() string {
	return fmt.Sprintf(
		"error code: %d; error description: %s\n\terror: %s; in file: %s; at line: %d",
		err.Code.ID,
		err.Code.Description,
		err.Message,
		err.Position.File,
		err.Position.Line,
	)
}

func (err Error) FormatWithTraces() string {
	return iterago.Fold(err.Traces, err.Format(), func(acc string, trace Trace) string {
		return acc + fmt.Sprintf("\n\t\t%s", trace.Format())
	})
}

func (err Error) FormatJSON() string {
	data, _ := json.Marshal(err)
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

func (trace Trace) Format() string {
	return fmt.Sprintf("error: %s; in file: %s; at line: %d", trace.Message, trace.Position.File, trace.Position.Line)
}
