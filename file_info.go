package gopherpanic

import "runtime"

type Position struct {
	File string `json:"file"`
	Line int    `json:"line"`
}

func (position Position) Spawn() Position {
	_, file, line, _ := runtime.Caller(1)
	return Position{
		File: file,
		Line: line,
	}
}

func (position Position) spawn(parentLevel int) Position {
	_, file, line, _ := runtime.Caller(parentLevel)
	return Position{
		File: file,
		Line: line,
	}
}
