package errors

import "fmt"

// RepositoryError is the type of errors thrown by repositories.
type RepositoryError struct {
	errType string
	code    int
	msg     string
	details string
}

// NewRepositoryError creates a new RepositoryError instance.
func NewRepositoryError(message string, code int, details string) error {
	return &RepositoryError{
		errType: "RepositoryError",
		code:    code,
		msg:     message,
		details: details,
	}
}

// Error returns the RepositoryError message.
func (e *RepositoryError) Error() string {
	return fmt.Sprintf("%s: %s", e.errType, e.msg)
}

func (e *RepositoryError) Type() string { return e.errType }
func (e *RepositoryError) Code() int { return e.code }
func (e *RepositoryError) Msg() string { return e.msg }
func (e *RepositoryError) Trace() string { return e.details }
