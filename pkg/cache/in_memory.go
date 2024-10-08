package cache

import (
	"sync"

	"github.com/maypok86/otter"
)

var (
	cacheInstance *otter.CacheWithVariableTTL[string, string]
	once          sync.Once
	cacheErr      error
)

// GetCache initializes and returns the singleton cache instance.
// Note: Don't use this in production. Use a distributed cache like Redis or Memcached,
// which can be scaled and managed more easily in a Kubernetes environment.
func GetCache() (*otter.CacheWithVariableTTL[string, string], error) {
	once.Do(func() {
		cache, err := otter.MustBuilder[string, string](10_000).
			CollectStats().
			Cost(func(key string, value string) uint32 {
				return 1
			}).
			WithVariableTTL().
			Build()

		if err != nil {
			cacheErr = err
			return
		}

		cacheInstance = &cache
	})

	return cacheInstance, cacheErr
}

func GetDataCache() {
	//TODO impl GetDataCache
}

func SetDataCache() {
	//TODO impl SetDataCache
}
