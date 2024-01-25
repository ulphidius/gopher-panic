package gopherpanic

// Error type
//
// The use can create new Error Code by extending the existing Code
type ErrorKind uint

const (
	Unknown ErrorKind = iota
	IO
	Network
	Internal
	Client
	Unauthorized
	Timeout
	Unimplemented
)

var (
	UnknownError Code = Code{
		ID:          Unknown,
		Description: "failed to perform task",
	}
	IOError Code = Code{
		ID:          IO,
		Description: "failed to perform IO task",
	}
	NetworkError Code = Code{
		ID:          Network,
		Description: "failed to perform network task",
	}
	InternalError Code = Code{
		ID:          Internal,
		Description: "failed to perform application task",
	}
	ClientError Code = Code{
		ID:          Client,
		Description: "failed to perform client api task",
	}
	UnauthorizedError Code = Code{
		ID:          Unauthorized,
		Description: "cannot perform unauthorized task",
	}
	TimeoutError Code = Code{
		ID:          Timeout,
		Description: "failed to perform the task, the deadline is exceeded",
	}
	UnimplementedError Code = Code{
		ID:          Unimplemented,
		Description: "unimplemented behavior",
	}
)

// Type of error
type Code struct {
	ID          ErrorKind `json:"id"`
	Description string    `json:"description,omitempty"`
}
