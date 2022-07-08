package cache

import "time"

// Item represents a cache item.
type Item[T any] struct {
	Key       string
	Val       T
	ExpiredAt time.Time
}

// IsExpired returns true if the item is expired.
func (item Item[T]) IsExpired() bool {
	return !item.ExpiredAt.IsZero() && time.Now().After(item.ExpiredAt)
}
