package pokecache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	val       []byte
	createdAt time.Time
}
type Cache struct {
	mutex    sync.Mutex
	hmap     map[string]CacheEntry
	interval time.Duration
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		hmap:     make(map[string]CacheEntry),
		interval: interval,
	}
	go cache.ReapLoop()
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	entry := CacheEntry{val: val,
		createdAt: time.Now()}
	c.hmap[key] = entry
}
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	entry, found := c.hmap[key]
	if !found {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) ReapLoop() {
	for {
		time.Sleep(c.interval)
		c.mutex.Lock()
		now := time.Now()
		for key, value := range c.hmap {
			if now.Sub(value.createdAt) > c.interval {
				delete(c.hmap, key)
			}
		}
		c.mutex.Unlock()
	}

}
