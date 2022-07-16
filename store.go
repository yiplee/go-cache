package cache

// Store defines the interface of a cache store.
type Store[K comparable, T any] interface {
	// Get returns the item with the given key.
	// If the item is not found, the second argument will be false.
	Get(key K) (Item[T], bool)

	// Set sets the item with the given key.
	Set(key K, item Item[T])

	// Delete deletes the item with the given key.
	Delete(key K)

	// Each calls the given function for each item in the store.
	Each(func(key K, item Item[T]) bool)
}

// defaultStore returns items as the default store.
func defaultStore[K comparable, T any]() Store[K, T] {
	return items[K, T]{}
}

// items is a Store implementation that uses a map to store items.
type items[K comparable, T any] map[K]Item[T]

func (m items[K, T]) Get(key K) (Item[T], bool) {
	item, found := m[key]
	return item, found
}

func (m items[K, T]) Set(key K, item Item[T]) {
	m[key] = item
}

func (m items[K, T]) Delete(key K) {
	delete(m, key)
}

func (m items[K, T]) Each(fn func(key K, item Item[T]) bool) {
	for k, v := range m {
		if !fn(k, v) {
			break
		}
	}
}
