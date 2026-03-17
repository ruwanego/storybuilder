package mappers

// Error is the base mapper for error payloads.
type Error struct {
	Payload any `json:"errors"`
}
