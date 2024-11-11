package itemscache

import (
	"restapiv2/internal/mutex"
)

type CacheType map[string]string

var Cache CacheType

func Init() {
	Cache = make(CacheType)
}

func (c CacheType) GetItem(k string) (string, bool) {
	mutex.M.RLock()
	item, found := c[k]
	mutex.M.RUnlock()
	return item, found
}

func (c CacheType) UpdateItem(k, v string) {
	mutex.M.Lock()
	c[k] = v
	mutex.M.Unlock()
}

func (c CacheType) DeleteItem(k string) {
	mutex.M.Lock()
	delete(c, k)
	mutex.M.Unlock()
}