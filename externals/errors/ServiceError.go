package errors

import "fmt"

// ServiceError is the type of errors thrown by services talking to third party APIs.
type ServiceError struct {
	errType string
	code    int
	msg     string
	details string
}

// NewServiceError creates a new ServiceError instance.
func NewServiceError(message string, code int, details string) error {
	return &ServiceError{
		errType: "ServiceError",
		code:    code,
		msg:     message,
		details: details,
	}
}

// Error returns the ServiceError message.
func (e *ServiceError) Error() string {
	return fmt.Sprintf("%s: %s", e.errType, e.msg)
}

func (e *ServiceError) Type() string { return e.errType }
func (e *ServiceError) Code() int { return e.code }
func (e *ServiceError) Msg() string { return e.msg }
func (e *ServiceError) Trace() string { return e.details }
