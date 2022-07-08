package cache

// Store defines the interface of a cache store.
type Store[T any] interface {
	// Get returns the item with the given key.
	// If the item is not found, the second argument will be false.
	Get(key string) (Item[T], bool)

	// Set sets the item with the given key.
	Set(key string, item Item[T])
}

// defaultStore returns items as the default store.
func defaultStore[T any]() Store[T] {
	return items[T]{}
}

// items is a Store implementation that uses a map to store items.
type items[T any] map[string]Item[T]

func (m items[T]) Get(key string) (Item[T], bool) {
	item, found := m[key]
	return item, found
}

func (m items[T]) Set(key string, item Item[T]) {
	m[key] = item
}
