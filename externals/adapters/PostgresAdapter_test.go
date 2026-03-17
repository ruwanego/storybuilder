package adapters_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/storybuilder/storybuilder/app/config"
	"github.com/storybuilder/storybuilder/externals/adapters"
)

func TestPostgresAdapter_Destruct(t *testing.T) {
	cfg := config.DBConfig{
		Host:     "localhost",
		Port:     5432,
		Database: "test",
		User:     "test",
		Password: "test",
		PoolSize: 10,
		Check:    false,
	}
	adapter, err := adapters.NewPostgresAdapter(cfg)
	assert.NoError(t, err)
	assert.NotNil(t, adapter)

	// Call Destruct and ensure no panics occur
	assert.NotPanics(t, func() {
		adapter.Destruct()
	})
}
