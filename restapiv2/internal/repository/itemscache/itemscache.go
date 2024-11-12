package itemscache

import (
	"restapiv2/internal/mutex"
)

type CacheType map[string]string

var Cache CacheType = CacheType{}


func (c *CacheType) GetItem(key string) (string, bool) {
	mutex.M.RLock()
	item, found := (*c)[key]
	mutex.M.RUnlock()
	return item, found
}

func (c *CacheType) UpdateItem(key, value string) {
	mutex.M.Lock()
	(*c)[key] = value
	mutex.M.Unlock()
}

func (c *CacheType) DeleteItem(key string) {
	mutex.M.Lock()
	delete(*c, key)
	mutex.M.Unlock()
}