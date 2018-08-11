package cache

// MemoryCache which stores values at string keys in memory.
type MemoryCache struct {
	data map[string]interface{}
}

// NewMemoryCache makes an empty MemoryCache.
func NewMemoryCache() *MemoryCache {
	c := &MemoryCache{}
	c.Clear()
	return c
}

// Get the value at the key.
//
// Returns true if the value exists and false otherwise.
func (c MemoryCache) Get(k string) (interface{}, bool) {
	v, ok := c.data[k]
	return v, ok
}

// Put the value at the key.
func (c *MemoryCache) Put(k string, v interface{}) {
	c.data[k] = v
}

// Delete the value and the key.
func (c *MemoryCache) Delete(k string) {
	delete(c.data, k)
}

// Clear the MemoryCache.
func (c *MemoryCache) Clear() {
	c.data = make(map[string]interface{})
}
