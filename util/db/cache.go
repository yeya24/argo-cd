package db

import (
	"fmt"
	"sync"
)

type gitURLCache struct {
	cache map[string]string
	sync.RWMutex
}

func newGitURLCache() *gitURLCache {
	return &gitURLCache{cache: map[string]string{}}
}

func (c *gitURLCache) Load(url string) (string, bool) {
	c.RLock()
	normalizedURL, ok := c.cache[url]
	c.RUnlock()
	return normalizedURL, ok
}

func (c *gitURLCache) Store(url, normalizedURL string) {
	c.Lock()
	c.cache[url] = normalizedURL
	fmt.Printf("cache size is %d", len(c.cache))
	c.Unlock()
}
