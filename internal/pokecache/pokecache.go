package pokecache

import (
	"sync"
	"time"
)

// cacheEntry Holds information about the time in wich the entry was created and the value it holds as a []byte
type cacheEntry struct {
	createdAt time.Time
	value     []byte
}

func newCacheEntry(timeStamp time.Time, value []byte) cacheEntry {
	return cacheEntry{
		createdAt: timeStamp,
		value:     value,
	}
}

// Cache Is a map de holds mutiple cacheEntry. It is thread safe
type Cache struct {
	mux      sync.RWMutex
	info     map[string]cacheEntry
	interval time.Duration
}

// NewCache Creates a new cache struct
func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		info:     make(map[string]cacheEntry),
		interval: interval,
	}
	go cache.reapLoop()
	return &cache
}

// Add adds the value into the cache
func (c *Cache) Add(key string, value []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	timeStamp := time.Now()
	c.info[key] = newCacheEntry(timeStamp, value)
	return
}

// Get Returns the information at the cache, if any. The boolean will be true if the key has a value in the cache, false otherwise.
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.RLock()
	defer c.mux.RUnlock()
	entry, ok := c.info[key]
	if !ok {
		return nil, false
	}
	return entry.value, true
}

// reapLoop This loop clears the cache in the interval set in at the cache creation
func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for {
		<-ticker.C
		now := time.Now()
		for key, value := range c.info {
			if value.createdAt.Sub(now) < c.interval {
				c.mux.Lock()
				delete(c.info, key)
				c.mux.Unlock()
			}
		}
	}
}
