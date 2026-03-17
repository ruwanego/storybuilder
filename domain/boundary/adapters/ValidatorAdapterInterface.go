package adapters

// ValidatorAdapterInterface is implemented by all validator adapters.
type ValidatorAdapterInterface interface {
	// Validate validates bound values of an unpacker struct against
	// validation rules defined in that unpacker struct.
	Validate(data any) map[string]string
	// ValidateField validates a single variable.
	ValidateField(field any, rules string) map[string]string
}
