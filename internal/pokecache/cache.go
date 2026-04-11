package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu           sync.Mutex
	reapInterval time.Duration
	vals         map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		reapInterval: interval,
		vals:         make(map[string]cacheEntry),
	}

	go c.reapLoop()

	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.vals[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, ok := c.vals[key]

	return val.val, ok
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.reapInterval)

	for range ticker.C {
		c.mu.Lock()

		for key, value := range c.vals {
			if time.Since(value.createdAt) >= c.reapInterval {
				delete(c.vals, key)
			}
		}

		c.mu.Unlock()
	}
}
