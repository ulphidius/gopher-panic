package gopherpanic

type ErrorKind uint

const (
	Unknown ErrorKind = iota
	IO
	Network
	Internal
	Client
	Authorization
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
	AuthorizationError Code = Code{
		ID:          Authorization,
		Description: "cannot perform unauthorize task",
	}
)

type Code struct {
	ID          ErrorKind `json:"id"`
	Description string    `json:"description,omitempty"`
}
