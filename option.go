package cache

import "time"

type option struct {
	expiredAt time.Time
}

// OptionFunc is a function that can be used to configure a cache.
type OptionFunc func(*option)

// WithTTL returns an OptionFunc that sets the expiredAt to the current time plus the given duration.
func WithTTL(dur time.Duration) OptionFunc {
	return func(opt *option) {
		opt.expiredAt = time.Now().Add(dur)
	}
}

// WithExpiredAt returns an OptionFunc that sets the expiredAt to the given time.
func WithExpiredAt(at time.Time) OptionFunc {
	return func(opt *option) {
		opt.expiredAt = at
	}
}
