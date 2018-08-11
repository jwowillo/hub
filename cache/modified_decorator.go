package cache

import "time"

// HasBeenModified returns true if the value at the key has been modified since
// the time.Time.
type HasBeenModified func(string, time.Time) bool

// ModifiedDecorator is a Decorator which deletes entries that have been
// modified since a time.Time.
type ModifiedDecorator struct {
	cache Cache

	timeSource      TimeSource
	hasBeenModified HasBeenModified

	added map[string]time.Time
}

// NewModifiedDecorator that decorates the Cache and deletes entries that
// HasBeenModified says has been modified since the time returned from the
// TimeSource.
func NewModifiedDecorator(c Cache,
	ts TimeSource, hbm HasBeenModified) *ModifiedDecorator {
	return &ModifiedDecorator{
		cache: c,

		timeSource:      ts,
		hasBeenModified: hbm,

		added: make(map[string]time.Time),
	}
}

// NewModifiedDecoratorFactory that creates a ModifiedDecorator with the
// TimeSource and HasBeenModified.
func NewModifiedDecoratorFactory(ts TimeSource,
	hbm HasBeenModified) DecoratorFactory {
	return func(c Cache) Decorator {
		return NewModifiedDecorator(c, ts, hbm)
	}
}

// Get finds the time.Time the key was originally, calls the decorated Cache's
// Get, checks if the value at the key has been modified since that time,
// and returns the value if it hasn't.
//
// Returns false if the value doesn't exist. Calls the decorated Cache's Delete
// with the stale key and returns false if the value exists but has been
// modified since the time.
func (d *ModifiedDecorator) Get(k string) (interface{}, bool) {
	t, ok := d.added[k]
	if !ok {
		return nil, false
	}
	if d.hasBeenModified(k, t) {
		delete(d.added, k)
		d.cache.Delete(k)
		return nil, false
	}
	return d.cache.Get(k)
}

// Put notes the time.Time the key and value were added calls the decorated
// Cache's Put.
func (d *ModifiedDecorator) Put(k string, v interface{}) {
	d.added[k] = d.timeSource()
	d.cache.Put(k, v)
}

// Delete calls the decorated Cache's Delete.
func (d *ModifiedDecorator) Delete(k string) {
	delete(d.added, k)
	d.cache.Delete(k)
}

// Clear calls the decorated Cache's Clear.
func (d *ModifiedDecorator) Clear() {
	d.added = make(map[string]time.Time)
	d.cache.Clear()
}
