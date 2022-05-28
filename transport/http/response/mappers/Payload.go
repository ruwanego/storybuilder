package mappers

import (
	"encoding/json"
	"fmt"
)

type ByteOrPayload interface {
	MyByte | Payload
	ToJSON() []byte
}

type MyByte []byte

func (b MyByte) ToJSON() []byte {
	return b
}

// Payload is the base mapper for data payloads.
type Payload struct {
	Data      interface{} `json:"data"`
	Paginator interface{} `json:"paginator,omitempty"`
}

func (p Payload) ToJSON() []byte {
	msg, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("JSON Marshaling Error: %v", err)
	}
	return msg
}
