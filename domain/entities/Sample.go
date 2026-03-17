package entities

// Sample represents a sample domain entity.
//
// swagger:model Sample
type Sample struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
