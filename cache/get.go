package cache

// Fallback gets the value at a key if the key isn't in a Cache.
type Fallback func(string) interface{}

// Get the value at the key by checking in the Cache first and getting the value
// from the Fallback second.
//
// The value returned from the Fallback is stored back in the Cache.
func Get(c Cache, k string, fb Fallback) interface{} {
	v, ok := c.Get(k)
	if !ok {
		v = fb(k)
		c.Put(k, v)
	}
	return v
}
