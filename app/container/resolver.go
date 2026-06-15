package container

import "github.com/storybuilder/storybuilder/app/config"

// Resolve resolves the entire container.
//
// The order of resolution is very important. Low level dependencies need to be resolved before high level dependencies.
// It generally happens in this order.
//   - Adapters
//   - Repositories
//   - Services
func Resolve(cfg *config.Config) *Container {
	return &Container{
		Adapters:     resolveAdapters(cfg),
		Repositories: resolveRepositories(),
		Services:     resolveServices(cfg.ServicesConfig),
	}
}
