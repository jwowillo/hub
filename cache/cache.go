package cache

// Cache associates keys with values.
type Cache interface {
	// Get the key.
	//
	// Returns false if the key isn't in the Cache.
	Get(string) (interface{}, bool)
	// Put the value at the key.
	Put(string, interface{})
	// Delete the key.
	Delete(string)
	// Clear the Cache.
	Clear()
}
