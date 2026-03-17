package errors

import "fmt"

// DomainError is the type of errors thrown by business logic.
type DomainError struct {
	errType  string
	code     int
	httpCode int
	msg      string
	details  string
}

// NewDomainError creates a new DomainError.
func NewDomainError(message string, code, httpCode int, details string) error {
	return &DomainError{
		errType:  "ServiceError",
		code:     code,
		httpCode: httpCode,
		msg:      message,
		details:  details,
	}
}

// Error returns the DomainError message.
func (e *DomainError) Error() string {
	return fmt.Sprintf("%s: %s", e.errType, e.msg)
}

func (e *DomainError) Type() string { return e.errType }
func (e *DomainError) Code() int { return e.code }
func (e *DomainError) Msg() string { return e.msg }
func (e *DomainError) Trace() string { return e.details }
