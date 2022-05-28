package adapters

import (
	"runtime/debug"
	"time"

	"github.com/coocood/freecache"
	_ "github.com/go-sql-driver/mysql"

	"github.com/storybuilder/storybuilder/app/config"
	"github.com/storybuilder/storybuilder/domain/boundary/adapters"
)

// FreeCacheAdapter is used to communicate with a free cache.
type FreeCacheAdapter struct {
	cache      *freecache.Cache
	expireTime time.Duration
}

// NewFreeCacheAdapter creates a new Cache adapter instance.
func NewFreeCacheAdapter(cfg config.CacheConfig) (adapters.CacheAdapterInterface, error) {
	cacheSize := cfg.HardMaxSize * 1024 * 1024
	cache := freecache.NewCache(cacheSize)
	debug.SetGCPercent(20)
	a := &FreeCacheAdapter{
		cache:      cache,
		expireTime: cfg.LifeWindow.Dur(),
	}
	return a, nil
}

func (f FreeCacheAdapter) Del(key string) {
	if len(key) == 0 {
		return
	}
	f.cache.Del([]byte(key))
}

func (f FreeCacheAdapter) Get(key string) []byte {
	if len(key) == 0 {
		return nil
	}
	value, err := f.cache.Get([]byte(key))
	if err != nil {
		return nil
	}
	return value
}

func (f FreeCacheAdapter) Set(key string, value []byte) {
	if len(key) == 0 {
		return
	}
	f.cache.Set([]byte(key), value, int(f.expireTime.Seconds()))
}

func (f FreeCacheAdapter) Destruct() {
	f.cache.Clear()
}
