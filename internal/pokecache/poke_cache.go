package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entry map[string]cacheEntry
	mux   *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		entry: make(map[string]cacheEntry),
		mux:   &sync.Mutex{},
	}

	go cache.reapLoop(interval)

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.entry[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	val, ok := c.entry[key]
	return val.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	tick := time.NewTicker(interval)
	for range tick.C {
		c.reap(time.Now(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	for k, v := range c.entry {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.entry, k)
		}
	}
}
