package response

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	errHandler "github.com/storybuilder/storybuilder/transport/http/error"
	httpErrs "github.com/storybuilder/storybuilder/transport/http/errors"
	"github.com/storybuilder/storybuilder/transport/http/response/mappers"
)

// Send sets all required fields and write the response.
func Send(w http.ResponseWriter, payload any, code int) {
	// set headers
	w.Header().Set("Content-Type", "application/json")
	// set response code
	w.WriteHeader(code)
	// set payload
	var p []byte
	switch py := payload.(type) {
	case []byte:
		p = py
	case mappers.Payload:
		p = toJSON(py)
	}
	_, err := w.Write(p)
	if err != nil {
		fmt.Printf("JSON Writing Error: %v", err)
	}
}

// Error formats and sends the error response.
func Error(ctx context.Context, w http.ResponseWriter, err any, logger *slog.Logger) {
	var msg any
	code := http.StatusInternalServerError
	// check whether err is a general error or a validation error
	errG, isG := err.(error)
	errV, isV := err.(httpErrs.ValidationMapError)
	if isV {
		msg, code = errHandler.HandleValidationErrors(ctx, map[string]string(errV), logger)
	} else if isG {
		msg, code = errHandler.Handle(ctx, errG, logger)
	}
	Send(w, msg, code)
}

// toJSON converts the payload to JSON.
func toJSON(payload any) []byte {
	msg, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("JSON Marshaling Error: %v", err)
	}
	return msg
}
