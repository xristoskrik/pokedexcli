package pokecache

import (
	"sync"
	"time"
)

type PokeCache struct {
	map_chache map[string]cacheEntry
	mu         *sync.Mutex
}

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

func NewCache(interval time.Duration) PokeCache {
	c := PokeCache{
		map_chache: make(map[string]cacheEntry),
		mu:         &sync.Mutex{},
	}
	go c.reapLoop(interval)
	return c

}

func (c *PokeCache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.map_chache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}
func (c *PokeCache) Get(key string) ([]byte, bool) {
	if _, ok := c.map_chache[key]; !ok {
		return nil, false
	}
	return c.map_chache[key].val, true
}
func (c *PokeCache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval)
	}
}
func (c *PokeCache) reap(interval time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	timeago := time.Now().Add(-interval)
	for k, v := range c.map_chache {
		if v.createdAt.Before(timeago) {
			delete(c.map_chache, k)
		}
	}
}
