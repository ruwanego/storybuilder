package errors

// ValidationMapError is the type of errors thrown by the field wise validator.
type ValidationMapError map[string]string

// Error returns a generic validation failed message.
func (e ValidationMapError) Error() string {
	return "validation failed"
}
