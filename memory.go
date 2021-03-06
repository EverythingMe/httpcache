package httpcache

import (
	"errors"
	"github.com/EverythingMe/groupcache/lru"
	"net/url"
	"sync"
)

// A memoryCache implements a cache with files saved in Path.
// It is not safe for concurrency.
type memoryCache struct {
	Cache *lru.Cache
	mutex *sync.Mutex
}

func newMemoryCache(maxItems int) *memoryCache {
	return &memoryCache{
		Cache: lru.New(maxItems),
		mutex: &sync.Mutex{},
	}
}

// Get gets data saved for an URL if present in cache.
func (m memoryCache) Get(u *url.URL) (*Entry, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	data, ok := m.Cache.Get(u.String())
	if !ok {
		return nil, errors.New("not in cache")
	}

	return data.(*Entry), nil
}

// Put puts data of an URL in cache.
func (m memoryCache) Put(u *url.URL, e *Entry) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.Cache.Add(u.String(), e)
	return nil
}
