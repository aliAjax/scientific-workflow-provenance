package errs

import "fmt"

type Code string

const (
	BadRequest  Code = "bad_request"
	NotFound    Code = "not_found"
	Conflict    Code = "conflict"
	Unavailable Code = "unavailable"
	Internal    Code = "internal"
)

type Error struct {
	Code    Code
	Message string
	Cause   error
}

func (e Error) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s: %s: %v", e.Code, e.Message, e.Cause)
	}
	return string(e.Code) + ": " + e.Message
}
func (e Error) Unwrap() error    { return e.Cause }
func New(c Code, m string) Error { return Error{Code: c, Message: m} }
