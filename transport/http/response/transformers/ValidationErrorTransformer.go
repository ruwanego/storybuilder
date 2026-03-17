package transformers

// ValidationErrorTransformer is used to transform the response payload for validation errors.
//
// swagger:model ValidationErrorTransformer
type ValidationErrorTransformer struct {
	Type  string `json:"type,omitempty"`
	Trace any    `json:"trace,omitempty"`
}
