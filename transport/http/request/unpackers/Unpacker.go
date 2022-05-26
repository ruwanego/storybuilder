package unpackers

// Unpacker is the interface implemented by all unpacker data structures.
type Unpacker interface {
	// RequiredFormat string representation of the required format of the relevant request body.
	RequiredFormat() string
}
