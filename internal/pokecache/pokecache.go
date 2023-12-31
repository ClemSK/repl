package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	pokecache map[string]cacheEntry
	mux       *sync.Mutex
}

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		pokecache: make(map[string]cacheEntry), // how we make a new cache
		mux:       &sync.Mutex{},
	}
	go c.reapLoop(interval) // happens in a separate go routine, otherwise will never execute
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock() // locks the protected resource to prevent concurrency issues
	defer c.mux.Unlock()
	c.pokecache[key] = cacheEntry{
		val:       val,
		createdAt: time.Now().UTC(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	cacheE, ok := c.pokecache[key]
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
	c.mux.Lock()
	defer c.mux.Unlock()
	cacheTimeout := time.Now().UTC().Add(-interval) // setting a negative interval gets time in the past
	for k, v := range c.pokecache {
		if v.createdAt.Before(cacheTimeout) {
			delete(c.pokecache, k)
		}
	}
}
