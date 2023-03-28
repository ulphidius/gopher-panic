package gopherpanic

type Builder interface {
	New() ErrorBuilder
	Default() ErrorBuilder
	WithMessage(message string) ErrorBuilder
	WithPosition(position Position) ErrorBuilder
	WithCause(causes ...Error) ErrorBuilder
	Build() Error
}

type ErrorBuilder struct {
	message  string
	position Position
	traces   []Trace
}

func (builder ErrorBuilder) New() ErrorBuilder {
	return ErrorBuilder{}
}

func (builder ErrorBuilder) Default() ErrorBuilder {
	return ErrorBuilder{message: "an unexpected error occured", position: Position{}.spawn(2), traces: nil}
}

func (builder ErrorBuilder) WithMessage(message string) ErrorBuilder {
	builder.message = message
	return builder
}

func (builder ErrorBuilder) WithPosition(position Position) ErrorBuilder {
	builder.position = position
	return builder
}

func (builder ErrorBuilder) WithTraces(traces ...Trace) ErrorBuilder {
	builder.traces = traces
	return builder
}

func (builder ErrorBuilder) Build() Error {
	return Error{
		Message:  builder.message,
		Position: builder.position,
		Traces:   builder.traces,
	}
}
