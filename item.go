package cache

import "time"

// Item represents a cache item.
type Item[T any] struct {
	Val       T
	ExpiredAt time.Time
}

// IsExpired returns true if the item is expired.
func (item Item[T]) IsExpired() bool {
	return !item.ExpiredAt.IsZero() && time.Now().After(item.ExpiredAt)
}

// IsValid returns true if the item is valid.
func (item Item[T]) IsValid() bool {
	switch {
	case item.IsExpired():
		return false
	default:
		return true
	}
}
