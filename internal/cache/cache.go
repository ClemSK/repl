package cache

import "time"

type Cache struct {
	cache map[string]cacheEntry
}

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry), // how we make a new cache
	}
	go c.reapLoop(interval) // happens in a separate go routine, otherwise will never execute
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.cache[key] = cacheEntry{
		val:       val,
		createdAt: time.Now().UTC(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	cacheE, ok := c.cache[key]
	return cacheE.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval) // chanel that receives a value every time the interval passes
	for range ticker.C {
		// the code in this block will run every interval
		c.reap(interval)
	}
}

func (c *Cache) reap(interval time.Duration) {
	cacheTimeout := time.Now().UTC().Add(-interval) // setting a negative interval gets time in the past
	for k, v := range c.cache {
		if v.createdAt.Before(cacheTimeout) {
			delete(c.cache, k)
		}
	}
}
