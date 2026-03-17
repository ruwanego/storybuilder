package transformers

// APITransformer is used to transform the response payload for API details.
//
// swagger:model APITransformer
type APITransformer struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Purpose string `json:"purpose"`
}
