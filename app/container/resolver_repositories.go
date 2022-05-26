package container

import (
	"github.com/storybuilder/storybuilder/externals/repositories"
)

var resolvedRepositories Repositories

// resolveRepositories resolve all repositories.
func resolveRepositories() Repositories {
	resolvedRepositories.SampleRepository = repositories.NewSampleMySQLRepository(resolvedAdapters.DBAdapter)
	return resolvedRepositories
}
