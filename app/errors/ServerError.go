package errors

import "fmt"

// ServerError is the type of errors thrown by the framework while booting.
type ServerError struct {
	errType string
	code    int
	msg     string
	details string
}

// NewServerError creates a new ServerError instance.
func NewServerError(message string, code int, details string) error {
	return &ServerError{
		errType: "ServerError",
		code:    code,
		msg:     message,
		details: details,
	}
}

// Error returns the ServerError message.
func (e *ServerError) Error() string {
	return fmt.Sprintf("%s: %s", e.errType, e.msg)
}

func (e *ServerError) Type() string { return e.errType }
func (e *ServerError) Code() int { return e.code }
func (e *ServerError) Msg() string { return e.msg }
func (e *ServerError) Trace() string { return e.details }
