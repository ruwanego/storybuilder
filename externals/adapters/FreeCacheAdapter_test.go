package adapters_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/storybuilder/storybuilder/app/config"
	"github.com/storybuilder/storybuilder/externals/adapters"
)

func TestNewFreeCacheAdapter(t *testing.T) {
	cfg := config.CacheConfig{
		HardMaxSize: 10,
		LifeWindow:  config.DurString("10m"),
	}
	adapter, err := adapters.NewFreeCacheAdapter(cfg)
	assert.NoError(t, err)
	assert.NotNil(t, adapter)
}

func TestFreeCacheAdapter_Set(t *testing.T) {
	cfg := config.CacheConfig{
		HardMaxSize: 10,
		LifeWindow:  config.DurString("10m"),
	}
	adapter, err := adapters.NewFreeCacheAdapter(cfg)
	assert.NoError(t, err)

	adapter.Set("key1", []byte("value1"))
	val := adapter.Get("key1")
	assert.Equal(t, []byte("value1"), val)

	// test empty key
	adapter.Set("", []byte("empty"))
	valEmpty := adapter.Get("")
	assert.Nil(t, valEmpty)
}
