package itemscache

import (
	"sync"
	"strconv"
	"fmt"
)

type CacheError struct {
	Code int
	Desc string
}

const (
	IncreaseErrorValueNotIntCode = 1001
	AbsentKeyErrorCode = 1002
)

func NewCacheError(errCode int, errDesc string) *CacheError {
	return &CacheError{
		Code: errCode,
		Desc: errDesc,
	}
}

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

func (c *Cache) IncreaseValue(key string, increment int) *CacheError {
	
	c.m.Lock()
	defer c.m.Unlock()

	val, exists := c.cache[key]
	if !exists {
		return NewCacheError(AbsentKeyErrorCode, "Key not found\n")
	}
	
	currentValInt, err := strconv.Atoi(val)
	if err != nil {
		return NewCacheError(IncreaseErrorValueNotIntCode, "Value should be int\n")
	}

	c.cache[key] = fmt.Sprintf("%d", currentValInt+increment)
	return nil
}
