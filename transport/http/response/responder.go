package response

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/storybuilder/storybuilder/domain/boundary/adapters"
	errHandler "github.com/storybuilder/storybuilder/transport/http/error"
	"github.com/storybuilder/storybuilder/transport/http/response/mappers"
)

// Send sets all required fields and write the response.
func Send(w http.ResponseWriter, payload interface{}, code int) {
	// set headers
	w.Header().Set("Content-Type", "application/json")
	// set response code
	w.WriteHeader(code)
	// set payload
	var p []byte
	switch payload.(type) {
	case []byte:
		p = payload.([]byte)
	case mappers.Payload:
		p = toJSON(payload)
	}
	_, err := w.Write(p)
	if err != nil {
		fmt.Printf("JSON Writing Error: %v", err)
	}
}

// Error formats and sends the error response.
func Error(ctx context.Context, w http.ResponseWriter, err interface{}, logger adapters.LogAdapterInterface) {
	var msg interface{}
	code := http.StatusInternalServerError
	// check whether err is a general error or a validation error
	errG, isG := err.(error)
	errV, isV := err.(map[string]string)
	if isG {
		msg, code = errHandler.Handle(ctx, errG, logger)
	}
	if isV {
		msg, code = errHandler.HandleValidationErrors(ctx, errV, logger)
	}
	Send(w, msg, code)
}

// toJSON converts the payload to JSON.
func toJSON(payload interface{}) []byte {
	msg, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("JSON Marshaling Error: %v", err)
	}
	return msg
}
