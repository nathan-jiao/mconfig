package pkg

import (
	"sync"
)

// MconfigCache ...
type MconfigCache struct {
	cache map[Appkey]*AppConfigsMap
	sync.RWMutex
}

var mconfigCache *MconfigCache

func init() {
	mconfigCache = &MconfigCache{
		cache: make(map[Appkey]*AppConfigsMap),
	}
}

func (cache *MconfigCache) getConfigCache(key Appkey) (*AppConfigsMap, error) {
	cache.RLock()
	value, ok := cache.cache[key]
	cache.RUnlock()
	if ok {
		return value, nil
	}
	return nil, Error_AppConfigNotFound
}

func (cache *MconfigCache) putConfigCache(key Appkey, configs *AppConfigs) error {
	configsMap := &AppConfigsMap{
		AppConfigs: configs,
	}
	cache.Lock()
	defer cache.Unlock()
	cache.cache[key] = configsMap
	return nil
}