package itemscache

import (
	"sync"
)

type Cache struct {
	cache map[string]string
	m     sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		cache: make(map[string]string),
		m:     sync.RWMutex{},
	}
}

func (c *Cache) GetItem(key string) (string, bool) {
	c.m.RLock()
	item, found := c.cache[key]
	c.m.RUnlock()
	return item, found
}

func (c *Cache) UpdateItem(key, value string) {
	c.m.Lock()
	c.cache[key] = value
	c.m.Unlock()
}

func (c *Cache) DeleteItem(key string) {
	c.m.Lock()
	delete(c.cache, key)
	c.m.Unlock()
}

func (c *Cache) ReturnValueIfExists(key string) (string, bool) {
	val, exists :=  c.cache[key]
	if !exists {
		return "", false
	}
	return val, true
}