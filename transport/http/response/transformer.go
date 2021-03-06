package response

import (
	"github.com/storybuilder/storybuilder/transport/http/response/mappers"
	"github.com/storybuilder/storybuilder/transport/http/response/transformers"
)

// Transform transforms data either as an object or as a collection depending on the `isCollection` boolean value.
func Transform(data interface{}, t transformers.TransformerInterface, isCollection bool) (interface{}, error) {
	if isCollection {
		return t.TransformAsCollection(data)
	}
	return t.TransformAsObject(data)
}

// Map wraps payload in a standard response payload object.
func Map(data []interface{}) (m mappers.Payload) {
	for _, v := range data {
		switch v.(type) {
		default:
			m.Data = v
		}
	}
	return m
}
