package container

import (
	"log/slog"

	"github.com/storybuilder/storybuilder/domain/boundary/adapters"
	"github.com/storybuilder/storybuilder/domain/boundary/repositories"
	"github.com/storybuilder/storybuilder/domain/boundary/services"
)

// Container holds all resolved dependencies that needs to be injected at run time.
type Container struct {
	Adapters     Adapters
	Repositories Repositories
	Services     Services
}

// Adapters hold resolved adapter instances.
//
// These are wrappers around third party libraries. All adapters will be of a corresponding adapter interface type.
type Adapters struct {
	DBAdapter        adapters.DBAdapterInterface
	DBTxAdapter      adapters.DBTxAdapterInterface
	LogAdapter       *slog.Logger
	ValidatorAdapter adapters.ValidatorAdapterInterface
	CacheAdapter     adapters.CacheAdapterInterface
}

// Repositories hold resolved repository instances.
//
// These are used to connect to databases. All repositories will be of a corresponding repository interface type.
// They act as an abstraction layer between the application and the database.
// Generally a single repository represents a single table in the database along with all the actions
// that can be performed on that table.
type Repositories struct {
	SampleRepository repositories.SampleRepositoryInterface
}

// Services hold resolved service instances.
//
// These are abstractions to third party APIs. All services will be of a corresponding service interface type.
type Services struct {
	SampleService services.SampleServiceInterface
}
