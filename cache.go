// package cache provides a simple in-memory cache.
//
// All methods are safe for concurrent use.
package cache

import (
	"sync"
)

// NewWithStore returns a new Cache with the given store.
func NewWithStore[T any](store Store[T]) *Cache[T] {
	return &Cache[T]{
		store: store,
	}
}

// New returns a new Cache with the default store.
func New[T any]() *Cache[T] {
	return NewWithStore(defaultStore[T]())
}

// Cache represents a cache.
type Cache[T any] struct {
	store Store[T]
	mux   sync.RWMutex
}

// Get returns the value for the given key.
// If the value is not found, the second argument will be false.
func (c *Cache[T]) Get(key string) (T, bool) {
	c.mux.RLock()
	item, found := c.store.Get(key)
	c.mux.RUnlock()

	if found && item.IsValid() {
		return item.Val, true
	}

	var t T
	return t, false
}

// Contain check if the key cached
func (c *Cache[T]) Contain(key string) bool {
	_, ok := c.Get(key)
	return ok
}

// Set sets the value for the given key.
func (c *Cache[T]) Set(key string, val T, opts ...OptionFunc) {
	var opt option
	for _, fn := range opts {
		fn(&opt)
	}

	item := Item[T]{
		Val:       val,
		ExpiredAt: opt.expiredAt,
	}

	c.mux.Lock()
	c.store.Set(key, item)
	c.mux.Unlock()
}

// Delete deletes the value for the given key.
func (c *Cache[T]) Delete(key string) {
	c.mux.Lock()
	c.store.Delete(key)
	c.mux.Unlock()
}

// Keys returns a list of all keys in the cache.
func (c *Cache[T]) Keys() []string {
	var keys []string
	c.Each(func(key string, _ T) bool {
		keys = append(keys, key)
		return true
	})

	return keys
}

// Each iterates over all items in the cache.
func (c *Cache[T]) Each(fn func(key string, val T) bool) {
	c.mux.RLock()
	c.store.Each(func(key string, item Item[T]) bool {
		if item.IsValid() {
			return fn(key, item.Val)
		}

		return true
	})
	c.mux.RUnlock()
}
