package error

import (
	"context"
	"log/slog"
	"net/http"

	baseErrs "github.com/storybuilder/storybuilder/app/errors"
	domainErrs "github.com/storybuilder/storybuilder/domain/errors"
	externalErrs "github.com/storybuilder/storybuilder/externals/errors"
	httpErrs "github.com/storybuilder/storybuilder/transport/http/errors"
)

// Handle handles all errors globally.
func Handle(ctx context.Context, err error, logger *slog.Logger) (errMessage []byte, status int) {
	switch err.(type) {
	case *baseErrs.ServerError, *httpErrs.TransformerError:
		logger.ErrorContext(ctx, "Server Error", "error", err)
		status = http.StatusInternalServerError
	case *externalErrs.AdapterError, *httpErrs.MiddlewareError,
		*externalErrs.RepositoryError, *externalErrs.ServiceError,
		*domainErrs.DomainError:
		logger.ErrorContext(ctx, "Other Error", "error", err)
		status = http.StatusBadRequest
	case *httpErrs.ValidationError:
		logger.ErrorContext(ctx, "Unpacker Error", "error", err)
		status = http.StatusUnprocessableEntity
	default:
		logger.ErrorContext(ctx, "Unknown Error", "error", err)
		status = http.StatusInternalServerError
	}
	errMessage = format(err)
	return errMessage, status
}

// HandleValidationErrors specifically handles validation errors thrown by the validator.
func HandleValidationErrors(ctx context.Context, errs map[string]string, logger *slog.Logger) (errMessage []byte, status int) {
	errMessage = formatValidationErrors(errs)
	logger.ErrorContext(ctx, "Validation Errors", "errors", string(errMessage))
	return errMessage, http.StatusUnprocessableEntity
}
