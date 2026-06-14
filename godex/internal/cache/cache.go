package cache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries  map[string]cacheEntry
	mu       sync.Mutex
	interval time.Duration
}

func NewCache(d time.Duration) *Cache {
	c := &Cache{
		interval: d,
	}
	go c.reapLoop()
	return c
}

func (c *Cache) Set(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.entries == nil {
		c.entries = make(map[string]cacheEntry)
	}

	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if entry, found := c.entries[key]; found {
		return entry.val, found
	} else {
		return nil, found
	}
}

func (c *Cache) reapLoop() {
	t := time.NewTicker(c.interval)
	defer t.Stop()

	for range t.C {
		c.mu.Lock()
		for key, entry := range c.entries {
			if time.Since(entry.createdAt) > c.interval {
				delete(c.entries, key)
			}
		}
		c.mu.Unlock()
	}
}
