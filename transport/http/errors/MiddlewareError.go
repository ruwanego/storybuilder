package errors

import "fmt"

// MiddlewareError is the type of errors thrown by middleware.
type MiddlewareError struct {
	errType string
	code    int
	msg     string
	details string
}

// NewMiddlewareError creates a new MiddlewareError instance.
func NewMiddlewareError(message string, code int, details string) error {
	return &MiddlewareError{
		errType: "MiddlewareError",
		code:    code,
		msg:     message,
		details: details,
	}
}

// Error returns the MiddlewareError message.
func (e *MiddlewareError) Error() string {
	return fmt.Sprintf("%s: %s", e.errType, e.msg)
}

func (e *MiddlewareError) Type() string { return e.errType }
func (e *MiddlewareError) Code() int { return e.code }
func (e *MiddlewareError) Msg() string { return e.msg }
func (e *MiddlewareError) Trace() string { return e.details }
