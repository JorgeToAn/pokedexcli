package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	interval time.Duration
	entries  map[string]cacheEntry
	mux      *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	newCache := Cache{
		interval: interval,
		entries:  make(map[string]cacheEntry),
		mux:      &sync.Mutex{},
	}
	go newCache.reapLoop()
	return newCache
}

func (c Cache) Add(key string, val []byte) {
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c Cache) Get(key string) ([]byte, bool) {
	entry, ok := c.entries[key]
	if !ok {
		return nil, false
	}

	return entry.val, true
}

func (c Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for {
		t := <-ticker.C
		c.mux.Lock()
		defer c.mux.Unlock()
		for key, entry := range c.entries {
			diff := t.Sub(entry.createdAt)
			if diff > c.interval {
				delete(c.entries, key)
			}
		}
	}
}
