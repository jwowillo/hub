package cache

import "time"

// TimeSource returns a time.Time.
type TimeSource func() time.Time

// TimeDecorator is a Decorator which clears all entries after a time.Duration
// passes since the last clear.
type TimeDecorator struct {
	cache Cache

	timeSource TimeSource
	duration   time.Duration

	lastClear time.Time
}

// NewTimeDecorator that decorates the Caceh and clears all entries after a
// time.Duration passes after the time.Time of the last clear from the
// TimeSource.
func NewTimeDecorator(c Cache, ts TimeSource, d time.Duration) *TimeDecorator {
	return &TimeDecorator{
		cache: c,

		timeSource: ts,
		duration:   d,

		lastClear: ts(),
	}
}

// NewTimeDecoratorFactory that creates a TimeDecorator with the TimeSource and
// time.Duration.
func NewTimeDecoratorFactory(ts TimeSource, d time.Duration) DecoratorFactory {
	return func(c Cache) Decorator {
		return NewTimeDecorator(c, ts, d)
	}
}

// Get checks if the current time.Time from the TimeSource is after the time the
// last clear was performed and returns the value at the key if it isn't.
//
// Clears the Cache if it is and returns false otherwise.
func (d *TimeDecorator) Get(k string) (interface{}, bool) {
	now := d.timeSource()
	if now.After(d.lastClear.Add(d.duration)) {
		d.cache.Clear()
		d.lastClear = now
		return nil, false
	}
	return d.cache.Get(k)
}

// Put calls the decorated Cache's Put.
func (d *TimeDecorator) Put(k string, v interface{}) {
	d.cache.Put(k, v)
}

// Delete calls the decorated Cache's Delete.
func (d *TimeDecorator) Delete(k string) {
	d.cache.Delete(k)
}

// Clear calls the decorated Cache's Clear.
func (d *TimeDecorator) Clear() {
	d.cache.Clear()
}
