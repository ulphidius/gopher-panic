package gopherpanic

import "runtime"

// Reprentation of spawn position in the code
type Position struct {
	File string `json:"file"` // Filepath where the error is spawned
	Line int    `json:"line"` // Line where the error is spawned
}

// Create a new position with the data of where the method is called
func (position Position) Spawn() Position {
	_, file, line, _ := runtime.Caller(1)
	return Position{
		File: file,
		Line: line,
	}
}

// Used to fetch position data of where the parent function which calls this method is called itself
func (position Position) spawn(parentLevel int) Position {
	_, file, line, _ := runtime.Caller(parentLevel)
	return Position{
		File: file,
		Line: line,
	}
}
