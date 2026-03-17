package errors

import "fmt"

// AdapterError is the type of errors thrown by adapters.
type AdapterError struct {
	errType string
	code    int
	msg     string
	details string
}

// NewAdapterError creates a new AdapterError instance.
func NewAdapterError(message string, code int, details string) error {
	return &AdapterError{
		errType: "AdapterError",
		code:    code,
		msg:     message,
		details: details,
	}
}

// Error returns the AdapterError message.
func (e *AdapterError) Error() string {
	return fmt.Sprintf("%s: %s", e.errType, e.msg)
}

func (e *AdapterError) Type() string { return e.errType }
func (e *AdapterError) Code() int { return e.code }
func (e *AdapterError) Msg() string { return e.msg }
func (e *AdapterError) Trace() string { return e.details }
