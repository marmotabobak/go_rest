package cache

import (
	"hash/fnv"
	"sync"
)

/*
Нужно написать простую библиотеку in-memory cache. Для простоты считаем, что у нас бесконечная память и нам не нужно задумываться об удалении ключей из него.
Реализация должна удовлетворять интерфейсу:
type Cache interface {
    Set(k, v string)
    Get(k string) (v string, ok bool)
}
*/

// Эта структура и два ее метода ниже, в общем, решают поставленную задачу
// ПОСТАВЛЕННАЯ ЗАДАЧА
type Cache struct {
	m   map[string]string
	mtx sync.RWMutex
}

func (c *Cache) Set(k, v string) {
	c.mtx.Lock()
	c.m[k] = v
	c.mtx.Unlock()
}

func (c *Cache) Get(k string) (v string, ok bool) {
	c.mtx.RLock() // При больших количес твах операций чтения обычный мьютекс блокировал бы и вызывал задержки. В данном случае лучше подойдет мьютекс, позхволяющий читать параллельно.
	val, exists := c.m[k]
	c.mtx.RUnlock()
	return val, exists
}

// ПОСТАВЛЕННАЯ ЗАДАЧА РЕШЕНА

// Но: в реализаии выше проблема при большом количестве записей - будет блокироваться вся структура.
// Решение ниже дополняет "базовое" решение - добавялем шардирование - теперь блокироваться будет только конкретная шарда, куда пишем
type ShardedCache struct {
	shards []*Cache
}

func (sc *ShardedCache) Set(k, v string) {
	h := hash(k)
	i := int(h) % len(sc.shards)
	sc.shards[i].Set(k, v)
}

func NewShardedCache(shardsCount int) *ShardedCache {
	return &ShardedCache{shards: make([]*Cache, shardsCount)}
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}
