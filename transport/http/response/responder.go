package response

import (
	"context"
	"fmt"
	"net/http"

	"github.com/storybuilder/storybuilder/domain/boundary/adapters"
	errHandler "github.com/storybuilder/storybuilder/transport/http/error"
	"github.com/storybuilder/storybuilder/transport/http/response/mappers"
)

// Send sets all required fields and write the response.
func Send[byteOrPayload mappers.ByteOrPayload](w http.ResponseWriter, payload byteOrPayload, code int) {
	// set headers
	w.Header().Set("Content-Type", "application/json")
	// set response code
	w.WriteHeader(code)
	_, err := w.Write(payload.ToJSON())
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
	Send(w, msg.(mappers.MyByte), code)
}
