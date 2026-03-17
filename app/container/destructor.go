package container

import (
	"context"
)

// Destruct releases all necessary resources that needs to be released.
func (ctr *Container) Destruct() {
	ctx := context.Background()
	ctr.Adapters.LogAdapter.Info(ctx, "Closing database connections...")
	ctr.Adapters.DBAdapter.Destruct()
	ctr.Adapters.LogAdapter.Info(ctx, "Clearing cache...")
	ctr.Adapters.CacheAdapter.Destruct()
	ctr.Adapters.LogAdapter.Info(ctx, "Closing logger...")
	ctr.Adapters.LogAdapter.Destruct()
}
