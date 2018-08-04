package main

import "html/template"

// TmplCache stores HTML templates.
type TmplCache struct {
	data map[string]*template.Template
}

// NewTmplCache with no stored templates.
func NewTmplCache() *TmplCache {
	c := &TmplCache{}
	c.Clear()
	return c
}

// Get template at the key.
//
// Returns true if the template exists.
func (c TmplCache) Get(key string) (*template.Template, bool) {
	tmpl, ok := c.data[key]
	return tmpl, ok
}

// Put template at the key.
func (c *TmplCache) Put(key string, tmpl *template.Template) {
	c.data[key] = tmpl
}

// Clear the cache.
func (c *TmplCache) Clear() {
	c.data = make(map[string]*template.Template)
}
