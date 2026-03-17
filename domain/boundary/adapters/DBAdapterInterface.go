package adapters

import (
	"context"
	"database/sql"
)

// DBAdapterInterface is implemented by all database adapters.
type DBAdapterInterface interface {
	// Ping checks whether the database is accessible.
	Ping() error
	// Query runs a query and return the result.
	Query(ctx context.Context, query string, parameters map[string]any) ([]map[string]any, error)
	// NewTransaction creates a new database transaction.
	NewTransaction() (*sql.Tx, error)
	// Destruct will close the database adapter releasing all resources.
	Destruct()
}
